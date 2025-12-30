package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"text/template"
)

var serviceFile = `
[Unit]
Description=Nidavellir Package Repository
After=network.target
Wants=network.target

[Service]
Type=simple

User={{.User}}

Environment=XDG_CONFIG_HOME=%h/.config
Environment=XDG_DATA_HOME=%h/.local/share

EnvironmentFile=-%h/.config/nidavellir/conf

ExecStart=/usr/local/bin/nidavellir
Restart=on-failure
RestartSec=5

NoNewPrivileges=true

StandardOutput=journal
StandardError=journal

[Install]
WantedBy=multi-user.target
`

func installService() {
	if os.Getuid() == 0 {
		fmt.Println("Do not run this command with sudo")
		return
	}
	t, err := template.New("service").Parse(serviceFile)
	if err != nil {
		log.Fatalln(err)
	}
	f, err := os.Create("nidavellir.service")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	user, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	if err := t.Execute(f, map[string]string{
		"User": user.Username,
	}); err != nil {
		log.Fatalln(err)
	}
	if err := f.Sync(); err != nil {
		log.Fatalln(err)
	}
	cmd := exec.Command("sudo", "cp", "nidavellir.service", "/etc/systemd/system/")
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	log.Println("Service file created at /etc/systemd/system/nidavellir.service")
}
