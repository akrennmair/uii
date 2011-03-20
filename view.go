package main

import (
	"stfl"
	"strings"
	"fmt"
	"strconv"
)

type View struct {
	f *stfl.Form
	quit bool
	lines uint
	linequeue chan LineMsg
	info ServerInfo
}

func CreateView(linequeue chan LineMsg, info ServerInfo) (v *View) {
	v = new(View)
	v.f = stfl.Create("<screen.stfl>")
	v.quit = false
	v.lines = 0
	v.linequeue = linequeue
	v.info = info
	return
}

func(v *View) Run() {

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
	WriteLine("ii-data/" + v.info.Server + "/in", line)
	if line[0] == '/' {
		elements := strings.Split(line[1:], " ", -1)
		v.execCmd(elements[0], elements[1:])
	}
}


func(v *View) execCmd(cmd string, args []string) {
	switch cmd {
	case "quit":
		v.quit = true
	}
}

func(v *View) AddLine(line string) {
	v.f.Modify("mainlist", "append_inner", fmt.Sprintf("{list {listitem text:%s}}", stfl.Quote(line)))
	v.f.Set("mainlistpos", strconv.Uitoa(v.lines))
	v.lines++
	// TODO: implement properly
}

func(v *View) showError(errmsg string) {
	v.AddLine("Error: " + errmsg)
}

func(v *View) UpdateScreen() {
	v.f.Run(-1)
}
