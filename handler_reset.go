package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	err := s.db.DeleteUsers(ctx)
	if err != nil {
		return fmt.Errorf("failed to delate all users: %v", err)
	}
	fmt.Println("successful delete all users")
	return nil
}
