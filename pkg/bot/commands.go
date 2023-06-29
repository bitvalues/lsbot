package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) onMessageCreated(s *discordgo.Session, m *discordgo.MessageCreate) {
	// cache the session, just in case
	b.UpdateSession(s)

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

	// now, iterate through the plugins and act accordingly
	for _, plugin := range b.plugins {
		if strings.ToLower(parts[0]) == strings.ToLower(plugin.GetPrimaryCommand()) {
			plugin.HandlePrimaryCommand(parts[1:])
		}
	}
}

func (b *Bot) doHelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	validCommands := []string{}

	for _, plugin := range b.plugins {
		validCommands = append(validCommands, "`"+plugin.GetPrimaryCommand()+"`")
	}

	b.GetSession().ChannelMessageSend(m.ChannelID, "Valid commands are: "+strings.Join(validCommands, ", "))
}
