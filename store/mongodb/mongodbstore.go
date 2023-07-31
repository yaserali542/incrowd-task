package mongodb

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"github.com/yaserali542/incrowd-task/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Struct to store the dependencies for MongoDB
type MongoDbStore struct {
	Client     *mongo.Client
	Collection *mongo.Collection
}

// method to create MongoDbStore object. which contains resolved dependencies.
func ConnectMongoDbClient(viper *viper.Viper) *MongoDbStore {
	connectionString := viper.GetString("mongodb-connection-string")
	d := viper.GetString("mongodb-database")
	c := viper.GetString("mongodb-collection")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString)) //connect to mongo DB

	if err != nil {
		panic(err)
	}
	collection := client.Database(d).Collection(c) // create the database and collection
	err = client.Ping(ctx, readpref.Primary())     //ping to verify the connectivity
	if err != nil {
		fmt.Println("error in ping", err)
	}
	return &MongoDbStore{
		Client:     client,
		Collection: collection,
	}
}

// method to disconnect the mongodb.
// prefereed to use as defer in the main function
func (client *MongoDbStore) DisconnectMongoDBClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

// method to insert the data in bulk
func (client *MongoDbStore) BulkInsert(data []models.Newsarticles) error {
	s := make([]interface{}, len(data))
	for i, v := range data { //convert the data of type  []models.Newsarticles to []interface{}
		s[i] = v
	}
	insertManyOptions := options.InsertManyOptions{} // options for the query
	insertManyOptions.SetOrdered(false)              // this is allow to inset if any error occurs.
	ctx := context.TODO()
	_, err := client.Collection.InsertMany(ctx, s, &insertManyOptions)
	if err != nil {
		bulkerr := new(mongo.BulkWriteException)
		if errors.As(err, bulkerr) { // converts generic type of error to BulkWriteException error
			for _, v := range bulkerr.WriteErrors {
				if v.Code != 11000 { // ignore the duplicate error with code 11000
					return err
				}
			}
		}
	}
	return nil
}

// insert singel record of data.
func (client *MongoDbStore) InsertSingle(data *models.Newsarticles) error {
	var s interface{}
	s = data
	ctx := context.TODO()
	_, err := client.Collection.InsertOne(ctx, s, nil)
	if err != nil {
		return err
	}
	return nil
}

// this method fetches all the records in batches of 10000
// This is to prevent out of memory and heavy read lock wait time.
func (client *MongoDbStore) FindAllRecords() ([]models.Newsarticles, error) {
	ctx := context.Background()
	cur, err := client.Collection.Find(ctx, bson.D{},
		options.Find().SetBatchSize(10000).SetNoCursorTimeout(true).SetSort(bson.D{{Key: "publishDate", Value: -1}}))

	if err != nil {
		return nil, err
	}

	var data []models.Newsarticles

	for cur.Next(ctx) {
		var result models.Newsarticles
		err := cur.Decode(&result)
		if err != nil {

			fmt.Println("error while decode the cursor", err)
		}
		data = append(data, result)

	}
	return data, nil
}

// method finds the document by id.
func (client *MongoDbStore) FindById(id string) (*models.Newsarticles, error) {
	var data models.Newsarticles
	err := client.Collection.FindOne(context.Background(), bson.M{"_id": id}, nil).Decode(&data)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		//if err == mongo.ErrNoDocuments {
		return nil, err
	}
	return &data, nil
}
