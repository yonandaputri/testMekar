package repositories

import "testMekar/master/models"

type UserRepo interface {
	GetAllUser() ([]*models.DetailUser, error)
	GetUserById(id int) (*models.DetailUser, error)
	AddUser(user models.User) error
	UpdateUser(id int, user models.User) error
	DeleteUser(id int) error
	ReadJob() ([]*models.Pekerjaan, error)
	ReadEducation() ([]*models.Pendidikan, error)
}
