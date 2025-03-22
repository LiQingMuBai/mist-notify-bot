package repositories

import (
	"github.com/jmoiron/sqlx"
	"homework_bot/internal/domain"
	"log"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(user domain.User) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	query := "INSERT INTO tg_users (id, username,amount,tron_amount,tron_address, eth_address,eth_amount, associates) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	row := tx.QueryRow(query, user.Id, user.Username, user.Amount, user.TronAmount, user.TronAddress, user.EthAddress, user.EthAmount, user.Associates)
	//row := tx.QueryRow(query, 1, user.Username, user.TronAmount, user.TronAddress, user.EthAddress, user.EthAmount, user.Associates)
	if row.Err() != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *UserRepository) Update(user domain.User) error {
	query := "UPDATE tg_users SET associates = $1, tron_amount = $2 WHERE username = $3"
	_, err := r.db.Exec(query, user.Associates, user.TronAmount, user.Username)
	return err
}

//associates VARCHAR(255),
//amount VARCHAR(255) ,
//tron_amount VARCHAR(255),
//tron_address VARCHAR(50),
//eth_address VARCHAR(50),
//eth_amount VARCHAR(255),

func (r *UserRepository) GetByUsername(_username string) (domain.User, error) {
	//query := "SELECT id,username,associates,amount, tron_amount,tron_address,eth_address,eth_amount FROM tg_users WHERE username = ? "
	//var user domain.User
	//
	//err := r.db.Get(&user, query, _username)

	jason := domain.User{}
	err := r.db.Get(&jason, "SELECT  id,username,amount,associates, tron_amount,tron_address,eth_address,eth_amount,create_at,update_at FROM tg_users WHERE username=$1", _username)

	log.Println(err)
	return jason, err
}
