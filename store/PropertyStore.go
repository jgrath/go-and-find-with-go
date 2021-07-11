package store

import (
	"database/sql"
	"errors"
	"fmt"
	. "github.com/jgrath/go-and-find-with-go/util"
)

var MainPropertyStore PropertyStore

type PropertyStore interface {
	FindProperties() ([]*SystemProperty, error)
	AddProperty(*SystemProperty)(error)
}

type MySQLPropertyStore struct {
	MainDatabase *sql.DB
}

func (store *MySQLPropertyStore) FindProperties() ([]*SystemProperty, error) {

	rows, err := store.MainDatabase.Query(propertyRequestSQL)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	listOfSystemProperties := []*SystemProperty{}

	for rows.Next() {
		property := &SystemProperty{}
		if err := rows.Scan(&property.Name, &property.Value, &property.DefaultValue,
			&property.Description, &property.DataType, &property.Enabled, &property.ActiveFromDate,
			&property.GroupCode, &property.GroupName, &property.GroupDescription);
		err != nil {
			return nil, err
		}
		listOfSystemProperties = append(listOfSystemProperties, property)
	}
	return listOfSystemProperties, nil
}

func (store *MySQLPropertyStore) AddProperty(property *SystemProperty) (error) {

	insertSQL := propertyInsertSQL

	_, err := store.MainDatabase.Exec(insertSQL, &property.Name, &property.Value, &property.DefaultValue,
		&property.Description, &property.DataType, 1, &property.ActiveFromDate, &property.GroupCode)

	if err != nil {
		var message = fmt.Sprintf("error in executing sql: [%s]", insertSQL)
		LogError.Println(message)
		errors.New(message)
	}

	return err
}

func InitializePropertyStore(store PropertyStore) {
	LogInfo.Println("initialising property store...")
	MainPropertyStore = store
	LogInfo.Println("initialised successful")
}
