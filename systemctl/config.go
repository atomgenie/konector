package systemctl

import "io/ioutil"

const systemctlConfig = `
[Unit]
Description=Konector daemon
After=network-online.target
StartLimitIntervalSec=0

[Service]
Type=simple
User=%u
ExecStart=/usr/bin/env konector service
Restart=on-failure
RestartSec=30
# Configures the time to wait before service is stopped forcefully.

[Install]
WantedBy=multi-user.target
`

// Init systemctl
func Init() error {
	err := ioutil.WriteFile("/etc/systemd/system/konector.service", []byte(systemctlConfig), 0777)
	return err
}
