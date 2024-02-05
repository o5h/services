package main

import (
	"fmt"
	"time"

	"github.com/o5h/services/config"
	"github.com/o5h/services/context"
	"github.com/o5h/services/db"
	"github.com/o5h/services/services"
)

var (
	version string = "dev"
	date    string = time.Now().Format("2006-01-02 15:04:05")
)

func main() {

	fmt.Printf("version=%s, date=%s\n", version, date)

	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		return
	}
	conf.BuildVersion = version
	conf.BuildDate = date

	globalContext, cancel := context.Init(conf)
	db.Init(globalContext)
	services.Start(globalContext, cancel)

}
