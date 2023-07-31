package httpclient

import (
	"testing"

	"github.com/spf13/viper"
)

func TestClient(t *testing.T) {
	v := viper.GetViper()

	client := CreateHtafcClient(v)

	_, err := client.GetBulkArticles()

	if err != nil {
		t.Error("error while fetching all records :", err)
	}

	_, notfound, err := client.GetArticleById("123")
	if !notfound {
		t.Error("expected not found to be true")
	}

	_, notfound, err = client.GetArticleById("608677")
	if err != nil {
		t.Error("error while fetching by correct id  :", err)
	}
	if notfound {
		t.Error("expected not found to be false")
	}

}
