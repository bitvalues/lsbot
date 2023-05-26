package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) handleAuctionCommand(s *discordgo.Session, m *discordgo.MessageCreate, commands []string) {
	if len(commands) < 1 {
		b.session.ChannelMessageSend(m.ChannelID, "The `auction` command must contain a valid subcommand (`create`, `cancel`, `bid`).")
		return
	}

	switch strings.ToLower(commands[0]) {
	case "create":
		b.doAuctionCreate(s, m, commands[1:])
		break
	case "cancel":
		b.doAuctionCancel(s, m, commands[1:])
		break
	case "bid":
		b.doAuctionBid(s, m, commands[1:])
	}
}
