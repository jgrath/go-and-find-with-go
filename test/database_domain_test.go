package test

import (
	"database/sql"
	"fmt"
	store2 "github.com/jgrath/go-and-find-with-go/store"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PropertyStoreSuite struct {
	suite.Suite
	store *store2.MySQLPropertyStore
	_db   *sql.DB
}

func (testSuite *PropertyStoreSuite) SetupSuite() {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "postgres", "postgres", "postgres")

	db, err := sql.Open("postgres", connString)

	if err != nil {
		testSuite.T().Fatal(err)
	}

	testSuite._db = db
	testSuite.store = &store2.MySQLPropertyStore{MainDatabase: db}
}

func (testSuite *PropertyStoreSuite) SetupTest() {
	_, err := testSuite._db.Query("DELETE FROM public.\"SYSTEM_SETTINGS\";")
	if err != nil {
		testSuite.T().Fatal(err)
	}
}

func (testSuite *PropertyStoreSuite) TearDownSuite() {
	testSuite._db.Close()
}

func TestPropertyStoreSuite(t *testing.T) {
	s := new(PropertyStoreSuite)
	suite.Run(t, s)
}

func (testSuite *PropertyStoreSuite) TestFindSystemProperties() {
	_, err := testSuite._db.Query(`INSERT INTO public."SYSTEM_SETTINGS" (PROPERTY_NAME, PROPERTY_VALUE, default_value, description, data_type, enabled, active_from_date, group_code) 
		VALUES('name-1','value-1', 'default-value', 'description', 'char', 1, current_timestamp, 'main_group')`)

	if err != nil {
		testSuite.T().Fatal(err)
	}

	propertySettings, err := testSuite.store.FindProperties()

	if err != nil {
		testSuite.T().Fatal(err)
	}

	numberOfPropertySettings := len(propertySettings)

	if numberOfPropertySettings != 1 {
		testSuite.T().Errorf("incorrect count, wanted 1, got %d", numberOfPropertySettings)
	}

	if propertySettings[0].Name != "name-1" {
		testSuite.T().Errorf("incorrect details, expected %v, got %v", "name-1", propertySettings[0].Name)
	}
}

