package repository

import (
	"context"

	"github.com/mars1385/e-com-go/auth/ent"
	"github.com/mars1385/e-com-go/auth/ent/user"
)

type User struct {
	*ent.UserClient
}

func (u *User) GetById(id int, ctx context.Context) (*ent.User, error) {

	return u.Query().Where(user.ID(id)).Only(ctx)
}
