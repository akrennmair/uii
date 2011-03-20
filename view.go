package main

import (
	"stfl"
	"strings"
	"fmt"
	"strconv"
	"time"
)

type View struct {
	f *stfl.Form
	quit bool
	lines uint
}

func CreateView() (v *View) {
	v = new(View)
	v.f = stfl.Create("<screen.stfl>")
	v.quit = false
	v.lines = 0
	return
}

func(v *View) Run() {

	go func() {
		for {
			time.Sleep(1000000000)
			v.sendLine(time.LocalTime().String())
			v.UpdateScreen()
		}
	}()

	for !v.quit {
		key := v.f.Run(0)
		switch key {
		case "ENTER":
			focus := v.f.GetFocus()
			if focus == "input" {
				text := v.f.Get("inputtext")
				if len(text) > 0 {
					v.handleInput(text)
					v.f.Set("inputtext", "")
				}
			}
		// TODO: more keys
		}
	}

	v.f.Free()
	stfl.Reset()
}

func(v *View) handleInput(line string) {
	if line[0] == '/' {
		elements := strings.Split(line[1:], " ", -1)
		v.execCmd(elements[0], elements[1:])
	} else {
		v.sendLine(line)
	}
}

func(v *View) execCmd(cmd string, args []string) {
	switch cmd {
	case "quit":
		v.quit = true
	default:
		v.showError(fmt.Sprintf("unknown command '%s'", cmd))
	}
}

func(v *View) sendLine(line string) {
	v.f.Modify("mainlist", "append_inner", fmt.Sprintf("{list {listitem text:%s}}", stfl.Quote(line)))
	v.f.Set("mainlistpos", strconv.Uitoa(v.lines))
	v.lines++
	// TODO: implement properly
}

func(v *View) showError(errmsg string) {
	v.sendLine("Error: " + errmsg)
}

func(v *View) UpdateScreen() {
	v.f.Run(-1)
}
