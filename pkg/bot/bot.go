package bot

import (
	"strings"

	"github.com/bitvalues/lsbot/pkg/config"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	session *discordgo.Session
	log     *logrus.Logger
	cfg     config.Config
}

func NewBot(cfg config.Config, logger *logrus.Logger) (*Bot, error) {
	session, err := discordgo.New("Bot " + cfg.DiscordAuthToken)
	if err != nil {
		return nil, err
	}

	bot := Bot{
		session: session,
		log:     logger,
		cfg:     cfg,
	}

	return &bot, nil
}

func (b *Bot) Startup() {
	b.session.AddHandler(b.onMessageCreate)
	b.session.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	if err := b.session.Open(); err != nil {
		b.log.WithError(err).Error("error opening connection")
		return
	}
}

func (b *Bot) onMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// cache the new session pointer (just in case)
	b.session = s

	// make sure we weed-out non-commands
	if !strings.HasPrefix(strings.ToLower(m.Content), "!lsbot") {
		return
	}

	// make sure there are at least 7 characters before we continue
	if len(m.Content) <= 7 {
		b.doHelpCommand(s, m)
		return
	}

	// parse the message command
	parts := strings.Split(m.Content[7:], " ")

	// now, take an action accordingly
	switch strings.ToLower(parts[0]) {
	case "auction":
		b.handleAuctionCommand(s, m, parts[1:])
		break
	case "tod":
		b.doTODCommand(s, m)
		break
	default:
		b.doHelpCommand(s, m)
		break
	}
}

func (b *Bot) Shutdown() {
	b.session.Close()
}
