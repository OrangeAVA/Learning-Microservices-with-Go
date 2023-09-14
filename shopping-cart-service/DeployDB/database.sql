USE shopping_cart;

CREATE TABLE shopping_cart (
    id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE product (
    id CHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    quantity VARCHAR(50) NOT NULL,
    product_description VARCHAR(50) NOT NULL,
    shopping_cart_id CHAR(36) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (shopping_cart_id) REFERENCES shopping_cart(id)
);