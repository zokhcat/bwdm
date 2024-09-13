package graph

import (
	"fmt"
	"math"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func DrawGraph(byteSlice []uint64) {
	app := tview.NewApplication()

	graph := tview.NewBox().SetBorder(true).SetTitle("Packet Count Graph")

	maxVal := uint64(0)

	for _, v := range byteSlice {
		if v > maxVal {
			maxVal = v
		}
	}

	draw := func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
		for i, val := range byteSlice {
			scaledHeight := int(math.Round(float64(val) / float64(maxVal) * float64(height-2)))

			for j := 0; j < scaledHeight; j++ {
				screen.SetContent(x+i*2, y+height-2-j, '█', nil, tcell.StyleDefault.Foreground(tcell.ColorGreen))
			}

			valStr := fmt.Sprintf("%d", val)
			for k, ch := range valStr {
				screen.SetContent(x+i*2+k, y+height-3-scaledHeight, rune(ch), nil, tcell.StyleDefault.Foreground(tcell.ColorYellow))
			}
		}

		for i := range byteSlice {
			screen.SetContent(x+i*2, y+height-1, rune('0'+i), nil, tcell.StyleDefault)
		}

		return x, y, width, height
	}

	graph.SetDrawFunc(draw)

	if err := app.SetRoot(graph, true).Run(); err != nil {
		panic(err)
	}
}
