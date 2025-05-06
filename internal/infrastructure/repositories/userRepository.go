package repositories

import (
	_ "github.com/go-sql-driver/mysql"
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

	query := "INSERT INTO tg_users (user_id, username,amount,tron_amount,tron_address, eth_address,eth_amount, associates) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	row := tx.QueryRow(query, user.UserID, user.Username, user.Amount, user.TronAmount, user.TronAddress, user.EthAddress, user.EthAmount, user.Associates)
	//row := tx.QueryRow(query, 1, user.Username, user.TronAmount, user.TronAddress, user.EthAddress, user.EthAmount, user.Associates)
	if row.Err() != nil {

		log.Println("add err:", row.Err())
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

func (r *UserRepository) UpdateAddress(user domain.User) error {
	query := "UPDATE tg_users SET address = ? , private_key = ?  WHERE username = ?"
	_, err := r.db.Exec(query, user.Address, user.Key, user.Username)
	return err
}

func (r *UserRepository) UpdateTimes(_times uint64, _username string) error {
	query := "UPDATE tg_users SET times = ?  WHERE username = ?"
	_, err := r.db.Exec(query, _times, _username)
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
	err := r.db.Get(&jason, "SELECT  user_id,username,amount,associates, tron_amount,tron_address,eth_address,eth_amount,times FROM tg_users WHERE username=?", _username)

	//log.Println(err)
	return jason, err
}
func (r *UserRepository) GetByUserID(_userID string) (domain.User, error) {
	jason := domain.User{}
	err := r.db.Get(&jason, "SELECT  id, user_id,username,amount,associates, tron_amount,tron_address,eth_address,eth_amount,times FROM tg_users WHERE associates=?", _userID)
	return jason, err
}

func (r *UserRepository) FetchNewestAddress() ([]domain.User, error) {
	query := `SELECT address,associates
    FROM 
      sys_address  where disable=0 ;
    `
	var addresses []domain.User
	err := r.db.Select(&addresses, query)
	return addresses, err
}
func (r *UserRepository) DisableTronAddress(_address string) error {
	query := "UPDATE sys_address SET disable = 1 WHERE address = ?"
	_, err := r.db.Exec(query, _address)
	return err
}

func (r *UserRepository) BindChat(_associates string, _username string) error {
	query := "UPDATE tg_users SET associates = ? WHERE username = ?"
	_, err := r.db.Exec(query, _associates, _username)
	return err
}

func (r *UserRepository) BindTronAddress(_address string, _username string) error {
	query := "UPDATE tg_users SET tron_address = ? WHERE username = ?"
	_, err := r.db.Exec(query, _address, _username)
	return err
}

func (r *UserRepository) BindEthereumAddress(_address string, _username string) error {
	query := "UPDATE tg_users SET eth_address = ? WHERE username = ?"
	_, err := r.db.Exec(query, _address, _username)
	return err
}

func (r *UserRepository) NotifyTronAddress() ([]domain.User, error) {
	query := `SELECT t.username,t.tron_address,t.associates
    FROM
        tg_users t
    LEFT JOIN
        sys_address s ON t.tron_address = s.address

    WHERE s.disable = 0;
    `
	var addresses []domain.User
	err := r.db.Select(&addresses, query)
	return addresses, err
}
func (r *UserRepository) NotifyEthereumAddress() ([]domain.User, error) {
	query := `SELECT t.username,t.eth_address,t.associates
    FROM
        tg_users t
    LEFT JOIN
        sys_address s ON t.eth_address = s.address

    WHERE s.disable = 0;
    `
	var addresses []domain.User
	err := r.db.Select(&addresses, query)
	return addresses, err
}

//query := `SELECT t.username
//    FROM
//        tg_users t
//    LEFT JOIN
//        sys_address s ON t.tron_address = s.address
//    WHERE s.disable = 0
//    GROUP BY
//        h.id, h.name, h.description, h.images, h.create_at, h.deadline, h.update_at;
//    `
