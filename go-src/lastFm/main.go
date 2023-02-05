package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	err, track := getMostRecentTrack()
	if err != nil {
		return &events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: http.StatusInternalServerError,
		}, nil
	}
	resp := formatResults(track)
	json, _ := json.Marshal(resp)
	return &events.APIGatewayProxyResponse{
		Body:       string(json),
		StatusCode: http.StatusOK,
	}, nil

}

type Size string

const url string = "http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=holopollock&api_key=%s&format=json"

const (
	Small  Size = "small"
	Medium      = "medium"
	Large       = "large"
	XLarge      = "extralarge"
)

type RecentTracksReturn struct {
	Results RecentTrackResults `json:"recenttracks"`
}

type RecentTrackResults struct {
	RecentTracks []Track `json:"track"`
	Attributes   struct {
		user       string
		totalPages int
		page       int
		perPage    int
		total      int
	} `json:"@attr"`
}

type Track struct {
	Artist struct {
		Mbid string `json:"mbid"`
		Name string `json:"#text"`
	} `json:"artist"`
	IsStreamable int           `json:"streamable"`
	Image        []lastFMImage `json:"image"`
	Mbid         string        `json:"mbid"`
	Album        struct {
		Mbid string `json:"mbid"`
		Name string `json:"#text"`
	} `json:"album"`
	Name       string `json:"name"`
	Attributes struct {
		IsPlaying bool `json:"nowplaying"`
	} `json:"@attr"`
	Url string `json:"url"`
}

type lastFMImage struct {
	Size Size   `json:"size"`
	Link string `json:"#text"`
}

type RecentTrackResponse struct {
	SongName string `json:"name"`
	Artist   string `json:"artist"`
	Album    string `json:"album"`
	Url      string `json:"url"`
}

func formatResults(track Track) RecentTrackResponse {
	return RecentTrackResponse{
		SongName: track.Name,
		Artist:   track.Artist.Name,
		Album:    track.Album.Name,
		Url:      track.Url,
	}
}

func getMostRecentTrack() (error, Track) {
	myClient := &http.Client{Timeout: 30 * time.Second}
	results := &RecentTracksReturn{}
	apiKey := os.Getenv("LAST_FM_KEY")
	resp, err := myClient.Get(fmt.Sprintf(url, apiKey))
	if err != nil {
		return err, Track{}
	}
	defer resp.Body.Close()
	json.NewDecoder(resp.Body).Decode(results)
	return nil, results.Results.RecentTracks[0]
}

func main() {
	lambda.Start(Handler)
}
