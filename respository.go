package main

func GetUserByEmail(email string) (User, error) {
	var user User
	if err := DB.Where("email = ?", email).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func createUser(user *User) error {
	if err := DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(id string) (User, error) {
	var user User
	if err := DB.Where("id = ?", id).First(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
