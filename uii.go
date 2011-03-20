package main

import (
	"flag"
	"os"
	"fmt"
	"stfl"
	"bufio"
	"time"
	"exec"
	"strconv"
)

type ServerInfo struct {
	Server string
	Port int
	Nick string
}

func init() {
	stfl.Init()
}

func main() {
	info := ServerInfo{ "", 6666, "" }

	flag.IntVar(&info.Port, "port", 6666, "IRC server port")
	flag.StringVar(&info.Server, "server", "", "IRC server hostname")
	flag.StringVar(&info.Nick, "nick", os.Getenv("USER"), "Your nick")

	flag.Parse()

	if info.Server == "" {
		usage()
	}

	wmanchan := make(chan LineMsg, 16)
	v := CreateView(wmanchan, info)
	wman := CreateWindowManager(v, wmanchan)

	go wman.Run()

	startii(info, wman)

	v.Run()
}

func usage() {
	fmt.Println("usage: uii -server <server> [-port <port>] [-nick <nick>]")
	os.Exit(1)
}

func startii(info ServerInfo, wman *WindowManager) {
	ii_path, err := exec.LookPath("ii")
	if err != nil {
		fmt.Println("Error: couldn't find ii.")
		os.Exit(1)
	}

	_, err = exec.Run(ii_path, []string{ "ii", "-i", "ii-data", "-s", info.Server, "-p", strconv.Itoa(info.Port), "-n", info.Nick }, []string{ }, ".", exec.DevNull, exec.DevNull, exec.DevNull)
	if err != nil {
		fmt.Printf("Running ii failed: %s\n", err)
		os.Exit(1)
	}

	go monitorFile("ii-data/" + info.Server + "/out", info.Server, wman)
}

func monitorFile(filename string, ircchan string, wman *WindowManager) {
	var file *os.File

	for {
		f, err := os.Open(filename, os.O_RDONLY, 0)
		if err != nil {
			time.Sleep(1000000000)
		} else {
			file = f
			break
		}
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil && len(line) == 0 {
			time.Sleep(1000000000)
			if err == os.EOF {
				// XXX hack to reset EOF
				reader = bufio.NewReader(file)
			}
			continue
		}
		msg := LineMsg{ string(line[:len(line)-1]), ircchan }
		wman.LineQueue <- msg
	}
}

func WriteLine(filename string, line string) {
	f, err := os.Open(filename, os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	f.Write([]byte(line + "\n"))
	f.Close()
}
