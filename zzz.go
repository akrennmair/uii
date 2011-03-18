package main

import (
	"flag"
	"os"
	"fmt"
)

type ServerInfo struct {
	Server string
	Port int
	Nick string
}

func main() {
	info := ServerInfo{ "", 6767, "" }

	flag.IntVar(&info.Port, "port", 6767, "IRC server port")
	flag.StringVar(&info.Server, "server", "", "IRC server hostname")
	flag.StringVar(&info.Nick, "nick", os.Getenv("USER"), "Your nick")

	flag.Parse()

	if info.Server == "" {
		usage()
	}

	v := CreateView()

	v.Run()

}

func usage() {
	fmt.Println("usage: zzz -server <server> [-port <port>] [-nick <nick>]")
	os.Exit(1)
}
