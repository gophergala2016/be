package main

import (
	"log"
	"os"
	"strconv"

	api "github.com/gophergala2016/be/insightapi"
	"github.com/gophergala2016/be/tui"
	"github.com/nsf/termbox-go"
)

var (
	state        api.BlockList
	selected     int
	screenWidth  int
	screenHeight int
	screenScroll int
)

const (
	boxWidth  = 19
	boxHeight = 7
	xMargin   = 2
	yMargin   = 1
	xSpace    = 4
	ySpace    = 2
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

	selected = 1
	screenWidth, screenHeight = termbox.Size()
	screenScroll = 0

	draw()
	tuiPoll()
	termbox.Close()
	cliBlock(state.Blocks[selected-1].Hash)
	os.Exit(0)
}

func box(lines []string, x, y int, background, foreground termbox.Attribute) tui.Box {
	return tui.Box{
		Lines: lines,
		X:     xMargin + x*(boxWidth+xSpace), Y: yMargin + (y-screenScroll)*(boxHeight+ySpace),
		Width: boxWidth, Height: boxHeight,
		Background: background, Foreground: foreground,
	}
}

func horizontalLine(x, y int) tui.Box {
	line := ""
	if y%2 != 0 {
		line = line + "<"
	}
	for i := 0; i < xSpace-1; i++ {
		line = line + "─"
	}
	if y%2 == 0 {
		line = line + ">"
	}

	return tui.Box{
		Lines: []string{line},
		X:     xMargin + boxWidth + (xSpace+boxWidth)*x,
		Y:     yMargin + boxHeight/2 + (ySpace+boxHeight)*(y-screenScroll),
		Width: xSpace, Height: 1,
		Foreground: termbox.ColorWhite,
	}
}

func verticalLine(x, y int) tui.Box {
	lines := []string{}
	for i := 0; i < ySpace-1; i++ {
		lines = append(lines, "│")
	}
	lines = append(lines, "V")

	return tui.Box{
		Lines: lines,
		X:     xMargin + boxWidth/2 + (xSpace+boxWidth)*x,
		Y:     yMargin + boxHeight + (ySpace+boxHeight)*(y-screenScroll),
		Width: 1, Height: ySpace,
		Foreground: termbox.ColorWhite,
	}
}

func calculateFit(pad, space, boxSize, containerSize int) (boxes int) {
	for {
		if pad+boxSize*(boxes+1)+space*boxes+pad > containerSize {
			return
		}

		boxes = boxes + 1
	}
}

func calculateXFit() int {
	x := calculateFit(xMargin, xSpace, boxWidth, screenWidth)
	if x < 1 {
		x = 1
	}
	return x
}

func calculateYFit() int {
	x := calculateFit(yMargin, ySpace, boxHeight, screenHeight)
	if x < 1 {
		x = 1
	}
	return x
}

func toSnake(i int) (x, y int) {
	xBoxes := calculateXFit()

	y = i / xBoxes

	if y%2 == 0 {
		x = i % xBoxes
	} else {
		x = xBoxes - 1 - (i % xBoxes)
	}

	return
}

func blockBox(block api.BlockInfo, i int) tui.Group {
	xBoxes := calculateXFit()

	x, y := toSnake(i)

	color := termbox.ColorBlack
	if i == selected {
		color = termbox.ColorWhite
	}

	box := box(
		[]string{
			"",
			"  #" + strconv.Itoa(block.Height),
			"",
			"  " + strconv.Itoa(block.Txlength) + " txs",
			"  " + strconv.Itoa(block.Size/1024) + " kb",
			"  " + block.PoolInfo.PoolName,
		},
		x, y, termbox.ColorBlue, color,
	)

	var line tui.Drawable

	if y%2 == 0 {
		if x == 0 {
			line = verticalLine(x, y-1)
		} else {
			line = horizontalLine(x-1, y)
		}
	} else {
		if x == xBoxes-1 {
			line = verticalLine(x, y-1)
		} else {
			line = horizontalLine(x, y)
		}
	}

	return tui.Group{box, line}
}

func nextBlockBox(block api.BlockInfo) tui.Box {
	return box(
		[]string{
			"",
			"  #" + strconv.Itoa(block.Height+1),
			"",
			"  next",
			"  block",
		},
		0, 0, termbox.ColorRed, termbox.ColorBlack,
	)
}

func draw() {
	canvas := tui.Canvas{}

	group := tui.Group{}
	for i, block := range state.Blocks {
		if i == 0 { // draw unconfirmed block
			group = append(group, nextBlockBox(block))
		}

		group = append(group, blockBox(block, i+1))
	}

	canvas.Drawable = group
	canvas.Redraw()
}

func move(x int) {
	selected = selected + x
	if selected < 1 {
		selected = 1
	}
	if selected > len(state.Blocks) {
		selected = len(state.Blocks)
	}
	for _, y := toSnake(selected); y+1 > calculateYFit()+screenScroll; {
		screenScroll = screenScroll + 1
	}
	for _, y := toSnake(selected); y < screenScroll; {
		screenScroll = screenScroll - 1
	}
	draw()
}

func tuiPoll() {
	for {
		e := termbox.PollEvent()

		if e.Type == termbox.EventKey {
			switch e.Key {
			case termbox.KeyArrowLeft:
				move(-1)
			case termbox.KeyArrowUp:
				move(-1)
			case termbox.KeyArrowRight:
				move(1)
			case termbox.KeyArrowDown:
				move(1)
			case termbox.KeyEnter:
				return
			default:
				termbox.Close()
				os.Exit(0)
			}
		}

		if e.Type == termbox.EventResize {
			screenWidth = e.Width
			screenHeight = e.Height
			draw()
			move(0)
		}
	}
}
