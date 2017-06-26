package services

// JenkinsSettings holds the information about one or more Jenkins servers that
// the bot should send or retrieve information from.
type JenkinsSettings struct {
	JenkinsURL   string `json:"jenkinsURL"`
	JenkinsUser  string `json:"jenkinsUser"`
	JenkinsToken string `json:"jenkinsToken"`
}
