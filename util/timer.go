package util

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
	"sync"
	"time"
)

var SmallTimerDiff = [11]int{7, 6, 6, 7, 6, 6, 7, 6, 4, 7, 7}
var SmallTimerWithoutHourDiff = [11]int{7, 6, 6, 7, 6, 4, 7, 7}

var SmallTimerWithoutHourAndMillisecondDiff = [11]int{7, 6, 6, 7, 6}
var SmallTimerWithoutMillisecondDiff = [11]int{7, 6, 6, 7, 6, 6, 7, 6}

// timerState tracks the running/paused state so the space bar can pause,
// resume, or reset-and-restart the countdown from the event loop while the
// render goroutine keeps reading the current remaining time.
type timerState struct {
	mu        sync.Mutex
	running   bool
	elapsed   time.Duration
	lastStart time.Time
}

func (s *timerState) remaining(total time.Duration) time.Duration {
	s.mu.Lock()
	defer s.mu.Unlock()
	elapsed := s.elapsed
	if s.running {
		elapsed += time.Since(s.lastStart)
	}
	return total - elapsed
}

// togglePauseResume pauses a running countdown or resumes a paused one
// from where it left off.
func (s *timerState) togglePauseResume() {
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

func (s *timerState) resetAndStart() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.elapsed = 0
	s.lastStart = time.Now()
	s.running = true
}

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

	if cCtx.String("time") != "" {
		seconds, _ := timeStringToSeconds(cCtx.String("time"))
		total = total + seconds
	}

	if total == 0 {
		total = 300
	}
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	colonColor := FlagColor(cCtx.String("colon-color"))

	if cCtx.String("colon-color") == "" && cCtx.String("color") != "" {
		colonColor = FlagColor(cCtx.String("color"))
	}
	duration := time.Duration(total) * time.Second
	state := &timerState{running: true, lastStart: time.Now()}

	go func() {
		for {
			termWidth, termHeight := termbox.Size()
			current := state.remaining(duration)
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			NowTime := ""

			disableMillisecond := cCtx.String("disable-millisecond") == "true"
			if cCtx.String("disable-hour") == "true" {

				if current <= 0 {
					NowTime = "00:00.000"
				} else {
					NowTime = StopwatchFormatTimeWihtoutHour(current, disableMillisecond)
				}
			} else {
				if current <= 0 {
					NowTime = "00:00:00.000"
				} else {
					NowTime = StopwatchFormatTime(current, disableMillisecond)
				}
			}
			diff := -38
			totalString := 12
			if cCtx.String("disable-hour") == "true" {
				if disableMillisecond {
					totalString = 5
					diff = -16
				} else {
					totalString = 9
					diff = -29
				}
			} else if disableMillisecond {
				totalString = 8
				diff = -25
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
