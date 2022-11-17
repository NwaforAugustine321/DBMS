package userRepo

import (
	"context"
	"dbms/repo/interfaces"
)

type userSource struct{}

func NewUserSource() interfaces.UserSourceInterface {
	return &userSource{}
}

func (user *userSource) CreateUser(ctx context.Context, data interfaces.User) {

}


