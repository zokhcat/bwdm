package graph

// import (
// 	"fmt"
// 	"math"

// 	"github.com/gdamore/tcell/v2"
// 	"github.com/rivo/tview"
// )

// type GraphWidget struct {
// 	*tview.Box
// 	data     []float64
// 	maxValue float64
// }

// func NewGraphWidget() *GraphWidget {
// 	return &GraphWidget{
// 		Box:      tview.NewBox().SetTitle("Graph"),
// 		data:     make([]float64, 10),
// 		maxValue: 1,
// 	}
// }

// func (g *GraphWidget) DrawGraph(screen tcell.Screen) {
// 	g.Box.DrawForSubclass(screen, g)
// 	x, y, width, height := g.GetRect()

// 	for i := 0; i < height; i++ {
// 		screen.SetContent(x, y+i, '|', nil, tcell.StyleDefault)
// 	}

// 	for i := 0; i < width; i++ {
// 		screen.SetContent(x+i, y+height-1, '-', nil, tcell.StyleDefault)
// 	}

// 	for i, value := range g.data {
// 		if value > 0 {
// 			normalizedValue := int(math.Floor(value / g.maxValue * float64(height-2)))
// 			screen.SetContent(x+i+1, y+height-2-normalizedValue, '•', nil, tcell.StyleDefault.Foreground(tcell.ColorGreen))
// 		}
// 	}

// 	tview.Print(screen, "Bytes", x+1, y, 5, tview.AlignLeft, tcell.ColorWhite)
// 	tview.Print(screen, fmt.Sprintf("%.0f", g.maxValue), x+1, y+1, 5, tview.AlignLeft, tcell.ColorWhite)
// 	tview.Print(screen, "Time(s)", x+width-8, y+height-1, 8, tview.AlignRight, tcell.ColorWhite)
// 	tview.Print(screen, "60", x+width-2, y+height-2, 2, tview.AlignRight, tcell.ColorWhite)
// }

// func (g *GraphWidget) AddPoint(value float64) {
// 	g.data = append(g.data[1:], value)
// 	if value > g.maxValue {
// 		g.maxValue = value
// 	}
// }
