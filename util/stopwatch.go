package util

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/urfave/cli/v2"
	"os"
	"sync"
	"time"
)

var SmallStopWatchDiff = [11]int{7, 6, 6, 7, 6, 6, 7, 6, 4, 7, 7}
var SmallStopWatchWithoutHourDiff = [8]int{7, 6, 6, 7, 6, 4, 7, 7}

// stopwatchState tracks the running/paused state so the space bar can
// pause, resume, or reset-and-restart the stopwatch from the event loop
// while the render goroutine keeps reading the current elapsed time.
type stopwatchState struct {
	mu        sync.Mutex
	running   bool
	elapsed   time.Duration
	lastStart time.Time
}

func (s *stopwatchState) current() time.Duration {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.running {
		return s.elapsed + time.Since(s.lastStart)
	}
	return s.elapsed
}

// togglePauseResume pauses a running stopwatch or resumes a paused one
// from where it left off.
func (s *stopwatchState) togglePauseResume() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.running {
		s.elapsed += time.Since(s.lastStart)
		s.running = false
	} else {
		s.lastStart = time.Now()
		s.running = true
	}
}

func (s *stopwatchState) resetAndStart() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.elapsed = 0
	s.lastStart = time.Now()
	s.running = true
}

func Stopwatch(cCtx *cli.Context) error {
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

	state := &stopwatchState{running: true, lastStart: time.Now()}

	go func() {
		for {
			termWidth, termHeight := termbox.Size()
			current := state.current().Round(time.Millisecond)
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			NowTime := ""
			if cCtx.String("disable-hour") == "true" {
				NowTime = StopwatchFormatTimeWihtoutHour(current, false)
			}else{
				NowTime = StopwatchFormatTime(current, false)
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
						diff = diff + SmallStopWatchWithoutHourDiff[i-1]
					} else {
						diff = diff + SmallStopWatchDiff[i-1]
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

	var lastSpacePress time.Time

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
			case ev.Key == termbox.KeySpace:
				now := time.Now()
				if !lastSpacePress.IsZero() && now.Sub(lastSpacePress) <= doublePressWindow {
					state.resetAndStart()
					lastSpacePress = time.Time{}
				} else {
					state.togglePauseResume()
					lastSpacePress = now
				}
			}
		case termbox.EventError:
			os.Exit(1)
		}
	}
	return nil
}
