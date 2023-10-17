package dao

import (
	"time"

	"github.com/Orendev/gokeeper/internal/pkg/domain/user"
	"github.com/Orendev/gokeeper/pkg/type/email"
	"github.com/Orendev/gokeeper/pkg/type/name"
	"github.com/Orendev/gokeeper/pkg/type/password"
	"github.com/Orendev/gokeeper/pkg/type/role"
	"github.com/Orendev/gokeeper/pkg/type/token"
	"github.com/google/uuid"
)

var TableNameUser = "users"

type User struct {
	ID        uuid.UUID `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`

	Password string `db:"password"`
	Email    string `db:"email"`
	Role     string `db:"role"`
	Token    string `db:"token"`
	Name     string `db:"name"`
}

var ColumnUser = []string{
	"id",
	"password",
	"email",
	"name",
	"role",
	"token",
	"created_at",
	"updated_at",
}

func ToDomainUser(dao *User) (*user.User, error) {

	emailObject, err := email.New(dao.Email)
	if err != nil {
		return nil, err
	}

	passwordObject, err := password.New(dao.Password)
	if err != nil {
		return nil, err
	}

	nameObject, err := name.New(dao.Name)
	if err != nil {
		return nil, err
	}

	roleObject, err := role.New(dao.Role)
	if err != nil {
		return nil, err
	}

	tokenObject := token.New(dao.Token)
	if err != nil {
		return nil, err
	}

	result, err := user.NewWithID(
		dao.ID,
		*passwordObject,
		*emailObject,
		*nameObject,
		*roleObject,
		dao.CreatedAt,
		dao.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	result.SetToken(*tokenObject)

	return result, nil
}

func ToDomainUsers(dao []*User) ([]*user.User, error) {
	var result = make([]*user.User, len(dao))
	for i, v := range dao {
		c, err := ToDomainUser(v)
		if err != nil {
			return nil, err
		}
		result[i] = c
	}
	return result, nil
}
