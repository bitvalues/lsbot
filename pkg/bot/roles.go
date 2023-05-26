package bot

import "github.com/bwmarrin/discordgo"

func (b *Bot) MemberHasRole(member *discordgo.Member, role string) bool {
	for _, memberRole := range member.Roles {
		if role == memberRole {
			return true
		}
	}

	return false
}
