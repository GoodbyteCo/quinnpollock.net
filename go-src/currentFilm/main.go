package main

import (
	"encoding/json"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mmcdole/gofeed"
)

func Handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	latestMovie := getLatestMovie(getRssFeed())
	e, _ := json.Marshal(latestMovie)
	return &events.APIGatewayProxyResponse{
		Body:       string(e),
		StatusCode: 200,
	}, nil

}

const feed string = "https://letterboxd.com/holopollock/rss/"

type film struct {
	Title string `json:"title"`
	Link string `json:"link"`
	Image string `json:"image"`
}

func getRssFeed() *gofeed.Feed {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feed)
	if err != nil {
		panic(err)
	}
	return feed
}

func getLatestMovie(feed *gofeed.Feed) film {
	return film{
		Title: feed.Items[0].Title,
		Link: linkTransformer(feed.Items[0].Link),
		Image: getImageLink(feed.Items[0].Description),
	}
}

func linkTransformer(link string) string {
	return strings.Replace(link, "holopollock/", "", 1)
}

func getImageLink(link string) string {
	i := strings.Index(link, "src")
	afterSrc := link[i+5:]
	end := strings.Index(afterSrc, "/>")
	fullLink := afterSrc[:end-1]
	return strings.Replace(fullLink, "0-600-0-900", "0-300-0-450", 1)
} 



func main() {
	lambda.Start(Handler)
}