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
			ticker := time.NewTicker(100 * time.Millisecond)
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

	// h, m, s, and ms are in milliseconds.
	h := g.Hour * 3_600_000
	m := g.Min * 60_000
	s := g.Sec * 1_000
	ms := g.Asec / 1_000_000_000_000_000

	beatTAI := (float64(h+m+s) + float64(ms)) / 86400

	return fmt.Sprintf(":%06.2f (%02d:%02d:%02d)", beatTAI, local.Hour(), local.Minute(), local.Second())
}
