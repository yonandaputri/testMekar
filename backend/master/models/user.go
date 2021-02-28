package models

type User struct {
	IdUser       int    `json:"id_user"`
	NamaUser     string `json:"nama_user"`
	TanggalLahir string `json:"tanggal_lahir"`
	NoKtp        string `json:"no_ktp"`
	IdPekerjaan  int    `json:"id_pekerjaan"`
	IdPendidikan int    `json:"id_pendidikan"`
}

type DetailUser struct {
	IdUser             int    `json:"id_user"`
	NamaUser           string `json:"nama_user"`
	TanggalLahir       string `json:"tanggal_lahir"`
	NoKtp              string `json:"no_ktp"`
	NamaPekerjaan      string `json:"nama_pekerjaan"`
	PendidikanTerakhir string `json:"pendidikan_terakhir"`
}
