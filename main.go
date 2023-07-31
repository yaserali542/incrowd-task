package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/yaserali542/incrowd-task/controllers"
	httpclient "github.com/yaserali542/incrowd-task/http-client"
	apiservice "github.com/yaserali542/incrowd-task/services/api-service"
	schedulerservice "github.com/yaserali542/incrowd-task/services/scheduler-service"
	"github.com/yaserali542/incrowd-task/store/mongodb"
)

func main() {

	viper, err := LoadConfig() // load the configuration from the config/config.json
	if err != nil {
		panic(err)
	}

	var apiservicer apiservice.ApiServicer     // interface
	var htafcClienter httpclient.HtafcClienter // interface
	var mongodbClienter mongodb.MongoDbStorer  // interface

	htafcClienter = httpclient.CreateHtafcClient(viper)   // create the struct object to resolve dependency
	mongodbClienter = mongodb.ConnectMongoDbClient(viper) // create the struct object to resolve dependency
	defer mongodbClienter.DisconnectMongoDBClient()       // create the struct object to resolve dependency

	schedulerService := schedulerservice.CreateSchedulerService(viper, htafcClienter, mongodbClienter) //pass the dependencies to the scheduler service
	schedulerService.StartCronJob()                                                                    // schedule and start cron job async

	apiservicer = apiservice.CreateNewsInformationService(htafcClienter, mongodbClienter) //pass the dependencies to the api service

	apicontroller := controllers.CreateNewsInformationController(apiservicer) //pass the dependencies to the api controller

	r := mux.NewRouter()                              //create a mux router
	s := r.PathPrefix("/test/incrowd/v1").Subrouter() // prefix with path

	s.HandleFunc("/news/articles", apicontroller.FetchAllRecords).Methods("GET") // register the endpoint
	s.HandleFunc("/news/articles/{id}", apicontroller.FetchById).Methods("GET")  // register the endpoint

	log.Fatal(http.ListenAndServe(":8000", r))

}

// This method set the configuration path, load and return viper instance.
func LoadConfig() (*viper.Viper, error) {
	viper.SetConfigType("json")
	viper.SetConfigFile("./config/config.json")

	viper.AutomaticEnv()

	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	return viper.GetViper(), nil
}
