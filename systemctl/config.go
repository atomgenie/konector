package systemctl

const systemctlConfig = `
[Unit]
Description=Konector daemon
After=network-online.target
StartLimitIntervalSec=0

[Service]
Type=simple
User=%i
ExecStart=
Restart=on-failure
RestartSec=30
# Configures the time to wait before service is stopped forcefully.

[Install]
WantedBy=multi-user.target
`
