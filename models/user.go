package models

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
	user.ID = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}
