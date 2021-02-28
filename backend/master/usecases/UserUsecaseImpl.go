package usecases

import (
	"testMekar/master/models"
	"testMekar/master/repositories"
)

type UserUsecaseImpl struct{
	userRepo repositories.UserRepo
}

func (u UserUsecaseImpl) GetUsers() ([]*models.DetailUser, error) {
	users, err := u.userRepo.GetAllUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u UserUsecaseImpl) GetUserById(id int) (*models.DetailUser, error) {
	user, err := u.userRepo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserUsecaseImpl) PostUser(user models.User) (*models.DetailUser, error) {
	err := u.userRepo.AddUser(user)
	if err != nil {
		return nil, err
	}
	response, err := u.userRepo.GetUserById(user.IdUser)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u UserUsecaseImpl) PutUser(id int, user models.User) (*models.DetailUser, error) {
	err := u.userRepo.UpdateUser(id, user)
	if err != nil {
		return nil, err
	}
	response, err := u.userRepo.GetUserById(user.IdUser)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (u UserUsecaseImpl) DeleteUser(id int) error {
	err := u.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}

func (u UserUsecaseImpl) GetJob() ([]*models.Pekerjaan, error) {
	listJob, err := u.userRepo.ReadJob()
	if err != nil {
		return nil, err
	}
	return listJob, nil
}

func (u UserUsecaseImpl) GetEducation() ([]*models.Pendidikan, error) {
	listEdu, err := u.userRepo.ReadEducation()
	if err != nil {
		return nil, err
	}
	return listEdu, nil
}

func InitUserUsecase(userRepo repositories.UserRepo) UserUsecase {
	return &UserUsecaseImpl{userRepo}
}

