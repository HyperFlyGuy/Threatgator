package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/hyperflyguy/ThreatGator/internal/database"
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

func handlerListFeeds(s *state, cmd command) error {
	if len(cmd.args) > 0 {
		fmt.Println("Too many arguements (rss_handler.go)")
		os.Exit(1)
		return nil
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		fmt.Println("Unable to retrieve feeds from Database (rss_handler.go): ", err)
		os.Exit(1)
	}
	for _, feed := range feeds {
		fmt.Println(feed)
	}
	os.Exit(0)
	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		fmt.Println("Not Enough Arguments (rss_handler.go)")
		os.Exit(1)
		return nil
	}
	user, err := s.db.GetUser(context.Background(), s.c.CurrentUsername)
	if err != nil {
		fmt.Println("Failed to fetch user UUID from Database (rss_handler.go):", err)
		os.Exit(1)
		return err
	}

	feed := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
		Url:       cmd.args[1],
		UserID:    user.ID,
	}
	s.db.CreateFeed(context.Background(), feed)
	fmt.Println(feed)
	os.Exit(0)
	return nil
}

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml" //cmd.args[0]
	res, err := fetchFeed(context.Background(), url)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// Format the request
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		fmt.Println("Failed to create the request (rss_handler.go):", err)
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")
	// Make the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to make the request (rss_handler.go):", err)
		return nil, err
	}
	defer resp.Body.Close()
	// Handle the response
	var feed RSSFeed
	body, err := io.ReadAll(resp.Body)
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		fmt.Println("Failed to unmarshal data (rss_handler.go):", err)
		return nil, err
	}
	feed.cleanFeed()
	return &feed, nil
}

func (rf *RSSFeed) cleanFeed() error {
	rf.Channel.Title = html.UnescapeString(rf.Channel.Title)
	rf.Channel.Description = html.UnescapeString(rf.Channel.Description)
	if len(rf.Channel.Item) == 0 {
		return fmt.Errorf("No items found at specified URL")
	}
	for i := range rf.Channel.Item {
		rf.Channel.Item[i].Title = html.UnescapeString(rf.Channel.Item[i].Title)
		rf.Channel.Item[i].Description = html.UnescapeString(rf.Channel.Item[i].Description)
	}
	return nil
}
