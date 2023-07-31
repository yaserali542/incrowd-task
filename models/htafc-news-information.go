package models

import "encoding/xml"

type HtafcNewsArticleInformation struct {
	XMLName        xml.Name         `xml:"NewsArticleInformation"`
	Text           string           `xml:",chardata"`
	ClubName       string           `xml:"ClubName"`
	ClubWebsiteURL string           `xml:"ClubWebsiteURL"`
	NewsArticle    HtafcNewsArticle `xml:"NewsArticle"`
}

type HtafcNewsArticle struct {
	Text              string `xml:",chardata"`
	ArticleURL        string `xml:"ArticleURL"`
	NewsArticleID     string `xml:"NewsArticleID"`
	PublishDate       string `xml:"PublishDate"`
	Taxonomies        string `xml:"Taxonomies"`
	TeaserText        string `xml:"TeaserText"`
	Subtitle          string `xml:"Subtitle"`
	ThumbnailImageURL string `xml:"ThumbnailImageURL"`
	Title             string `xml:"Title"`
	BodyText          string `xml:"BodyText"`
	GalleryImageURLs  string `xml:"GalleryImageURLs"`
	VideoURL          string `xml:"VideoURL"`
	OptaMatchId       string `xml:"OptaMatchId"`
	LastUpdateDate    string `xml:"LastUpdateDate"`
	IsPublished       bool   `xml:"IsPublished"`
}
