package usecases

import "testMekar/master/models"

type UserUsecase interface {
	GetUsers() ([]*models.DetailUser, error)
	GetUserById(id int) (*models.DetailUser, error)
	PostUser(user models.User) (*models.DetailUser, error)
	PutUser(id int, user models.User) (*models.DetailUser, error)
	DeleteUser(id int) error
	GetJob() ([]*models.Pekerjaan, error)
	GetEducation() ([]*models.Pendidikan, error)
}