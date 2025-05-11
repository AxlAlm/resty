package resty_test

import (
	"context"
	"testing"

	"github.com/AxlAlm/resty"
)

type UserCreate struct {
	Name string
}

type User struct {
	Name string
}

type UserController struct{}

func (uc UserController) Create(ctx context.Context, u UserCreate) (User, error) {
	return User{Name: "test"}, nil
}

func (uc UserController) GetOne(ctx context.Context, id string) (User, error) {
	return User{Name: "test"}, nil
}

func TestResource_Setup(t *testing.T) {
	resty.Resource(nil, "users", UserController{}, []string{})

}
