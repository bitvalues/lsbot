package bot

import "github.com/bwmarrin/discordgo"

func (b *Bot) doHelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	b.session.ChannelMessageSend(m.ChannelID, "Valid commands are: `auction`, `tod` & `help`. ")
}
