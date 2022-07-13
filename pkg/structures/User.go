package structures

import (
	"github.com/google/uuid"
)

type User struct{
  // username so we can easily reference a user
	Username string
  // token that is a combination of their password + their SessionId that was used to create the password
	Token uuid.UUID 
  // bananas =D
  IsAnonymous bool
}

func CreateUser (name string, token uuid.UUID) *User {
  return &User{
    Username: name,
    Token: token,
    IsAnonymous: false,
  }
}

func (self *User) AuthenticateUser () {
}