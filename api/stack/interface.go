package stack

import models "github.com/wolfgunblood/go-sdk-stack/model"

type UserService interface {
    //examples
    GetUser(userID string) (*models.User, error)
    CreateUser(user *models.User) error
}
