package store

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"ordernationn/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type AdminUserStore struct {
	adminUser domain.AdminUser
}

type SQLAdminUserRepository struct {
	DB *sql.DB
}

func NewSQLAdminUserRepository(db *sql.DB) *SQLAdminUserRepository {
	return &SQLAdminUserRepository{
		DB: db,
	}
}

const columnList = `id, first_name, last_name, gender, age, address, email, admin_info, password_hash, created_at, updated_at`

var (
	ErrAdminNotFound   = errors.New("admin not found")
	ErrInvalidPassword = errors.New("invalid password")
)

func (r *SQLAdminUserRepository) Login(ctx context.Context, email string, password string) (domain.AdminUser, error) {
	if email == "" || password == "" {
		return domain.AdminUser{}, errors.New("credentials can not be empty")
	}

	query := fmt.Sprintf(`SELECT %s FROM admin_users WHERE email = $1 LIMIT 1`, columnList)
	var addrsJSON, infoJSON []byte
	var passwordHash string

	row := r.DB.QueryRowContext(ctx, query, email)
	var u domain.AdminUser
	err := row.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Gender, &u.Age, &addrsJSON, &u.Email,
		&infoJSON, &passwordHash, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u, ErrAdminNotFound
		}
		return u, fmt.Errorf("unmarshal address %w", err)
	}
	if len(addrsJSON) > 0 {
		if err := json.Unmarshal(addrsJSON, &u.Address); err != nil {
			return u, fmt.Errorf("unmarshal address: %w", err)
		}
	}
	if len(infoJSON) > 0 {
		if err := json.Unmarshal(infoJSON, &u.AdminInfo); err != nil {
			return u, fmt.Errorf("unmarshal admin_info: %w", err)
		}
	}
	log.Println(password, passwordHash)
	if err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)); err != nil {
		return domain.AdminUser{}, ErrInvalidPassword
	}
	u.Password = ""
	return u, nil
}
