package models

import (
	"errors"
	"fmt"
)

//User a user
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers dsaa
func GetUsers() []*User {
	return users
}

//AddUser as
func AddUser(user User) (User, error) {
	if user.ID != 0 {
		return User{}, errors.New("User already has an ID")
	}

	user.ID = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}

// GetUserByID dfgasg
func GetUserByID(id int) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}

	return User{}, fmt.Errorf("User with id '%v' not found", id)
}

// UpdateUser fdasg
func UpdateUser(user User) (User, error) {
	for i, toUpdate := range users {
		if user.ID == toUpdate.ID {
			users[i] = &user
			return user, nil
		}
	}

	return User{}, fmt.Errorf("User with id %v not found", user.ID)
}

// RemoveUserByID dfgaf
func RemoveUserByID(id int) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with id %v not found", id)
}
