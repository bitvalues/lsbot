package config

import "os"

type Config struct {
	DiscordAuthToken   string
	DiscordOfficerRole string
	DiscordMemberRole  string
}

func GetConfig() Config {
	config := Config{
		DiscordAuthToken:   os.Getenv("DISCORD_BOT_TOKEN"),
		DiscordOfficerRole: os.Getenv("DISCORD_OFFICER_ROLE"),
		DiscordMemberRole:  os.Getenv("DISCORD_MEMBER_ROLE"),
	}

	return config
}
