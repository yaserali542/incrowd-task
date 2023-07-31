package httpclient

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/spf13/viper"
	"github.com/yaserali542/incrowd-task/models"
)

// Contains all the dependency required for making calls.
// For Example FetchByIdURL, fetAllDataURL
type HtafcClient struct {
	fetchAllDetailURL string
	fetchbyIdURL      string
	count             int
}

// Create a struct object by taking the values from the config file
func CreateHtafcClient(viper *viper.Viper) *HtafcClient {

	fetchAllurlString := viper.GetString("htafc-all-record-url") //get from config
	if fetchAllurlString == "" {                                 // set default value if config is not set
		fetchAllurlString = "https://www.htafc.com/api/incrowd/getnewlistinformation?count=50"
	}

	fetchByIdurlString := viper.GetString("htafc-by-id-record-url")
	if fetchByIdurlString == "" {
		fetchByIdurlString = "https://www.htafc.com/api/incrowd/getnewsarticleinformation?id=XXXX"
	}

	count := viper.GetInt("htafc-record-count") // get the count number
	if count == 0 {
		count = 50 //set the defualt value
	}

	return &HtafcClient{
		fetchAllDetailURL: fetchAllurlString,
		count:             count,
		fetchbyIdURL:      fetchByIdurlString,
	}
}

// All the news Articles from the external API in bulk. Count is configurable in the config.
// Decodes the data as the stream rather than consuming all at once.
func (client *HtafcClient) GetBulkArticles() (*models.NewListInformation, error) {

	u, err := url.Parse(client.fetchAllDetailURL) //parse the url to change the count.
	if err != nil {
		panic(err)
	}
	q := u.Query()
	q.Set("count", strconv.Itoa(client.count))
	u.RawQuery = q.Encode()

	res, err := MakeHttpCall(http.MethodGet, nil, u.String())

	if err != nil {
		return nil, err
	}

	data := models.NewListInformation{}
	//stream decode rather than unmarshal.
	if err := xml.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}

// Get the specific article by providing the Id
func (client *HtafcClient) GetArticleById(id string) (*models.HtafcNewsArticleInformation, bool, error) {

	u, err := url.Parse(client.fetchbyIdURL)
	if err != nil {
		panic(err)
	}
	q := u.Query()
	q.Set("id", id) // set the id
	u.RawQuery = q.Encode()
	fmt.Println("url to call single api : ", u.String())

	res, err := MakeHttpCall(http.MethodGet, nil, u.String())

	if err != nil {
		return nil, false, err
	}
	data := models.HtafcNewsArticleInformation{}
	//Stream decode.
	if err := xml.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, false, err
	}
	//Illegal argument also gives valid data. Adding additional check to ensure id is present in the response.
	if data.NewsArticle.NewsArticleID == "" {
		return nil, true, errors.New("No Record found")
	}

	return &data, false, nil
}
