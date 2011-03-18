package main

import (
	"stfl"
)

type View struct {
	f *stfl.Form
}

func CreateView() (v *View) {
	v = new(View)
	v.f = stfl.Create("<screen.stfl>")
	return
}

func(v View) Run() {

	quit := false

	for !quit {
		key := v.f.Run(0)
		switch key {
		case "q":
			quit = true
		// TODO: more keys
		}
	}

	v.f.Free()
	stfl.Reset()
}
