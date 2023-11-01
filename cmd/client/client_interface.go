package main

import (
	"image/color"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/text"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := app.NewWindow()

		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	th := material.NewTheme()

	var ops op.Ops
	var download widget.Clickable

	for {
		e := <-w.Events()
		switch e := e.(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			title := material.H5(th, "Raindrops")
			maroon := color.NRGBA{R: 217, G: 217, B: 217, A: 255}
			title.Color = maroon
			title.Alignment = text.Middle
			paint.Fill(&ops, color.NRGBA{R: 107, G: 130, B: 158, A: 255})
			title.Layout(gtx)

			downloadBth := material.Button(th, &download, "Download")
			downloadBth.Layout(gtx)

			e.Frame(gtx.Ops)
		}
	}
}
