package persistence

import (
	"database/sql"

	shoppingcart "github.com/go-microservices/shopping-cart-service/internal/domain/shopping_cart"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	template_errors "github.com/go-microservices/shopping-cart-service/internal/shared/error"
	_ "github.com/go-sql-driver/mysql"
)

type shoppingCartMySQLRepo struct {
	logger *logrus.Logger
	db     DataBaseConn
}

func NewshoppingCartMySQLRepo(db DataBaseConn, logger *logrus.Logger) shoppingCartMySQLRepo {
	return shoppingCartMySQLRepo{
		logger: logger,
		db:     db,
	}
}

func (sc shoppingCartMySQLRepo) Create(spc shoppingcart.ShoppingCart) error {
	sc.logger.Info("On ShoppingCart repository - creating shopping cart")

	query := "INSERT INTO shopping_cart (id, user_id) VALUES (?, ?)"
	randomUUID := uuid.NewV4()

	_, err := sc.db.Exec(query, randomUUID.String(), spc.UserID)
	if err != nil {
		sc.logger.Errorf("error saving shopping cart: %v", err)
		return err
	}

	return nil
}

func (sc shoppingCartMySQLRepo) GetByUserID(userId *uuid.UUID) (*shoppingcart.ShoppingCart, error) {
	sc.logger.Info("On ShoppingCart repository - get shopping cart info")

	query := "SELECT * FROM shopping_cart WHERE  user_id = ?"

	row := sc.db.QueryRow(query, userId.String())
	if row.Err() != nil {
		sc.logger.Errorf("error getting shopping cart: %v", row)
		return nil, row.Err()
	}

	var cart shoppingcart.ShoppingCart
	err := row.Scan(&cart.ID, &cart.UserID)
	if err == sql.ErrNoRows {
		return nil, template_errors.NewRecordNotFoundError("shopping cart not found", template_errors.RecordNotFound)
	}

	if err != nil {
		sc.logger.Errorf("error getting shopping cart values: %v", err)
		return nil, err
	}

	return &cart, nil
}
