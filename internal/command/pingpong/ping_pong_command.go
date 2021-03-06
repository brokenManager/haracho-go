package pingpong

import (
	"haracho-go/internal/client"
	"haracho-go/internal/client/arg"
)

func Build(collection *client.CommandCollection) {
	help := &client.HelpContext{
		Name:            "!ping",
		Description:     "PingPong",
		ArgsDescription: nil,
	}

	command := client.NewCommand(help, "!ping", func(arg *arg.Arg, ctx client.CommandContext) {
		ctx.SendMessage("pong!")
	})

	collection.AddCommand(command)
}
