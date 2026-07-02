CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    product_name VARCHAR(255) NOT NULL,
    price NUMERIC(10,2) NOT NULL
);

INSERT INTO product (product_name, price)
VALUES
    ('Notebook', 4500.00),
    ('Mouse Gamer', 189.90),
    ('Teclado Mecânico', 399.99),
    ('Monitor 24"', 899.90);