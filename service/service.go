package service

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/atomgenie/konector/config"
	"github.com/atomgenie/konector/github"
	"github.com/atomgenie/konector/ssh"
)

//StartService Start Konector service
func StartService() error {
	sigs := make(chan os.Signal, 1)

	// catch all signals since not explicitly listing
	signal.Notify(sigs, syscall.SIGQUIT, syscall.SIGTERM)

	for {

		config, err := config.Load()

		if err != nil {
			return err
		}

		var data github.APISSHKeys = []struct {
			ID  int    "json:\"id\""
			Key string "json:\"key\""
		}{}

		for _, username := range config.Usernames {

			_data, err := github.GetSSHKeys(username)

			if err != nil {
				return err
			}

			data = append(data, _data...)
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
