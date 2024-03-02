package util

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

var SmallStopWatchDiff = [11]int{7, 6, 6, 7, 6, 6, 7, 6, 4, 7, 7}

func Stopwatch(cCtx *cli.Context) error {
	err := termbox.Init()
	if err != nil {
		fmt.Println("failed to initialize termbox:", err)
		return err
	}
	defer termbox.Close()
	termbox.SetOutputMode(termbox.Output256)
	color := FlagColor(cCtx.String("color"))

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	go func() {
		start := time.Now()
		for {
			termWidth, termHeight := termbox.Size()
			current := time.Since(start).Round(time.Millisecond)
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			NowTime := StopwatchFormatTime(current)
			diff := -38
			for i := 0; i < 12; i++ {
				if i != 0 {
					diff = diff + SmallStopWatchDiff[i-1]
				}
				switch NowTime[i] {
				case '0':
					SmallZero(termWidth, termHeight, color, diff)
				case '1':
					SmallOne(termWidth, termHeight, color, diff)
				case '2':
					SmallTwo(termWidth, termHeight, color, diff)
				case '3':
					SmallThree(termWidth, termHeight, color, diff)
				case '4':
					SmallFour(termWidth, termHeight, color, diff)
				case '5':
					SmallFive(termWidth, termHeight, color, diff)
				case '6':
					SmallSix(termWidth, termHeight, color, diff)
				case '7':
					SmallSeven(termWidth, termHeight, color, diff)
				case '8':
					SmallEight(termWidth, termHeight, color, diff)
				case '9':
					SmallNine(termWidth, termHeight, color, diff)
				case ':':
					SmallColon(termWidth, termHeight, color, diff)
				case '.':
					SmallDot(termWidth, termHeight, color, diff)
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
