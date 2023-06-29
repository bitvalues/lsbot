package auctions

import (
	"github.com/bitvalues/lsbot/pkg/bot"
	"github.com/sirupsen/logrus"
)

type ExamplePlugin struct {
	bot *bot.Bot
	log *logrus.Entry
}

func NewExamplePlugin(bot *bot.Bot) *ExamplePlugin {
	return &ExamplePlugin{
		bot: bot,
	}
}

func (p ExamplePlugin) GetName() string {
	return "example"
}

func (p ExamplePlugin) GetPrimaryCommand() string {
	return "example"
}

func (p *ExamplePlugin) Startup(log *logrus.Entry) {
	p.log = log

	p.log.Debug("Starting up...")
}

func (p *ExamplePlugin) Shutdown() {
	p.log.Debug("Shutting down...")
}

func (p *ExamplePlugin) HandlePrimaryCommand(args []string) {
	p.log.Debug("Handling primary command...")
}
