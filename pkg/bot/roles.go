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

func (b *Bot) IsValidOfficer(member *discordgo.Member) bool {
	return b.MemberHasRole(member, b.cfg.DiscordOfficerRole)
}

func (b *Bot) IsValidMember(member *discordgo.Member) bool {
	return b.MemberHasRole(member, b.cfg.DiscordMemberRole)
}
