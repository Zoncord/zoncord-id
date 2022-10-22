package models

import (
	"github.com/DATA-DOG/go-sqlmock"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository Repository
	person     *model.Person
}

func (s *Suite) SetupSuite() {
	var (
		db  *gorm.DB
		err error
	)
	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)
	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)
	s.repository = NewRepository(s.DB)
}
