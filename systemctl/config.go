package systemctl

import (
	"flag"
	"io/ioutil"
	"strings"
)

const systemctlConfig = `
[Unit]
Description=Konector daemon
After=network-online.target
StartLimitIntervalSec=0

[Service]
Type=simple
User=__user_name__
ExecStart=/usr/bin/env konector service
Restart=on-failure
RestartSec=30
# Configures the time to wait before service is stopped forcefully.

[Install]
WantedBy=multi-user.target
`

// Init systemctl
func Init(argv []string) error {
	subCmd := flag.NewFlagSet("init-systemctl", flag.ExitOnError)
	user := subCmd.String("user", "root", "The os user")

	subCmd.Parse(argv)

	withUser := strings.ReplaceAll(systemctlConfig, "__user_name__", *user)

	err := ioutil.WriteFile("/etc/systemd/system/konector.service", []byte(withUser), 0777)
	return err
}
