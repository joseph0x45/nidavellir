package main

var serviceFile = `
[Unit]
Description=Nidavellir Package Repository
After=network.target
Wants=network.target

[Service]
Type=simple
EnvironmentFile=/etc/nidavellir/nidavellir.env
ExecStart=/usr/local/bin/nidavellir
Restart=on-failure
RestartSec=2

NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=full
ProtectHome=true

StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
`
