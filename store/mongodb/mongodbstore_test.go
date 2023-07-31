package mongodb

import (
	"testing"
	"time"

	"github.com/spf13/viper"
	"github.com/yaserali542/incrowd-task/models"
)

func TestBulkInsert(t *testing.T) {
	data := []models.Newsarticles{
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
	}

	viper.SetDefault("mongodb-connection-string", "mongodb://root:root@localhost:27017")
	viper.SetDefault("mongodb-database", "unitTest")
	viper.SetDefault("mongodb-collection", "newsFeed")
	v := viper.GetViper()

	db := ConnectMongoDbClient(v)
	defer db.DisconnectMongoDBClient()

	if err := db.BulkInsert(data); err != nil {
		t.Error("Error occured while inserting ", err)
	}

}

func TestInsertSingle(t *testing.T) {

	data := &models.Newsarticles{
		ArticleURL:        "3",
		NewsArticleID:     "3",
		PublishDate:       time.Now(),
		Taxonomies:        "3",
		TeaserText:        "3",
		ThumbnailImageURL: "3",
		Title:             "3",
		OptaMatchId:       "3",
		LastUpdateDate:    time.Now(),
		IsPublished:       true,
		ClubName:          "3",
		ClubWebsiteURL:    "3",
	}
	viper.SetDefault("mongodb-connection-string", "mongodb://root:root@localhost:27017")
	viper.SetDefault("mongodb-database", "unitTest")
	viper.SetDefault("mongodb-collection", "newsFeed")
	v := viper.GetViper()

	db := ConnectMongoDbClient(v)
	defer db.DisconnectMongoDBClient()

	if err := db.InsertSingle(data); err != nil {
		t.Error("Error occured while inserting ", err)
	}
}

// expected to run TestInsertSingle and TestBulkInsert
func TestFetchAllData(t *testing.T) {

	viper.SetDefault("mongodb-connection-string", "mongodb://root:root@localhost:27017")
	viper.SetDefault("mongodb-database", "unitTest")
	viper.SetDefault("mongodb-collection", "newsFeed")
	v := viper.GetViper()

	db := ConnectMongoDbClient(v)
	defer db.DisconnectMongoDBClient()

	data, err := db.FindAllRecords()
	if err != nil {
		t.Error("Error occured while inserting ", err)
	}
	if len(data) != 3 {
		t.Errorf("db returned wrong len got %v want %v", len(data), 3)
	}
}
