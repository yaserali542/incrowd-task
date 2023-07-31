// This package provides the functionality for the scheduler.
package schedulerservice

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/spf13/viper"
	httpclient "github.com/yaserali542/incrowd-task/http-client"
	"github.com/yaserali542/incrowd-task/models"
	"github.com/yaserali542/incrowd-task/store/mongodb"
)

// Struct to store the dependencies.
type SchedulerService struct {
	IntervalSeconds int
	HttpClient      httpclient.HtafcClienter
	MongodbClient   mongodb.MongoDbStorer
}

// Create SchedulerService object which contains injected dependencies.
func CreateSchedulerService(viper *viper.Viper, httpClient httpclient.HtafcClienter, mongodbClient mongodb.MongoDbStorer) *SchedulerService {

	interval := viper.GetInt("interval-in-seconds")
	if interval == 0 {
		interval = 3600
	}

	return &SchedulerService{
		IntervalSeconds: interval,
		HttpClient:      httpClient,
		MongodbClient:   mongodbClient,
	}

}

// Method starts cron job asynchronously.
func (service *SchedulerService) StartCronJob() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(service.IntervalSeconds).Seconds().Do(service.PopulateData)
	s.StartAsync()

}

// This method defines the test for the scheduler. Invoked at every interval.
func (service *SchedulerService) PopulateData() {
	fmt.Println("Cron Job has started at time :", time.Now())
	data, err := service.HttpClient.GetBulkArticles()
	if err != nil {
		fmt.Println("error occured while fetching from api: ", err)
		return
	}

	var records []models.Newsarticles

	for _, v := range data.NewsletterNewsItems.NewsletterNewsItem {
		layout := "2006-01-02 15:04:05"
		publishdate, err := time.Parse(layout, v.PublishDate)
		if err != nil {
			fmt.Println("Error parsing the date", v.PublishDate)
			publishdate = time.Now()
		}

		updatedDate, err := time.Parse(layout, v.LastUpdateDate)
		if err != nil {
			fmt.Println("Error parsing the date", v.LastUpdateDate)
			updatedDate = time.Now()
		}

		records = append(records, models.Newsarticles{ // convert the api response to mongodb response.
			ArticleURL:        v.ArticleURL,
			NewsArticleID:     v.NewsArticleID,
			PublishDate:       publishdate,
			Taxonomies:        v.Taxonomies,
			TeaserText:        v.TeaserText,
			ThumbnailImageURL: v.ThumbnailImageURL,
			Title:             v.Title,
			OptaMatchId:       v.OptaMatchId,
			LastUpdateDate:    updatedDate,
			IsPublished:       v.IsPublished,
			ClubName:          data.ClubName,
			ClubWebsiteURL:    data.ClubWebsiteURL,
		})
	}

	if err = service.MongodbClient.BulkInsert(records); err != nil {
		fmt.Println("error in inserting the data ", err)
	}

	fmt.Println("Task is completed at Time: ", time.Now())
}
