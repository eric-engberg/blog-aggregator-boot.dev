package main

import (
	"context"
	"fmt"
	"html"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <time_duration>", cmd.Name)
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("unable to parse duration: %w", err)
	}

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		scrapeFeeds(s)
	}
}

func scrapeFeeds(s *state) error {
	next_feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("unable to get next feed: %w", err)
	}

	feed, err := fetchFeed(context.Background(), next_feed.Url)
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), next_feed.ID)
	if err != nil {
		return fmt.Errorf("unable to mark feed as fetched: %w", err)
	}

	for i := range feed.Channel.Item {
		fmt.Print(html.UnescapeString(feed.Channel.Item[i].Title) + "\n")
	}

	return nil
}
