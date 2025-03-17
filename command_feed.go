package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hulkbusterks/BlogAggregator/internal/database"
)

func handlerAddFeed(s *state, cmd Command) error {
	if len(cmd.args) != 2 {
		return errors.New("too much parameters")
	}

	name := cmd.args[0]
	url := cmd.args[1]
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	rssfeed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      name,
		Url:       url,
	})
	if err != nil {
		return err
	}

	fmt.Println("Feed Created Successfully")
	fmt.Println(rssfeed)
	return nil
}

func handlerListFeeds(s *state, cmd Command) error {
	if len(cmd.args) != 0 {
		return errors.New("too many parameters")
	}
	feeds, err := s.db.ListFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Println("no feeds registered")
	}

	var name string
	var url string
	var user database.User
	for _, feed := range feeds {
		name = feed.Name
		url = feed.Url
		user, err = s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		fmt.Printf("name: %s, url: %s, user: %s \n", name, url, user.Name)
	}
	return nil
}
