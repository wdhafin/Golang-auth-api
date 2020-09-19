package auth

import "github.com/wdhafin/Golang-auth-api/entity"

// Repository represent the userauthentication's repository contract
type Repository interface {
	Register(entity.User) (*entity.User, error)
	CheckUserExist(string) (*entity.User, error)
}
