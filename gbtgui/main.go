// To build this on Debian and Ubuntu, install libgtk-3-dev first.

package main

import (
	"fmt"
	"time"

	"github.com/andlabs/ui"
	"github.com/brandondube/tai"
)

func main() {
	err := ui.Main(func() {
		w := ui.NewWindow("beatTAI GUI", 205, 24, false)
		w.SetMargined(true)

		l := ui.NewLabel("")
		w.SetChild(l)

		w.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})

		go func() {
			ticker := time.NewTicker(500 * time.Millisecond)
			defer ticker.Stop()

			for now := range ticker.C {
				ui.QueueMain(func() {
					l.SetText(timeInTaiString(now))
				})
			}
		}()

		w.Show()
	})
	if err != nil {
		panic(err)
	}
}

func timeInTaiString(local time.Time) string {
	g := tai.FromTime(local).AsGregorian()
	hourSeconds := g.Hour * 3600
	minuteSeconds := g.Min * 60
	seconds := g.Sec
	beatTai := float64(hourSeconds+minuteSeconds+seconds) / 86.4

	return fmt.Sprintf("@%06.2f (%02d:%02d:%02d)", beatTai, local.Hour(), local.Minute(), local.Second())
}
