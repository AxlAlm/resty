package resty_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
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

type GetOneParams struct {
	ID string
}

func (uc UserController) Create(ctx context.Context, u UserCreate) (User, error) {
	return User{Name: "test"}, nil
}

func (uc UserController) GetOne(ctx context.Context, id string) (User, error) {
	return User{Name: "test"}, nil
}

func TestResource_Setup(t *testing.T) {

	mux := http.NewServeMux()
	resty.Resource(mux, "users", UserController{}, []string{})
	srv := httptest.NewServer(mux)
	defer srv.Close()

	req, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("%s/%s/%s?test=test", srv.URL, "users", "123"),
		nil,
	)

	res, _ := http.DefaultClient.Do(req)
	fmt.Println(res)
}
