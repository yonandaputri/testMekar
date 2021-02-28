package repositories

import (
	"database/sql"
	"testMekar/master/models"
	"testMekar/utils"
)

type UserRepoImpl struct{
	db *sql.DB
}

func (u UserRepoImpl) GetAllUser() ([]*models.DetailUser, error) {
	var users []*models.DetailUser
	data, err := u.db.Query(utils.READ_USER)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		var user = new(models.DetailUser)
		var err = data.Scan(&user.IdUser, &user.NamaUser, &user.TanggalLahir, &user.NoKtp, &user.NamaPekerjaan, &user.PendidikanTerakhir)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = data.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u UserRepoImpl) GetUserById(id int) (*models.DetailUser, error) {
	var user = new(models.DetailUser)
	err := u.db.QueryRow(utils.READ_USER_BY_ID, id).Scan(&user.IdUser, &user.NamaUser, &user.TanggalLahir, &user.NoKtp, &user.NamaPekerjaan, &user.PendidikanTerakhir)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserRepoImpl) AddUser(user models.User) error {
	data, err := u.db.Begin()

	if err != nil {
		return err
	}

	row, err := u.db.Prepare(utils.CREATE_USER)

	if err != nil {
		return err
	}

	_, err = row.Exec(&user.NamaUser, &user.TanggalLahir, &user.NoKtp, &user.IdPekerjaan, &user.IdPendidikan)
	if err != nil {
		data.Rollback()
		return err
	}

	err = data.Commit()
	if err != nil {
		return err
	}
	row.Close()
	return nil
}

func (u UserRepoImpl) UpdateUser(id int, user models.User) error {
	data, err := u.db.Begin()
	if err != nil {
		return err
	}

	row, err := u.db.Prepare(utils.UPDATE_USER)
	if err != nil {
		return err
	}

	_, err = row.Exec(&user.NamaUser, &user.TanggalLahir, &user.NoKtp, &user.IdPekerjaan, &user.IdPendidikan, id)
	if err != nil {
		data.Rollback()
		return err
	}

	err = data.Commit()
	if err != nil {
		return err
	}
	row.Close()
	return nil
}

func (u UserRepoImpl) DeleteUser(id int) error {
	data, err := u.db.Begin()
	if err != nil {
		return err
	}

	row, err := u.db.Prepare(utils.DELETE_USER)
	if err != nil {
		return err
	}

	_, err = row.Exec(id)
	if err != nil {
		data.Rollback()
		return err
	}

	err = data.Commit()
	if err != nil {
		return err
	}
	row.Close()
	return nil
}

func (u UserRepoImpl) ReadJob() ([]*models.Pekerjaan, error) {
	var jobs []*models.Pekerjaan
	data, err := u.db.Query(utils.READ_PEKERJAAN)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		var job = new(models.Pekerjaan)
		var err = data.Scan(&job.IdPekerjaan, &job.NamaPekerjaan)

		if err != nil {
			return nil, err
		}

		jobs = append(jobs, job)
	}

	if err = data.Err(); err != nil {
		return nil, err
	}

	return jobs, nil
}

func (u UserRepoImpl) ReadEducation() ([]*models.Pendidikan, error) {
	var educations []*models.Pendidikan
	data, err := u.db.Query(utils.READ_PENDIDIKAN)
	if err != nil {
		return nil, err
	}
	defer data.Close()

	for data.Next() {
		var education = new(models.Pendidikan)
		var err = data.Scan(&education.IdPendidikan, &education.LabelPendidikan)

		if err != nil {
			return nil, err
		}

		educations = append(educations, education)
	}

	if err = data.Err(); err != nil {
		return nil, err
	}

	return educations, nil
}

func InitUserRepoImpl(db *sql.DB) UserRepo {
	return &UserRepoImpl{db}
}

