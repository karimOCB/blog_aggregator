package main

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	rssfeed := &RSSFeed{}

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return rssfeed, err
	}
	req.Header.Set("User-Agent", "gator")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return rssfeed, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return rssfeed, err
	}

	err = xml.Unmarshal(data, rssfeed)
	if err != nil {
		return rssfeed, err
	}

	return rssfeed, nil
}
