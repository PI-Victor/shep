package services

// IRCSettings holds information about one or more IRC servers that the bot
// could join.
type IRCSettings struct {
	IRCServers []IRCServer `json:"ircServers"`
}

// IRCServer holds information regarding one or more IRC servers that the bot
// should connect to.
type IRCServer struct {
	ServerName string   `json:"serverName"`
	Nick       string   `json:"nick"`
	Channels   []string `json:"channels"`
}
