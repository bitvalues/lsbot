package auctions

import (
	"github.com/bitvalues/lsbot/pkg/bot"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
)

type AuctionsPlugin struct {
	bot *bot.Bot
	log *logrus.Entry
}

func NewAuctionsPlugin(bot *bot.Bot) *AuctionsPlugin {
	return &AuctionsPlugin{
		bot: bot,
	}
}

func (p AuctionsPlugin) GetName() string {
	return "auction"
}

func (p AuctionsPlugin) GetPrimaryCommand() string {
	return "auction"
}

func (p *AuctionsPlugin) Startup(log *logrus.Entry) {
	p.log = log

	p.log.Debug("Starting up...")
}

func (p *AuctionsPlugin) Shutdown() {
	p.log.Debug("Shutting down...")
}

func (p *AuctionsPlugin) HandlePrimaryCommand(args []string) {
	p.log.Debug("Handling primary command...")
	spew.Dump(args)
}
