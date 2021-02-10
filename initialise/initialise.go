package initialise

import (
	"flag"
	"fmt"

	"github.com/atomgenie/konector/config"
	"github.com/atomgenie/konector/github"
	"github.com/atomgenie/konector/ssh"
)

// Init Init konector
func Init(argv []string) error {
	subCommand := flag.NewFlagSet("init", flag.ExitOnError)

	username := subCommand.String("username", "", "Github username")
	interval := subCommand.Int("interval", 1, "Check SSH keys every x minutes")

	subCommand.Parse(argv[1:])

	if *username == "" {
		subCommand.Usage()
		return fmt.Errorf("Invalid username")
	}

	data, err := github.GetSSHKeys(*username)

	if err != nil {
		return err
	}

	config := config.UserConfig{
		Username: *username,
		Interval: *interval,
	}

	err = config.Save()

	if err != nil {
		return err
	}

	ssh.SaveKeys(data)

	return nil
}
