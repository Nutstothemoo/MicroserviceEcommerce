package memory

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"microservice/pkg/shop/domain/products"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MemoryRepository struct {
	products []products.Product	
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{[]products.Product{}}
}

func (m *MemoryRepository) Save( productToSave *product.Product) error {
	for i, p := range m.products {
		if p.ID() == productToSave.ID() {
			m.products[i] = productToSave
			return nil
		}
	}
	m.products = append(m.products, productToSave)
	return nil	
}

func (m MemoryRepository) ByID(id products.ID) (*products.Product, error){
	for _, p := range m.products {
		if p.ID() == id {
			return &p, nil
		}
	}
	return nil, product.ErrNotFound
}

func (m *MemoryRepository) AllProducts()([]products.Product, error) {
	return m.products, nil
}

type MongoDBRepository struct {
    collection *mongo.Collection
}

func NewMongoDBRepository(mongoURI, dbName, collectionName string) (*MongoDBRepository, error) {
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
    if err != nil {
        return nil, err
    }

    collection := client.Database(dbName).Collection(collectionName)

    return &MongoDBRepository{collection}, nil
}

func (m *MongoDBRepository) Save(productToSave *product.Product) error {
    _, err := m.collection.ReplaceOne(
        context.Background(),
        bson.M{"id": productToSave.ID()},
        productToSave,
        options.Replace().SetUpsert(true),
    )
    return err
}

func (m *MongoDBRepository) ByID(id products.ID) (*products.Product, error) {
	var productFound products.Product
	err := m.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&productFound)
	if err != nil {
			if err == mongo.ErrNoDocuments {
					return nil, product.ErrNotFound
			}
			return nil, err
	}
	return &productFound, nil
}

func (m *MongoDBRepository) AllProducts() ([]products.Product, error) {
    cursor, err := m.collection.Find(context.Background(), bson.M{})
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    var products []products.Product
    if err = cursor.All(context.Background(), &products); err != nil {
        return nil, err
    }
    return products, nil
}