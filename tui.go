package main

import (
	"log"
	"strconv"

	api "github.com/gophergala2016/be/insightapi"
	"github.com/gophergala2016/be/tui"
	"github.com/nsf/termbox-go"
)

var (
	state api.BlockList
)

func tuiLatestBlocks() {
	var err error
	state, err = api.GetLatestBlocks()
	if err != nil {
		log.Fatal(err)
	}

	err = termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	draw()
	tuiPoll()
}

func box(lines []string, x, y int, background termbox.Attribute) tui.Box {
	return tui.Box{
		Lines: lines,
		X:     1 + x*18, Y: 1 + y*8,
		Width: 15, Height: 5,
		Background: background, Foreground: termbox.ColorBlack,
	}
}

func blockBox(block api.BlockInfo, i int) tui.Box {
	return box(
		[]string{
			"",
			"  #" + strconv.Itoa(block.Height),
			"  " + strconv.Itoa(block.Txlength) + "txs",
			"  " + strconv.Itoa(block.Size/1024) + "kb",
		},
		i+1, 0, termbox.ColorBlue,
	)
}

func unconfirmedBlockBox(block api.BlockInfo) tui.Box {
	return box(
		[]string{
			"",
			"  #" + strconv.Itoa(block.Height+1),
		},
		0, 0, termbox.ColorRed,
	)
}

func draw() {
	canvas := tui.Canvas{}

	group := tui.Group{}
	for i, block := range state.Blocks {
		if i == 0 { // draw unconfirmed block
			group = append(group, unconfirmedBlockBox(block))
		}

		group = append(group, blockBox(block, i))
	}

	canvas.Drawable = group
	canvas.Redraw()
}

func tuiPoll() {
	for {
		e := termbox.PollEvent()

		if e.Type == termbox.EventKey {
			return
		}

		if e.Type == termbox.EventResize {
			draw()
		}
	}
}
