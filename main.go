package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/agungwicaksana/privy-pretest/cmd"
)

var (
	appName    string = "cakestore"
	appVersion string = "1.0.0"
	err        error
)

func main() {
	if tz := os.Getenv("TZ"); tz != "" {
		time.Local, err = time.LoadLocation(tz)
		if err != nil {
			log.Printf("Error loading location '%s': %v\n", tz, err)
		}
	}

	fmt.Printf(appName, appVersion)
	runtime.GOMAXPROCS(runtime.NumCPU())

	cmd.Run()
}
