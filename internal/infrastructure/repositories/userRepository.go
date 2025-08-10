package repositories

import (
	"context"
	_ "github.com/go-sql-driver/mysql"

	"gorm.io/gorm"
	"ushield_bot/internal/domain"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
func (r *UserRepository) Update2(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}
func (r *UserRepository) UpdateBackupChat(ctx context.Context, backup string, _associates int64) error {
	query := "UPDATE tg_users SET backup_chat_id = ?  WHERE associates = ?"
	tx := r.db.Exec(query, backup, _associates)
	return tx.Error
}

func (r *UserRepository) Create2(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) Create(user domain.User) error {

	query := "INSERT INTO tg_users (user_id, username,amount,tron_amount,tron_address, eth_address,eth_amount, associates) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	tx := r.db.Exec(query, user.UserID, user.Username, user.Amount, user.TronAmount, user.TronAddress, user.EthAddress, user.EthAmount, user.Associates)

	return tx.Error
}

func (r *UserRepository) Update(user domain.User) error {
	query := "UPDATE tg_users SET associates = $1, tron_amount = $2 WHERE username = $3"
	tx := r.db.Exec(query, user.Associates, user.TronAmount, user.Username)
	return tx.Error
}

func (r *UserRepository) UpdateAddress(user domain.User) error {
	query := "UPDATE tg_users SET address = ? , private_key = ?  WHERE id = ?"
	tx := r.db.Exec(query, user.Address, user.Key, user.Id)
	return tx.Error
}

func (r *UserRepository) UpdateTimes(_times uint64, _username string) error {
	query := "UPDATE tg_users SET times = ?  WHERE username = ?"
	tx := r.db.Exec(query, _times, _username)
	return tx.Error
}

func (r *UserRepository) UpdateTimesByChatID(_times uint64, _chatID int64) error {
	query := "UPDATE tg_users SET times = ?  WHERE associates = ?"
	tx := r.db.Exec(query, _times, _chatID)
	return tx.Error
}

//associates VARCHAR(255),
//amount VARCHAR(255) ,
//tron_amount VARCHAR(255),
//tron_address VARCHAR(50),
//eth_address VARCHAR(50),
//eth_amount VARCHAR(255),

func (r *UserRepository) GetByUsername(_username string) (domain.User, error) {

	jason := domain.User{}

	err := r.db.Where(" username=?", _username).First(&jason).Error

	return jason, err
}
func (r *UserRepository) GetByUserID(_chatID int64) (domain.User, error) {
	jason := domain.User{}

	err := r.db.Where(" associates=?", _chatID).First(&jason).Error

	return jason, err
}

func (r *UserRepository) FetchNewestAddress() ([]domain.User, error) {
	query := `SELECT address,associates
    FROM 
      sys_address  where disable=0 ;
    `
	var addresses []domain.User
	r.db.Select(&addresses, query)
	return addresses, nil
}
func (r *UserRepository) DisableTronAddress(_address string) error {
	query := "UPDATE sys_address SET disable = 1 WHERE address = ?"
	tx := r.db.Exec(query, _address)
	return tx.Error
}

func (r *UserRepository) BindChat(_associates string, _username string) error {
	query := "UPDATE tg_users SET associates = ? WHERE username = ?"
	tx := r.db.Exec(query, _associates, _username)
	return tx.Error
}

func (r *UserRepository) BindTronAddress(_address string, _username string) error {
	query := "UPDATE tg_users SET tron_address = ? WHERE username = ?"
	tx := r.db.Exec(query, _address, _username)
	return tx.Error
}

func (r *UserRepository) BindEthereumAddress(_address string, _username string) error {
	query := "UPDATE tg_users SET eth_address = ? WHERE username = ?"
	tx := r.db.Exec(query, _address, _username)
	return tx.Error
}

func (r *UserRepository) NotifyTronAddress() ([]domain.User, error) {
	query := `SELECT t.username,t.tron_address,t.associates
    FROM
        tg_users t
    LEFT JOIN
        sys_address s ON t.tron_address = s.address

    WHERE s.disable = 0
    `
	var addresses []domain.User
	r.db.Select(&addresses, query)
	return addresses, nil
}
func (r *UserRepository) NotifyEthereumAddress() ([]domain.User, error) {
	query := `SELECT t.username,t.eth_address,t.associates
    FROM
        tg_users t
    LEFT JOIN
        sys_address s ON t.eth_address = s.address

    WHERE s.disable = 0
    `
	var addresses []domain.User
	r.db.Select(&addresses, query)
	return addresses, nil
}
