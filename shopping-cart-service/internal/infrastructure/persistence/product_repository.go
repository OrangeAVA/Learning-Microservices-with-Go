package persistence

import (
	"database/sql"
	"log"

	"github.com/go-microservices/shopping-cart-service/internal/domain/product"
	template_errors "github.com/go-microservices/shopping-cart-service/internal/shared/error"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type ProductMySQLRepo struct {
	logger *logrus.Logger
	db     *sql.DB
}

func NewProductMySQLRepo(db *sql.DB, logger *logrus.Logger) ProductMySQLRepo {
	return ProductMySQLRepo{
		logger: logger,
		db:     db,
	}
}

func (pr ProductMySQLRepo) Create(product product.Product) error {
	log.Println("On Create product repository")

	query := "INSERT INTO product (id,name,quantity,product_description, shopping_cart_id) VALUES (?, ?,?,?,?)"

	randomUUID := uuid.NewV4()

	_, err := pr.db.Exec(query, randomUUID.String(),
		product.Name, product.Quantity, product.Description, product.ShoppingCartID)
	if err != nil {
		log.Println("error saving product: ", err.Error())
		return err
	}

	return nil
}

func (pr ProductMySQLRepo) Get(shoppingCartID uuid.UUID) ([]product.Product, error) {
	pr.logger.Info("On List products repository")

	products := []product.Product{}

	query := "SELECT * FROM product WHERE  shopping_cart_id = ?"
	rows, err := pr.db.Query(query, shoppingCartID.String())

	if err != nil {
		pr.logger.Errorf("error getting products cart: %v", err)
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		var product product.Product

		err := rows.Scan(&product.ID, &product.Name, &product.Quantity,
			&product.Description, &product.ShoppingCartID)
		if err != nil {
			pr.logger.Errorf("error scanning row: %v", err)
			continue
		}

		products = append(products, product)
	}

	return products, nil
}

func (pr ProductMySQLRepo) Delete(productID uuid.UUID) error {
	pr.logger.Info("On Delete product repository")

	if err := pr.GetByID(productID); err != nil {
		return err
	}

	query := "DELETE FROM product WHERE id = ?"
	_, err := pr.db.Exec(query, productID)
	if err != nil {
		pr.logger.Errorf("On Delete product repository - error deleting product: %v", err)
		return err
	}

	return nil
}

func (pr ProductMySQLRepo) UpdateQuantity(productID uuid.UUID, quantity int) error {
	pr.logger.Info("On Update product quantity repository")

	if err := pr.GetByID(productID); err != nil {
		return err
	}

	query := "UPDATE product SET quantity = ? WHERE id = ?"
	_, err := pr.db.Exec(query, quantity, productID)
	if err != nil {
		pr.logger.Errorf("error updating product: %v", err)
		return err
	}

	return nil
}

func (pr ProductMySQLRepo) GetByID(productID uuid.UUID) error {
	pr.logger.Info("On get product by ID quantity repository")

	querySelect := "SELECT id FROM product WHERE id = ?"
	var existingID uuid.UUID

	err := pr.db.QueryRow(querySelect, productID).Scan(&existingID)
	if err == sql.ErrNoRows {
		pr.logger.Infof("On Get product quantity repository - product not found %d", productID)
		return template_errors.NewRecordNotFoundError("product not found", template_errors.RecordNotFound)
	} else if err != nil {
		pr.logger.Errorf("error quering product: %v", err)
		return err
	}

	return nil
}
