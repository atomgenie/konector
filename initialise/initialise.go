package initialise

import (
	"flag"
	"fmt"

	"github.com/atomgenie/konector/config"
	"github.com/atomgenie/konector/github"
	"github.com/atomgenie/konector/ssh"
	"github.com/atomgenie/konector/systemctl"
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

	if err := config.Save(); err != nil {
		return err
	}

	if err := ssh.SaveKeys(data); err != nil {
		return err
	}

	if err := systemctl.Init(); err != nil {
		return err
	}

	return nil
}
