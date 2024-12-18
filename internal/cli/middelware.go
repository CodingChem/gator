package cli

import (
	"context"

	"github.com/codingchem/gator/internal/database"
)

func middlewareLoggedIn(handler func(s * state, cmd command, user database.User) error) (func(s *state, cmd command) error) {
	return func (s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.config.CurrentUser)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}
