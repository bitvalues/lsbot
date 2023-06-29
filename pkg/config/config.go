package config

import "os"

type Config struct {
	DiscordAuthToken           string
	DiscordOfficerRole         string
	DiscordMemberRole          string
	DiscordDKPAuctionChannelID string
}

func GetConfig() Config {
	config := Config{
		DiscordAuthToken:           os.Getenv("DISCORD_BOT_TOKEN"),
		DiscordOfficerRole:         os.Getenv("DISCORD_OFFICER_ROLE"),
		DiscordMemberRole:          os.Getenv("DISCORD_MEMBER_ROLE"),
		DiscordDKPAuctionChannelID: os.Getenv("DISCORD_DKP_AUCTION_CHANNEL_ID"),
	}

	return config
}
