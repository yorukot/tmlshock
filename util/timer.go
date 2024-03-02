package util

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
	"time"
)

var SmallTimerDiff = [11]int{7, 6, 6, 7, 6, 6, 7, 6, 4, 7, 7}
var SmallTimerWithoutHourDiff = [11]int{7, 6, 6, 7, 6, 4, 7, 7}

func Timer(cCtx *cli.Context) error {
	err := termbox.Init()
	if err != nil {
		fmt.Println("failed to initialize termbox:", err)
		return err
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)
	color := FlagColor(cCtx.String("color"))

	min, _ := strconv.Atoi(cCtx.String("minute"))
	sec, _ := strconv.Atoi(cCtx.String("second"))
	hour, _ := strconv.Atoi(cCtx.String("hour"))

	total := hour*60*60 + min*60 + sec*60
	if total == 0 {
		total = 300
	}
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	colonColor := FlagColor(cCtx.String("colon-color"))

	if cCtx.String("colon-color") == "" && cCtx.String("color") != "" {
		colonColor = FlagColor(cCtx.String("color"))
	}
	go func() {
		duration := time.Duration(total) * time.Second
		end := time.Now().Add(duration)
		for {

			termWidth, termHeight := termbox.Size()
			current := time.Until(end)
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			NowTime := ""
			if cCtx.String("disable-hour") == "true" {
				NowTime = StopwatchFormatTimeWihtoutHour(current)
			} else {
				NowTime = StopwatchFormatTime(current)
			}
			diff := -38
			totalString := 12
			if cCtx.String("disable-hour") == "true" {
				totalString = 9
				diff = -29
			}
			for i := 0; i < totalString; i++ {
				if i != 0 {
					if cCtx.String("disable-hour") == "true" {
						diff = diff + SmallTimerWithoutHourDiff[i-1]
					} else {
						diff = diff + SmallTimerDiff[i-1]
					}
				}
				if i == 2 || i == 5 {
					CaseNumber(termWidth, termHeight, colonColor, NowTime[i], diff)
				} else {
					CaseNumber(termWidth, termHeight, color, NowTime[i], diff)
				}
			}

			termbox.Flush()
			time.Sleep(time.Millisecond)
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
