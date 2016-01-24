package main

import "github.com/nsf/termbox-go"

type Drawable interface {
	Draw()
}

type Box struct {
	Lines                  []string
	Foreground, Background termbox.Attribute
	Width, Height          int
	X, Y                   int
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

type Cell struct {
	Rune                   rune
	Foreground, Background termbox.Attribute
	X, Y                   int
}

func (c Cell) Draw() {
	termbox.SetCell(c.X, c.Y, c.Rune, c.Foreground, c.Background)
}

type Group []Drawable

func (g Group) Draw() {
	for _, drawable := range g {
		drawable.Draw()
	}
}

type Canvas struct {
	Foreground, Background termbox.Attribute
	Drawable               Drawable
}

func (c Canvas) Redraw() {
	termbox.Clear(c.Foreground, c.Background)
	c.Drawable.Draw()
	termbox.Flush()
}
