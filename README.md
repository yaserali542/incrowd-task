# incrowd-task
---
This service runs both scheduler and RestAPI concurrently

## Scheduler

1. interval of the scheduler can be configured in `./config/config.json` and the key is `interval-in-seconds`
2. Interval has to be in seconds.
3. At every interval an external Rest API call is made to fetch the data
4. Number of records to fetch from external api is configured `./config/config.json` and the key is `htafc-record-count` (default is 50 if key is absent or set to 0 )
5. at every interval the records are transformed and stored in DB by `insertMany`
6. any duplicate document error is ignored by checking the error code
7. Ordered is set to false so that it continues to insert the other document.

---
## RestAPI

1. provides 2 endpoints on port `:8000`
   - http://localhost:8000/test/incrowd/v1/news/articles
   - http://localhost:8000/test/incrowd/v1/news/articles/{id}
2. 1st API gets all the records from the mongoDB
3. While fetching, the records are pulling in batches of 1000 to avoid read locks while writing, memory overload
4. 2nd API returns single result by providing id for example `http://localhost:8000/test/incrowd/v1/news/articles/608677`
5. If the id is not found in MongoDB, external call is made to get that record and returned after storing in DB
6. even then the record is not found 204 is returned indicating record not found.

---
## Steps to run this service 
1. In local machine
   - install Mongodb and provide the connection string in `/config/config.json` and the key is `mongodb-connection-string`
   - (optional) Change db name or collection from the same config if desired
   - (optional) change scheduler time as per requirement
   - run `go run main.go` in terminal by going to the repo root directory.
2. Using Docker
   - install docker and docker-compose
   -  run `docker-compose build`
   -  run `docker-compose up`
   -  try to access any URL from your local machine browser
