package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AdServiceTestSuite struct {
	suite.Suite
	db *sql.DB
}

func (s *AdServiceTestSuite) SetupTest() {
	s.db = connectToDB()
}

func (s *AdServiceTestSuite) TearDownTest() {
	s.db.Close()
}

func (s *AdServiceTestSuite) TestCalculateDiscount() {
	res := Calculate(s.db, 100)

	s.NoError(nil)
}

func TestRunSuite(t *testing.T)  {
	suite.Run(t, new(AdServiceTestSuite))
}
