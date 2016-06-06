package config

type config struct {
	ApiKeys struct {
		Slack    string `yaml:"slack"`
		Opsgenie string `yaml:"opsgenie"`
	} `yaml:"api_keys"`
	Schedules []Schedule
}

type Schedule struct {
	Name           string `yaml:"name"`
	SlackUserGroup string `yaml:"slackUserGroup"`
	UpdateMessage  struct {
		Message  string   `yaml:"message"`
		Channels []string `yaml:"channels"`
	} `yaml:"updateMessage"`
}
