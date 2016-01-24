package main

import (
	"log"
	"strconv"

	api "github.com/gophergala2016/be/insightapi"
	"github.com/gophergala2016/be/tui"
	"github.com/nsf/termbox-go"
)

func tuiLatestBlocks() {
	latestBlocks, err := api.GetLatestBlocks()
	if err != nil {
		log.Fatal(err)
	}

	err = termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	canvas := tui.Canvas{}

	group := tui.Group{}
	for i, block := range latestBlocks.Blocks {
		if i == 0 { // draw unconfirmed block
			box := tui.Box{
				Lines: []string{
					"",
					"  #" + strconv.Itoa(block.Height+1),
				},
				X: 1 + i*18, Y: 1,
				Width: 15, Height: 5,
				Background: termbox.ColorRed, Foreground: termbox.ColorBlack,
			}

			group = append(group, box)
		}

		box := tui.Box{
			Lines: []string{
				"",
				"  #" + strconv.Itoa(block.Height),
				"  " + strconv.Itoa(block.Txlength) + "txs",
				"  " + strconv.Itoa(block.Size/1024) + "kb",
			},
			X: 1 + (i+1)*18, Y: 1,
			Width: 15, Height: 5,
			Background: termbox.ColorBlue, Foreground: termbox.ColorBlack,
		}

		group = append(group, box)
	}

	canvas.Drawable = group
	canvas.Redraw()

	tuiPoll(canvas)
}

func tuiPoll(canvas tui.Canvas) {
	for {
		e := termbox.PollEvent()

		if e.Type == termbox.EventKey {
			return
		}

		if e.Type == termbox.EventResize {
			canvas.Redraw()
		}
	}
}
