package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/labstack/gommon/log"
	"github.com/lib/pq"
	"time"
	"timewise/db"
	"timewise/model"
)

type UserRepoImpl struct {
	sql *db.SQL
}

// NewUserRepo create object working with user logic
func NewUserRepo(sql *db.SQL) UserRepo {
	return UserRepoImpl{
		sql: sql,
	}
}

func (u UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
		INSERT INTO users(user_id, phone, password, email, avatar, role, full_name, created_at, updated_at)
		VALUES(:user_id, :phone, :password, :email, :avatar, :role, :full_name, :created_at, :updated_at)
	`
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, errors.New("Người dùng đã tồn tại")
			}
		}
		return user, errors.New("Đăng ký thất bại")
	}

	return user, nil
}

func (u UserRepoImpl) SelectUserByEmail(context context.Context, email string) (model.User, error) {
	var user = model.User{}

	statement := `SELECT * FROM users WHERE email=$1`
	err := u.sql.Db.GetContext(context, &user, statement, email)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("Không tồn tại người dùng này")
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}

func (u UserRepoImpl) SelectUserById(context context.Context, userId string) (model.User, error) {
	var user = model.User{}

	statement := `SELECT * FROM users WHERE user_id=$1`
	err := u.sql.Db.GetContext(context, &user, statement, userId)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("Không tồn tại người dùng này")
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}

func (u UserRepoImpl) SelectUsers(context context.Context) ([]model.User, error) {
	var user []model.User

	statement := `SELECT * FROM users ORDER BY created_at DESC`
	err := u.sql.Db.SelectContext(context, &user, statement)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, errors.New("Chưa có người dùng nào")
		}
		log.Error(err.Error())
		return user, err
	}

	return user, nil
}
