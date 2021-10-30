package main

import (
	"database/sql"
	"fmt"
	"github.com/jgrath/go-and-find-with-go/config/impl"
	"github.com/jgrath/go-and-find-with-go/handlers"
	"github.com/jgrath/go-and-find-with-go/store"
	"github.com/jgrath/go-and-find-with-go/util"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func createRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/system-settings", handlers.FindSystemProperties).Methods("GET")
	router.HandleFunc("/system-settings/group-name/{search-keyword}", handlers.FindSystemPropertiesByCriteria).Methods("GET")
	router.HandleFunc("/system-settings/group-code/{search-keyword}", handlers.FindSystemPropertiesByCriteria).Methods("GET")
	router.HandleFunc("/system-settings", handlers.AddSystemProperties).Methods("POST")
	return router
}

func main() {

	util.LogInfo.Println("starting application")

	databaseConfiguration := impl.GetConfiguration().DatabaseConfiguration
	applicationConfiguration := impl.GetConfiguration().ApplicationConfiguration

	databaseConnectionString := "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	connString := fmt.Sprintf(databaseConnectionString, databaseConfiguration.Host, databaseConfiguration.Port,
		databaseConfiguration.User, databaseConfiguration.Password, databaseConfiguration.DBName,
		databaseConfiguration.SSLMode)

	databasePointer, err := sql.Open(databaseConfiguration.DBDriverName, connString)

	if err != nil {
		panic(err)
	}
	err = databasePointer.Ping()

	if err != nil {
		panic(err)
	}

	store.InitializePropertyStore(&store.MySQLPropertyStore{MainDatabase: databasePointer})

	r := createRouter()
	http.ListenAndServe(":"+applicationConfiguration.AppPort, r)
}
