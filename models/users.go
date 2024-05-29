package models

type User struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Sex     string `json:"sex"`
}

func (u *User) GetAllUsers() ([]User, error) {
	var users []User
	result := db.Find(&users)
	return users, result.Error
}

func (u *User) GetUserByID(id uint) (User, error) {
	result := db.First(&u, id)
	return *u, result.Error
}

func (u *User) CreateUser() (User, error) {
	result := db.Create(&u)
	return *u, result.Error
}

func (u *User) UpdateUser(id uint) error {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	user.Name = u.Name
	user.Address = u.Address
	user.Sex = u.Sex
	return db.Save(&user).Error
}

func (u *User) DeleteUser(id uint) error {
	return db.Delete(&u, id).Error
}
