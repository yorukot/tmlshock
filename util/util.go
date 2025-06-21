package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/nsf/termbox-go"
)

func StopwatchFormatTime(d time.Duration, disableMillisecond bool) string {
    h := d / time.Hour
    d -= h * time.Hour
    m := d / time.Minute
    d -= m * time.Minute
    s := d / time.Second
    if disableMillisecond {
        return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
    }
    ms := d % time.Second / time.Millisecond
    return fmt.Sprintf("%02d:%02d:%02d.%03d", h, m, s, ms)
}

func StopwatchFormatTimeWihtoutHour(d time.Duration, disableMillisecond bool) string {
    m := d / time.Minute
    d -= m * time.Minute
    s := d / time.Second
    if disableMillisecond {
        return fmt.Sprintf("%02d:%02d", m, s)
    }
    ms := d % time.Second / time.Millisecond
    return fmt.Sprintf("%02d:%02d.%03d", m, s, ms)
}

func formatTime(d time.Time, use12HourFormat bool) string {
	var hours int
	if use12HourFormat {
		hours = d.Hour() % 12
		if hours == 0 {
			hours = 12
		}
	} else {
		hours = d.Hour()
	}
	minutes := d.Minute()
	seconds := d.Second()
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func drawString(x, y int, text string, fg, bg termbox.Attribute) {
	for _, c := range text {
		termbox.SetCell(x, y, c, fg, bg)
		x++
	}
}

func FlagColor(ColorString string) termbox.Attribute {
	switch ColorString {
	case "black":
		return termbox.ColorBlack
	case "blue":
		return termbox.ColorBlue
	case "cyan":
		return termbox.ColorCyan
	case "dark-gray":
		return termbox.ColorDarkGray
	case "green":
		return termbox.ColorGreen
	case "light-green":
		return termbox.ColorLightGreen
	case "light-blue":
		return termbox.ColorLightBlue
	case "light-cyan":
		return termbox.ColorLightCyan
	case "light-gray":
		return termbox.ColorLightGray
	case "light-magenta":
		return termbox.ColorLightMagenta
	case "light-red":
		return termbox.ColorLightRed
	case "light-yellow":
		return termbox.ColorLightYellow
	case "magenta":
		return termbox.ColorMagenta
	case "red":
		return termbox.ColorRed
	case "white":
		return termbox.ColorWhite
	case "yellow":
		return termbox.ColorYellow
	default:
		if ColorString == "" {
			return termbox.ColorGreen
		} else {
			num, err := strconv.Atoi(ColorString)
			if err != nil || num > 256 || num < 1 {
				return termbox.ColorGreen
			}
			return termbox.Attribute(num)
		}
	}
}

func timeStringToSeconds(timeStr string) (int, error) {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 3 {
		return 0, fmt.Errorf("时间格式错误")
	}

	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	minutes, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}

	seconds, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, err
	}

	totalSeconds := hours*3600 + minutes*60 + seconds
	return totalSeconds, nil
}
