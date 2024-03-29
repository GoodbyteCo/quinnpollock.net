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
	}
}

func linkTransformer(link string) string {
	return strings.Replace(link, "holopollock/", "", 1)
}



func main() {
	lambda.Start(Handler)
}