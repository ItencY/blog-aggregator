package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	ctx := context.Background()
	feeds, err := s.db.GetFeeds(ctx)
	if err != nil {
		return fmt.Errorf("failed get feeds: %v", err)
	}
	for _, feed := range feeds {
		fmt.Println(feed.FeedName)
		fmt.Println(feed.Url)
		fmt.Println(feed.UserName)
	}
	return nil
}
