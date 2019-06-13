package cmd

import (
	"github.com/arnaucube/gogame/config"
	"github.com/arnaucube/gogame/database"
	"github.com/arnaucube/gogame/endpoint"
	"github.com/arnaucube/gogame/services/usersrv"
	"github.com/urfave/cli"
)

var ServerCommands = []cli.Command{
	{
		Name:    "start",
		Aliases: []string{},
		Usage:   "start the server",
		Action:  start,
	},
}

func start(c *cli.Context) error {
	if err := config.MustRead(c); err != nil {
		return err
	}

	db, err := database.New(config.C.Mongodb.Url, config.C.Mongodb.Database)
	if err != nil {
		return err
	}

	// services
	userservice := usersrv.New(db)
	if err != nil {
		return err
	}

	apiService := endpoint.Serve(config.C, db, userservice)
	apiService.Run(config.C.Server.ServiceApi)

	return nil
}
