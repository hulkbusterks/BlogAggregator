package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hulkbusterks/BlogAggregator/internal/database"
)

func handlerFollow(s *state, cmd Command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New("need only 1 argument")
	}

	url := cmd.args[0]

	feed_ID, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	follow, err := s.db.CreateFeedFollow(context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    feed_ID.ID,
		})
	if err != nil {
		return err
	}

	fmt.Println(follow.UserName, follow.FeedName)

	fmt.Println("Followed success")

	return nil
}

func handlerFollowing(s *state, cmd Command, user database.User) error {
	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't get feed follows: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}

	fmt.Printf("Feed follows for this user %s:\n", user.Name)
	for _, ff := range feedFollows {
		fmt.Printf("* %s\n", ff.FeedName)
	}

	return nil
}

func handlerUnfollow(s *state, cmd Command, user database.User) error {
	if len(cmd.args) != 1 {
		return errors.New(fmt.Sprintf("Need only one argument for %s: cli <command> [args...]", cmd.name))
	}
	url := cmd.args[0]

	feed_ID, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}
	err = s.db.DeleteFeedFollow(context.Background(),
		database.DeleteFeedFollowParams{
			UserID: user.ID,
			FeedID: feed_ID.ID,
		})
	if err != nil {
		return err
	}

	fmt.Println("Unfollowed successfully!")
	return nil
}
