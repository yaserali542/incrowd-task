package models

import "time"

type Newsarticles struct {
	ArticleURL        string    `bson:"articleURL,omitEmpty" json:"articleURL,omitEmpty"`
	NewsArticleID     string    `bson:"_id" json:"newsArticleID"`
	PublishDate       time.Time `bson:"publishDate,omitEmpty,$date" json:"publishDate,omitEmpty"`
	Taxonomies        string    `bson:"taxonomies,omitEmpty" json:"taxonomies,omitEmpty"`
	TeaserText        string    `bson:"teaserText,omitEmpty" json:"teaserText,omitEmpty"`
	ThumbnailImageURL string    `bson:"thumbnailImageUrl,omitEmpty" json:"thumbnailImageUrl,omitEmpty"`
	Title             string    `bson:"title,omitEmpty" json:"title,omitEmpty"`
	OptaMatchId       string    `bson:"optaMatchId,omitEmpty" json:"optaMatchId,omitEmpty"`
	LastUpdateDate    time.Time `bson:"lastUpdateDate,omitEmpty,$date" json:"lastUpdateDate,omitEmpty,$date"`
	IsPublished       bool      `bson:"isPublished,omitEmpty" json:"isPublished,omitEmpty"`
	ClubName          string    `bson:"clubName" json:"clubName"`
	ClubWebsiteURL    string    `bson:"clubWebsiteURL" json:"clubWebsiteURL"`
}
