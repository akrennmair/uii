package main

type LineMsg struct {
	Line string
	Channel string
}

type WindowManager struct {
	LineQueue chan LineMsg
	view *View
}

func CreateWindowManager(view *View, lq chan LineMsg) *WindowManager {
	wman := new(WindowManager)
	wman.LineQueue = lq
	wman.view = view
	return wman
}

func(wman *WindowManager) Run() {
	for {
		msg := <-wman.LineQueue
		wman.view.AddLine(msg.Line)
		wman.view.UpdateScreen()
		// TODO: implement
	}
}
