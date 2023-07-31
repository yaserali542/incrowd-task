//This package provides the functionalilty to operate on mongoDB like insert, delete. 
package mongodb

import "github.com/yaserali542/incrowd-task/models"
//Interface to Inject the dependency at the runtime. Helps in writing the unit tests. 
type MongoDbStorer interface {
	DisconnectMongoDBClient()
	BulkInsert([]models.Newsarticles) error
	InsertSingle(*models.Newsarticles) error
	FindAllRecords() ([]models.Newsarticles, error)
	FindById(id string) (*models.Newsarticles, error)
}
