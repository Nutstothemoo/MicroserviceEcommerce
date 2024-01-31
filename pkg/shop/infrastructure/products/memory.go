package memory

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"microservice/pkg/shop/domain/products"
	"go.mongodb.org/mongo-driver/mongo/options"
	"database/sql"
	_ "github.com/lib/pq"
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


// MONGODB

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

// POSTGRES 


func NewPostgresRepository(dataSourceName string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
			return nil, err
	}

	return &PostgresRepository{db}, nil
}

func (p *PostgresRepository) Save(productToSave *product.Product) error {
	// Implémentez la logique pour sauvegarder le produit dans PostgreSQL
	// ...
	return nil
}

func (p *PostgresRepository) ByID(id products.ID) (*products.Product, error) {
	// Implémentez la logique pour obtenir un produit par ID de PostgreSQL
	// ...
	return nil, nil
}

func (p *PostgresRepository) AllProducts() ([]products.Product, error) {
	// Implémentez la logique pour obtenir tous les produits de PostgreSQL
	// ...
	return nil, nil
}