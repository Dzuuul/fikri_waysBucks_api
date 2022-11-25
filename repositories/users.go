package repositories

import (
	"ways-bucks-api/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User, ID int) (models.User, error)
	DeleteUser(user models.User, ID int) (models.User, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error // err := r.db.Raw("SELECT * FROM users").Scan(&users).Error

	return users, err
}

func (r *repository) GetUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error // err := r.db.Raw("SELECT * FROM users WHERE id=?", ID).Scan(&user).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error // 	err := r.db.Exec("INSERT INTO users(fullname,email,password,image) VALUES (?,?,?,?)", user.Fullname, user.Email, user.Password, user.Image).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User, ID int) (models.User, error) {
	err := r.db.Save(&user).Error // err := r.db.Raw("UPDATE users SET fullname=?, email=?, password=?, image=? WHERE id=?", user.Fullname, user.Email, user.Password, user.Image, ID).Scan(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User, ID int) (models.User, error) {
	err := r.db.Delete(&user).Error // err := r.db.Raw("DELETE FROM users WHERE id=?", ID).Scan(&user).Error

	return user, err
}
