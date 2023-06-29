package bot

import (
	"github.com/bitvalues/lsbot/pkg/config"
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

type Bot struct {
	session *discordgo.Session
	log     *logrus.Logger
	cfg     config.Config
	plugins []BotPlugin
}

type BotPlugin interface {
	GetName() string
	GetPrimaryCommand() string
	Startup(log *logrus.Entry)
	Shutdown()
	HandlePrimaryCommand(args []string)
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
		plugins: []BotPlugin{},
	}

	return &bot, nil
}

func (b *Bot) GetSession() *discordgo.Session {
	return b.session
}

func (b *Bot) UpdateSession(session *discordgo.Session) {
	b.session = session
}

func (b *Bot) GetConfig() config.Config {
	return b.cfg
}

func (b *Bot) LoadPlugin(plugin BotPlugin) {
	b.plugins = append(b.plugins, plugin)
}

func (b *Bot) GetLogger() *logrus.Logger {
	return b.log
}

func (b *Bot) Startup() {
	// house-keeping
	b.GetLogger().Debug("Starting up...")

	// Open a websocket connection to Discord and begin listening.
	if err := b.session.Open(); err != nil {
		b.GetLogger().WithError(err).Error("error opening connection")
		return
	}

	// begin loading all of our plugins
	for _, plugin := range b.plugins {
		plugin.Startup(b.GetLogger().WithField("plugin", plugin.GetName()))
	}

	// register a handler for when messages are created
	b.GetSession().AddHandler(b.onMessageCreated)
	b.GetSession().Identify.Intents = discordgo.IntentsGuildMessages
}

func (b *Bot) Shutdown() {
	b.GetLogger().Debug("Shutting down...")

	// make sure the session is closed
	defer b.session.Close()

	// begin shutting down all of our plugins
	for _, plugin := range b.plugins {
		plugin.Shutdown()
	}
}
