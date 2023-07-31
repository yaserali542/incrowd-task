package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/yaserali542/incrowd-task/models"
	apiservice "github.com/yaserali542/incrowd-task/services/api-service"
)

type testPositiveService struct {
}

func (service *testPositiveService) GetAllArticles() ([]models.Newsarticles, error) {

	return []models.Newsarticles{
		{
			ArticleURL:        "1",
			NewsArticleID:     "1",
			PublishDate:       time.Now(),
			Taxonomies:        "1",
			TeaserText:        "1",
			ThumbnailImageURL: "1",
			Title:             "1",
			OptaMatchId:       "1",
			LastUpdateDate:    time.Now(),
			IsPublished:       true,
			ClubName:          "a",
			ClubWebsiteURL:    "1",
		},
		{
			ArticleURL:        "2",
			NewsArticleID:     "2",
			PublishDate:       time.Now(),
			Taxonomies:        "2",
			TeaserText:        "2",
			ThumbnailImageURL: "2",
			Title:             "2",
			OptaMatchId:       "2",
			LastUpdateDate:    time.Now(),
			IsPublished:       true,
			ClubName:          "b",
			ClubWebsiteURL:    "2",
		},
	}, nil
}

func (service *testPositiveService) FetchById(string) (*models.Newsarticles, bool, error) {
	return &models.Newsarticles{
		ArticleURL:        "1",
		NewsArticleID:     "1",
		PublishDate:       time.Now(),
		Taxonomies:        "1",
		TeaserText:        "1",
		ThumbnailImageURL: "1",
		Title:             "1",
		OptaMatchId:       "1",
		LastUpdateDate:    time.Now(),
		IsPublished:       true,
		ClubName:          "a",
		ClubWebsiteURL:    "1",
	}, false, nil
}

// Unit test of the controller for fetching all the records.
func TestFetchAllContentHandler(t *testing.T) {

	var apiservicer apiservice.ApiServicer
	apiservicer = &testPositiveService{}
	apicontroller := CreateNewsInformationController(apiservicer)

	req, err := http.NewRequest("GET", "/test/incrowd/v1/news/articles", nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := http.HandlerFunc(apicontroller.FetchAllRecords)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK && rr.Header().Get("Content-Type") == "application/json" {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestFetByIdController(t *testing.T) {
	var apiservicer apiservice.ApiServicer
	apiservicer = &testPositiveService{}
	apicontroller := CreateNewsInformationController(apiservicer)

	req, err := http.NewRequest("GET", "/test/incrowd/v1/news/articles/", nil)
	if err != nil {
		t.Fatal(err)
	}
	handler := http.HandlerFunc(apicontroller.FetchById)

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK && rr.Header().Get("Content-Type") == "application/json" {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
