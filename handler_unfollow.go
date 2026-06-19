package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/itency/blog_aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("the unfollow handler expects a single argument, the command name")
	}

	ctx := context.Background()

	feed, err := s.db.GetFeedByURL(ctx, cmd.args[0])
	if err != nil {
		return fmt.Errorf("couldn't get feed: %v", err)
	}

	err = s.db.DeleteFeedFollow(ctx, database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't delete feed follow: %v", err)
	}
	fmt.Printf("%s unfollowed successfully!\n", feed.Name)
	return nil
}
