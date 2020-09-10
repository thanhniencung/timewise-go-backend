package repository

import (
	"context"
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
	return model.User{}, nil
}