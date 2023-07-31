// This package provides the business logic on how the individua requests are handled.
package apiservice

import "github.com/yaserali542/incrowd-task/models"

// Interface to Inject the dependency at the runtime. Helps in writing the unit tests.
type ApiServicer interface {
	GetAllArticles() ([]models.Newsarticles, error)
	FetchById(string) (*models.Newsarticles, bool, error)
}
