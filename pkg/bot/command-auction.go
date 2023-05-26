package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) handleAuctionCommand(s *discordgo.Session, m *discordgo.MessageCreate, commands []string) {
	if len(commands) < 1 {
		b.doAuctionCommand(s, m)
		return
	}

	switch strings.ToLower(commands[0]) {
	case "create":
		b.doAuctionCreateCommand(s, m, commands[1:])
		break
	case "cancel":
		b.doAuctionCancelCommand(s, m, commands[1:])
		break
	case "bid":
		b.doAuctionBidCommand(s, m, commands[1:])
		break
	default:
		b.doAuctionCommand(s, m)
		break
	}
}

func (b *Bot) doAuctionCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	b.session.ChannelMessageSend(m.ChannelID, "The `auction` command must contain a valid subcommand (`create`, `cancel`, `bid`).")
}
