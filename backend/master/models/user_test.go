package models

import "github.com/stretchr/testify/suite"

type SuiteTestUser struct {
	suite.Suite
	VariableIdUser       int
	VariableNamaUser     string
	VariableTanggalLahir string
	VariableNoKtp        string
	VariableIdPekerjaan  int
	VariableIdPendidikan int
}

type SuiteTestDetailUser struct {
	suite.Suite
	VariableIdUser             int
	VariableNamaUser           string
	VariableTanggalLahir       string
	VariableNoKtp              string
	VariableNamaPekerjaan      string
	VariablePendidikanTerakhir string
}
