package apiservice

import (
	"testing"
	"time"

	httpclient "github.com/yaserali542/incrowd-task/http-client"
	"github.com/yaserali542/incrowd-task/models"
	"github.com/yaserali542/incrowd-task/store/mongodb"
)

func TestGetAllArticles(t *testing.T) {
	var htafcClienter httpclient.HtafcClienter // interface
	var mongodbClienter mongodb.MongoDbStorer  // interface

	htafcClienter = &HtafcClient{}
	mongodbClienter = &mongoDbStore{}
	s := CreateNewsInformationService(htafcClienter, mongodbClienter)

	data, err := s.GetAllArticles()

	if err != nil {
		t.Errorf("service returned error : %v", err)
	}
	if len(data) != 2 {
		t.Errorf("service returned wrong len got %v want %v", len(data), 2)
	}

}

func TestFetchById(t *testing.T) {
	var htafcClienter httpclient.HtafcClienter // interface
	var mongodbClienter mongodb.MongoDbStorer  // interface

	htafcClienter = &HtafcClient{}
	mongodbClienter = &mongoDbStore{}
	s := CreateNewsInformationService(htafcClienter, mongodbClienter)

	data, notfound, err := s.FetchById("abc")

	if notfound {
		t.Errorf("service returned not found  : %v", notfound)
	}

	if err != nil {
		t.Errorf("service returned error : %v", err)
	}
	if data.ArticleURL != "1" {
		t.Errorf("service returned wrong data %v want 1", data.ArticleURL)
	}
}

type HtafcClient struct {
}

func (client *HtafcClient) GetBulkArticles() (*models.NewListInformation, error) {

	list := []models.NewsletterNewsItem{

		{
			ArticleURL:        "1",
			NewsArticleID:     "1",
			PublishDate:       "2023-07-30 13:49:29",
			Taxonomies:        "1",
			TeaserText:        "1",
			ThumbnailImageURL: "1",
			Title:             "1",
			OptaMatchId:       "1",
			LastUpdateDate:    "2023-07-30 13:49:29",
			IsPublished:       true,
		},
		{
			ArticleURL:        "2",
			NewsArticleID:     "2",
			PublishDate:       "2023-07-30 13:49:29",
			Taxonomies:        "2",
			TeaserText:        "2",
			ThumbnailImageURL: "2",
			Title:             "2",
			OptaMatchId:       "2",
			LastUpdateDate:    "2023-07-30 13:49:29",
			IsPublished:       true,
		},
	}

	items := models.NewsletterNewsItems{
		NewsletterNewsItem: list,
	}

	return &models.NewListInformation{
		ClubName:            "1",
		ClubWebsiteURL:      "1",
		NewsletterNewsItems: items,
	}, nil
}
func (client *HtafcClient) GetArticleById(string) (*models.HtafcNewsArticleInformation, bool, error) {

	a := models.HtafcNewsArticle{
		ArticleURL:        "2",
		NewsArticleID:     "2",
		PublishDate:       "2023-07-30 13:49:29",
		Taxonomies:        "2",
		TeaserText:        "2",
		ThumbnailImageURL: "2",
		Title:             "2",
		OptaMatchId:       "2",
		LastUpdateDate:    "2023-07-30 13:49:29",
		IsPublished:       true,
	}

	return &models.HtafcNewsArticleInformation{
		ClubName:       "a",
		ClubWebsiteURL: "1",
		NewsArticle:    a,
	}, false, nil
}

type mongoDbStore struct{}

func (store *mongoDbStore) DisconnectMongoDBClient() {}
func (store *mongoDbStore) BulkInsert([]models.Newsarticles) error {
	return nil
}
func (store *mongoDbStore) InsertSingle(*models.Newsarticles) error {
	return nil
}
func (store *mongoDbStore) FindAllRecords() ([]models.Newsarticles, error) {
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

func (store *mongoDbStore) FindById(id string) (*models.Newsarticles, error) {
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
	}, nil
}
