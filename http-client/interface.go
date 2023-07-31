package httpclient

import "github.com/yaserali542/incrowd-task/models"

// Interface to Inject the dependency at the runtime. Helps in writing the unit tests.
type HtafcClienter interface {
	GetBulkArticles() (*models.NewListInformation, error)
	GetArticleById(string) (*models.HtafcNewsArticleInformation, bool, error)
}
