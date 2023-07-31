package apiservice

import (
	"fmt"
	"time"

	httpclient "github.com/yaserali542/incrowd-task/http-client"
	"github.com/yaserali542/incrowd-task/models"
	"github.com/yaserali542/incrowd-task/store/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

// Struct to store all the dependency members
type NewsInformationService struct {
	HttpClient    httpclient.HtafcClienter
	MongodbClient mongodb.MongoDbStorer
}

// Creates a struct object
func CreateNewsInformationService(httpClient httpclient.HtafcClienter, mongodbClient mongodb.MongoDbStorer) *NewsInformationService {
	return &NewsInformationService{
		HttpClient:    httpClient,
		MongodbClient: mongodbClient,
	}
}

// Business logic to get all the records.
func (service *NewsInformationService) GetAllArticles() ([]models.Newsarticles, error) {
	return service.MongodbClient.FindAllRecords()
}

// Service method to fetch and return the document by ID
// If the data is not found in the DB, an external call to api is make to check the data
// If present the data is stored inDB and returned to controller.
func (service *NewsInformationService) FetchById(id string) (*models.Newsarticles, bool, error) {
	data, err := service.MongodbClient.FindById(id)
	if err != nil {
		if err != mongo.ErrNoDocuments { // if no document is found in the DB makes external call to get the data
			return nil, false, err
		}
		apiResponse, notfound, err := service.HttpClient.GetArticleById(id)
		if err != nil || notfound {
			return nil, notfound, err
		}

		layout := "2006-01-02 15:04:05"
		publishdate, err := time.Parse(layout, apiResponse.NewsArticle.PublishDate)
		if err != nil {
			fmt.Println("Error parsing the date", apiResponse.NewsArticle.PublishDate)
			publishdate = time.Now()
		}

		updatedDate, err := time.Parse(layout, apiResponse.NewsArticle.LastUpdateDate)
		if err != nil {
			fmt.Println("Error parsing the date", apiResponse.NewsArticle.LastUpdateDate)
			updatedDate = time.Now()
		}

		data = &models.Newsarticles{ //convert API model to mongodb model.
			ArticleURL:        apiResponse.NewsArticle.ArticleURL,
			NewsArticleID:     apiResponse.NewsArticle.NewsArticleID,
			PublishDate:       publishdate,
			Taxonomies:        apiResponse.NewsArticle.Taxonomies,
			TeaserText:        apiResponse.NewsArticle.TeaserText,
			ThumbnailImageURL: apiResponse.NewsArticle.ThumbnailImageURL,
			Title:             apiResponse.NewsArticle.Title,
			OptaMatchId:       apiResponse.NewsArticle.OptaMatchId,
			LastUpdateDate:    updatedDate,
			IsPublished:       apiResponse.NewsArticle.IsPublished,
			ClubName:          apiResponse.ClubName,
			ClubWebsiteURL:    apiResponse.ClubWebsiteURL,
		}

		_ = service.MongodbClient.InsertSingle(data)

	}
	return data, false, nil

}
