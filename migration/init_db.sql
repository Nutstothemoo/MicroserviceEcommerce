CREATE TABLE address (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) ,
    street VARCHAR(255) ,
    city VARCHAR(255),
    postal_code VARCHAR(255),
    country VARCHAR(255) NOT NULL
);

CREATE TABLE products (
	id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price_cents INT NOT NULL,
    currency VARCHAR(3) NOT NULL
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL,
    address_id INT,
    paid BOOLEAN DEFAULT FALSE,
    FOREIGN KEY (product_id) REFERENCES products(id),
    FOREIGN KEY (address_id) REFERENCES address(id)
);

CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    order_id INT NOT NULL,
    amount_cents INT NOT NULL,
    currency VARCHAR(3) NOT NULL,
    paid_at TIMESTAMP NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id)
);