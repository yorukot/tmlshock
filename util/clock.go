package util

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

var ClockDiff = [11]int{7, 6, 6, 7}
var ClockWithSecDiff = [11]int{7, 6, 6, 7, 6, 6, 7}

func Clock(cCtx *cli.Context) error {
	err := termbox.Init()
	if err != nil {
		fmt.Println("failed to initialize termbox:", err)
		return err
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)
	color := FlagColor(cCtx.String("color"))
	colonColor := FlagColor(cCtx.String("colon-color"))

	if cCtx.String("colon-color") == "" && cCtx.String("color") != "" {
		colonColor = FlagColor(cCtx.String("color"))
	}
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	use12HourFormat := false
	if (cCtx.String("color") == "hour-format"){
		use12HourFormat = true
	}
	go func() {
		for {
			termWidth, termHeight := termbox.Size()
			current := time.Now()
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

			NowTime := formatTime(current, use12HourFormat)

			diff := 0
			totalString := 0
			if cCtx.String("second") == "true" {
				totalString = 8
				diff = -26
			} else {
				totalString = 5
				diff = -17
			}

			for i := 0; i < totalString; i++ {
				if i != 0 {
					diff = diff + SmallStopWatchDiff[i-1]
				}
				if i == 2 || i == 5 {
					CaseNumber(termWidth, termHeight, colonColor, NowTime[i], diff)
				} else {
					CaseNumber(termWidth, termHeight, color, NowTime[i], diff)
				}
			}

			if cCtx.String("date") == "true" {
				dateString := current.Format(cCtx.String("dateformate"))
				drawString(termWidth/2-(len(dateString)/2), termHeight/2+4, dateString, color, termbox.ColorDefault)
			}
			termbox.Flush()
			time.Sleep(time.Second)
		}
	}()

mainLoop:
	for {
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			switch {
			case ev.Key == termbox.KeyEsc || ev.Ch == 'q' || ev.Ch == 'Q':
				break mainLoop
			case ev.Key == termbox.KeyCtrlC:
				break mainLoop
			}
		case termbox.EventError:
			os.Exit(1)
		}
	}
	return nil
}
