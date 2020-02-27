package main

import (
	commands "XinAPI/app/console/commands"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())

	c.AddFunc("*/5 * * * * *", commands.Test)

	c.Start()
	select {}
}
