//This package handles all the requests for the api . It responsibility is to  extracting data for services packages.
//Also the response from the service is mapped to the response of the API.

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/yaserali542/incrowd-task/models"
	apiservice "github.com/yaserali542/incrowd-task/services/api-service"
)

// This struct have all the Depencencies required for the controllers.
type NewsInformationController struct {
	service apiservice.ApiServicer
}

// Response struct for FetchAllData API
type fetchAllApiResponse struct {
	Status   string                `json:"status"`
	Data     []models.Newsarticles `json:"data"`
	Metadata metadataFetchAllApi   `json:"metadata"`
}

// MetaData struct for FetchAllData API
type metadataFetchAllApi struct {
	CreatedAt  time.Time `json:"createdAt"`
	TotalItems int       `json:"totalItems"`
	Sort       string    `json:"sort"`
}

// Response struct for FetchById API
type fetchByIdResponse struct {
	Status   string               `json:"status"`
	Data     models.Newsarticles  `json:"data"`
	Metadata metadataFetchByIdApi `json:"metadata"`
}

// MetaData struct for FetchById API
type metadataFetchByIdApi struct {
	CreatedAt time.Time `json:"createdAt"`
}

// contructor to initialize the Struct and inject the dependency.This design is followed for granular Unit testing rather that whole integeration testing
func CreateNewsInformationController(s apiservice.ApiServicer) *NewsInformationController {
	return &NewsInformationController{
		service: s,
	}
}

// Http Handler for FetchAllRecords
func (controller *NewsInformationController) FetchAllRecords(w http.ResponseWriter, r *http.Request) {

	data, err := controller.service.GetAllArticles() //Forwards the call to

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error Occured Details being logged", http.StatusInternalServerError) //Do not display the error at the response
	}
	response := fetchAllApiResponse{ //convert the data to api response
		Status: "success",
		Data:   data,
		Metadata: metadataFetchAllApi{
			CreatedAt:  time.Now(),
			TotalItems: len(data),
			Sort:       "-publishedAt",
		},
	}
	w.Header().Set("Content-Type", "application/json") // set the header to application/json
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// http handler for fetch the details by id
func (controller *NewsInformationController) FetchById(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]                                 // get the id from the url
	data, notfound, err := controller.service.FetchById(id) //call service to get the data

	if notfound {
		http.Error(w, "content not found", http.StatusNoContent)
		return
	}

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error Occured Details being logged", http.StatusInternalServerError)
		return
	}
	response := fetchByIdResponse{ // convert to api response
		Status: "success",
		Data:   *data,
		Metadata: metadataFetchByIdApi{
			CreatedAt: time.Now(),
		},
	}
	w.Header().Set("Content-Type", "application/json") // set the header to application/json
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)

}
