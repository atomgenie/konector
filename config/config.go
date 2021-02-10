package config

// UserConfig Config of konector
type UserConfig struct {
	Usernames []string `json:"usernames"`
	Interval  int      `json:"interval"`
}
