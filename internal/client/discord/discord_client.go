package discord

import (
	"github.com/bwmarrin/discordgo"
	"haracho-go/internal/client"
	"haracho-go/internal/logger"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Client struct {
	Token      string
	session    *discordgo.Session
	Collection *client.CommandCollection
}

var l = logger.BasicLogger{Logger: log.New(os.Stdout, "DiscordClient: ", log.LstdFlags)}

func (c *Client) Start() {
	dg, err := discordgo.New("Bot " + c.Token)
	if err != nil {
		l.Errorfln("error creating discord session. %v", err)
		return
	}

	c.session = dg

	dg.AddHandler(func(session *discordgo.Session, create *discordgo.MessageCreate) {
		parser := Parser{message: create.Content, ctx: &CommandContext{client: c, channel: create.ChannelID, log: &l}}
		client.ExecuteCommand(parser, c.Collection)
	})

	err = dg.Open()
	if err != nil {
		l.Errorfln("error opening connection. %v", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	_ = dg.Close()
}
