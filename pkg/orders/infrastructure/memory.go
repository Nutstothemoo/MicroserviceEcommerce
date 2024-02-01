package memory

import (
	// "database/sql"
	"microservice/pkg/orders/domain"
	// _ "github.com/lib/pq"
)

type MemoryRepository struct {
    orders []orders.Order	
}

func NewMemoryRepository() *MemoryRepository {
    return &MemoryRepository{[]orders.Order{}}
}

func (m *MemoryRepository) Save(orderToSave *orders.Order) error {
    for i, o := range m.orders {
        if o.ID() == orderToSave.ID() {
            m.orders[i] = *orderToSave
            return nil
        }
    }
    m.orders = append(m.orders, *orderToSave)
    return nil	
}

func (m MemoryRepository) ByID(id orders.OrderID) (*orders.Order, error){
    for _, o := range m.orders {
        if o.ID() == id {
            return &o, nil
        }
    }
    return nil, orders.ErrNotFound
}

func (m *MemoryRepository) AllOrders()([]orders.Order, error) {
    return m.orders, nil
}

// POSTGRES

// type PostgresRepository struct {
// 	db *sql.DB
// }

// func NewPostgresRepository(dataSourceName string) (*PostgresRepository, error) {
// 	db, err := sql.Open("postgres", dataSourceName)
// 	if err != nil {
// 			return nil, err
// 	}
// 	if err := db.Ping(); err != nil {
// 			return nil, err
// 	}
// 	return &PostgresRepository{db}, nil
// }

// func (p *PostgresRepository) Save(orderToSave *orders.Order) error {
// 	// This is a placeholder SQL statement. You should replace it with your actual SQL statement.
// 	sqlStatement := `
// 	INSERT INTO orders (id, name, price)
// 	VALUES ($1, $2, $3)
// 	ON CONFLICT (id) DO UPDATE 
// 	SET name = $2, price = $3;`

// 	_, err := p.db.Exec(sqlStatement, orderToSave.OrderID(), orderToSave.Name(), orderToSave.Price())
// 	return err
// }

// func (p *PostgresRepository) ByID(id orders.OrderID) (*orders.Order, error) {
// 	// This is a placeholder SQL statement. You should replace it with your actual SQL statement.
// 	sqlStatement := `SELECT id, name, price FROM orders WHERE id = $1;`
// 	row := p.db.QueryRow(sqlStatement, id)

// 	var order orders.Order
// 	err := row.Scan(&order.ID, &order.Name, &order.Price)
// 	if err == sql.ErrNoRows {
// 			return nil, orders.ErrNotFound
// 	} else if err != nil {
// 			return nil, err
// 	}

// 	return &order, nil
// }

// func (p *PostgresRepository) AllOrders() ([]orders.Order, error) {
// 	// This is a placeholder SQL statement. You should replace it with your actual SQL statement.
// 	sqlStatement := `SELECT id, name, price FROM orders;`
// 	rows, err := p.db.Query(sqlStatement)
// 	if err != nil {
// 			return nil, err
// 	}
// 	defer rows.Close()

// 	var orders []orders.Order
// 	for rows.Next() {
// 			var order orders.Order
// 			if err := rows.Scan(&order.ID, &order.Name, &order.Price); err != nil {
// 					return nil, err
// 			}
// 			orders = append(orders, order)
// 	}

// 	if err := rows.Err(); err != nil {
// 			return nil, err
// 	}

// 	return orders, nil
// }