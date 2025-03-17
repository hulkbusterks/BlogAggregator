package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hulkbusterks/BlogAggregator/internal/database"
)

func handlerLogin(s *state, cmd Command) error {
	if len(cmd.args) != 1 {
		return errors.New(fmt.Sprintf("Need only one argument for %s: cli <command> [args...]", cmd.name))
	}
	name := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return errors.New("User does not exists")
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return errors.New("Couldn't set current user")
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd Command) error {
	if len(cmd.args) != 1 {
		return errors.New(fmt.Sprintf("Need only one argument for %s: cli <command> [args...]", cmd.name))
	}

	name := cmd.args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	})

	if err != nil {
		return err
	}

	err = s.cfg.SetUser(name)

	if err != nil {
		return err
	}

	fmt.Println(" ID : ", user.ID)
	fmt.Println(" Name: ", user.Name)

	return nil
}

func handlerReset(s *state, cmd Command) error {
	if len(cmd.args) != 0 {
		return errors.New("too much parameters")
	}

	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}
	fmt.Println("database reset successfully")
	return nil
}

func handlerUsers(s *state, cmd Command) error {
	if len(cmd.args) != 0 {
		return errors.New("too much parameters")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Println("* ", user.Name, "(current)")
		} else {
			fmt.Println("* ", user.Name)
		}
	}
	return nil
}
