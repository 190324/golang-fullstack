package main

import (
	"github.com/robfig/cron/v3"
	commands "XinAPI/app/console/commands"
)

func main() {
	c := cron.New(cron.WithSeconds())

    c.AddFunc("*/5 * * * * *", commands.Test)

    c.Start()
    select {}
}
