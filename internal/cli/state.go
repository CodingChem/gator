package cli

import (
	"database/sql"
	"github.com/codingchem/gator/internal/config"
	"github.com/codingchem/gator/internal/database"
)

type state struct {
	config *config.Config
	cmds   *commands
	db     *database.Queries
}

func NewState() (state, error) {
	globalConfig, err := config.Read()
	if err != nil {
		return state{}, err
	}
	cmds, err := NewCommands()
	if err != nil {
		return state{}, err
	}
	db, err := sql.Open("postgres", globalConfig.DB_CON_STRING)
	if err != nil {
		return state{}, err
	}
	dbQueries := database.New(db)
	return state{
		config: &globalConfig,
		cmds:   &cmds,
		db:     dbQueries,
	}, nil
}

func (s *state) Run(name string, args []string) error {
	return s.cmds.run(s, command{name: name, args: args})
}
