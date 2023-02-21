package user

func GetUserByUsername(username string) (*User, error) {
	user := &User{}
	err := Repository.App.DB.First(user, "username = ?", username)
	return user, err.Error
}

// func GetById(userId string) *User {
// 	// TODO : implement
// }

// func GetUsers() []*User {
// 	// TODO : implement
// }

func InsertUser(user *User) error {
	result := Repository.App.DB.Create(user)
	return result.Error
}

// func InsertUsers(users []*User) error {
// 	//TODO : implement
// }

// func UpdateUserById(userId string, userDetails *User) error {
// 	//TODO : implement
// }

// func DeleteUserById(user_id string) {
// 	//TODO : implement
// }
