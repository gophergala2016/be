package main

import "github.com/nsf/termbox-go"

type Drawable interface {
	Draw()
}

type Box struct {
	Lines      []string
	Foreground termbox.Attribute
	Background termbox.Attribute
	Width      int
	Height     int
	X          int
	Y          int
}

func (b Box) Draw() {
	for j := 0; j < b.Height; j++ {
		var line []rune
		if j < len(b.Lines) {
			line = []rune(b.Lines[j])
		}

		for i := 0; i < b.Width; i++ {
			var r rune
			if i < len(line) {
				r = line[i]
			} else {
				r = ' '
			}

			termbox.SetCell(b.X+i, b.Y+j, r, b.Foreground, b.Background)
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	boxes := []Box{
		{
			Lines:      []string{"", "  #394650"},
			Background: termbox.ColorRed,
			Foreground: termbox.ColorBlack,
			Width:      15,
			Height:     5,
			X:          1,
			Y:          1,
		},
		{
			Lines:      []string{"▶︎"},
			Foreground: termbox.ColorRed,
			Width:      1,
			Height:     1,
			X:          17,
			Y:          3,
		},
		{
			Lines:      []string{"", "  #394649", "  889 txs", "  927Kb"},
			Background: termbox.ColorYellow,
			Foreground: termbox.ColorBlack,
			Width:      15,
			Height:     5,
			X:          19,
			Y:          1,
		},
		{
			Lines:      []string{"▶︎"},
			Foreground: termbox.ColorYellow,
			Width:      1,
			Height:     1,
			X:          35,
			Y:          3,
		},
		{
			Lines:      []string{"", "  #394648", "  2 txs", "  0.5Kb"},
			Background: termbox.ColorYellow,
			Foreground: termbox.ColorBlack,
			Width:      15,
			Height:     5,
			X:          37,
			Y:          1,
		},
		{
			Lines:      []string{"▶︎"},
			Foreground: termbox.ColorYellow,
			Width:      1,
			Height:     1,
			X:          53,
			Y:          3,
		},
		{
			Lines:      []string{"", "  #394647", "  976 txs", "  966Kb"},
			Background: termbox.ColorYellow,
			Foreground: termbox.ColorBlack,
			Width:      15,
			Height:     5,
			X:          55,
			Y:          1,
		},
		{
			Lines:      []string{"▼"},
			Foreground: termbox.ColorYellow,
			Width:      1,
			Height:     1,
			X:          62,
			Y:          7,
		},
	}

	for _, box := range boxes {
		box.Draw()
	}

	termbox.Flush()

	pollExit()
}

func pollExit() {
	for {
		e := termbox.PollEvent()

		if e.Type == termbox.EventKey {
			return
		}
	}
}
