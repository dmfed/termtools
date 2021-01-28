package termtools

import (
	"errors"
	"fmt"

	"golang.org/x/sys/unix"
)

var (
	// ErrUnknownColor is returned whenever wrong color name or numeric id is requested
	ErrUnknownColor = errors.New("error: unknown color name or id out of range [0;255]")
)

func getTermSize() (int, int) {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		return -1, -1
	}
	return int(ws.Col), int(ws.Row)
}

func moveCursor(x, y int) {
	maxx, maxy := getTermSize()
	if x <= maxx && y <= maxy {
		fmt.Printf(MoveTemplate, y, x)
	}
}

func getColorByName(colorname string) (string, error) {
	if code, ok := colorMap[colorname]; ok {
		return code, nil
	}
	return "", ErrUnknownColor
}

func getBackgroundByName(colorname string) (string, error) {
	if code, ok := backgroundMap[colorname]; ok {
		return code, nil
	}
	return "", ErrUnknownColor
}

func getColorByID(id int) (string, error) {
	if id >= 0 && id < 256 {
		return fmt.Sprintf("\u001b[38;5;%vm", id), nil
	}
	return "", ErrUnknownColor
}

func getBackgroundByID(id int) (string, error) {
	if id >= 0 && id < 256 {
		return fmt.Sprintf("\u001b[48;5;%vm", id), nil
	}
	return "", ErrUnknownColor
}

func colorSprint(colorname string, a ...interface{}) string {
	code, err := getColorByName(colorname)
	if err != nil {
		return fmt.Sprint(a...)
	}
	return code + fmt.Sprint(a...) + Reset
}

func colorIDSprint(id int, a ...interface{}) string {
	code, err := getColorByID(id)
	if err != nil {
		return fmt.Sprint(a...)
	}
	return code + fmt.Sprint(a...) + Reset
}

func printAtPositionAndReturn(x, y int, a ...interface{}) {
	fmt.Print(SaveCursor)
	moveCursor(x, y)
	fmt.Print(a...)
	fmt.Print(RestoreCursor)
}
