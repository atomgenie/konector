package service

import (
	"os"
	"os/signal"
	"time"

	"github.com/atomgenie/konector/config"
	"github.com/atomgenie/konector/github"
	"github.com/atomgenie/konector/ssh"
)

//StartService Start Konector service
func StartService() error {
	sigs := make(chan os.Signal, 1)

	// catch all signals since not explicitly listing
	signal.Notify(sigs)

	config, err := config.Load()

	if err != nil {
		return err
	}

	for {
		data, err := github.GetSSHKeys(config.Username)

		if err != nil {
			return err
		}

		err = ssh.SaveKeys(data)

		if err != nil {
			return err
		}

		exitLoop := false

		select {
		case <-sigs:
			exitLoop = true
		case <-time.After(time.Duration(config.Interval) * time.Minute):
			break
		}

		if exitLoop {
			break
		}
	}

	return nil
}
