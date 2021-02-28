package utils

const (
	READ_USER = `SELECT * FROM user_detail`
	READ_USER_BY_ID = `SELECT * FROM user_detail WHERE id_user = ?`
	CREATE_USER = `INSERT INTO user(nama_user, tanggal_lahir, no_ktp, id_pekerjaan, id_pendidikan) VALUES (?, ?, ?, ?, ?)`
	UPDATE_USER = `UPDATE user SET nama_user = ?, tanggal_lahir = ?, no_ktp = ?, id_pekerjaan = ?, id_pendidikan = ? WHERE id_user = ?`
	DELETE_USER = `DELETE FROM user WHERE id_user = ?`
	READ_PEKERJAAN = `SELECT * FROM pekerjaan`
	READ_PENDIDIKAN = `SELECT * FROM pendidikan`
)

