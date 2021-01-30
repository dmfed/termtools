package termtools

import (
	"fmt"

	"golang.org/x/sys/unix"
)

// Cursor position manipulations

func getTermSize() (int, int, error) {
	ws, err := unix.IoctlGetWinsize(0, unix.TIOCGWINSZ)
	if err != nil {
		return -1, -1, ErrUnknownTermSize
	}
	return int(ws.Col), int(ws.Row), nil
}

func moveCursorTo(x, y int) {
	maxx, maxy, _ := getTermSize()
	if x <= maxx && y <= maxy {
		fmt.Printf(CursorGotoTemplate, y, x)
	}
}

func moveCursorUp(rows int) {
	fmt.Printf(CursorMoveUpTemplate, rows)
}

func moveCursorDown(rows int) {
	fmt.Printf(CursorMoveDownTemplate, rows)
}

func moveCursorLeft(columns int) {
	fmt.Printf(CursorMoveLeftTemplate, columns)
}

func moveCursorRight(columns int) {
	fmt.Printf(CursorMoveRightTemplate, columns)
}

func moveCursorToNextRow() {
	fmt.Print(CursorMoveToNextRowTemplate)
}

func moveCursorToRow(row int) {
	fmt.Printf(CursorMoveToRowTemplate, row)
}

func saveCursorPosition() {
	fmt.Print(CursorSave)
}

func restoreCursorPosition() {
	fmt.Print(CursorRestore)
}

// Color functions

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
		return fmt.Sprintf(ColorIDTemplate, id), nil
	}
	return "", ErrUnknownColor
}

func getBackgroundByID(id int) (string, error) {
	if id >= 0 && id < 256 {
		return fmt.Sprintf(BackgroundIDTemplate, id), nil
	}
	return "", ErrUnknownColor
}

// Printing functions

func colorSprint(colorname string, a ...interface{}) string {
	code, err := getColorByName(colorname)
	if err != nil {
		return fmt.Sprint(a...)
	}
	return code + fmt.Sprint(a...) + Reset
}

func colorSprintf(colorname string, format string, a ...interface{}) string {
	code, err := getColorByName(colorname)
	if err != nil {
		return fmt.Sprintf(format, a...)
	}
	return fmt.Sprintf(code+format+Reset, a...)
}

func colorIDSprint(id int, a ...interface{}) string {
	code, err := getColorByID(id)
	if err != nil {
		return fmt.Sprint(a...)
	}
	return code + fmt.Sprint(a...) + Reset
}

func colorIDSprintf(id int, format string, a ...interface{}) string {
	code, err := getColorByID(id)
	if err != nil {
		return fmt.Sprintf(format, a...)
	}
	return fmt.Sprintf(code+format+Reset, a...)
}

func printAtPositionAndReturn(x, y int, a ...interface{}) {
	saveCursorPosition()
	moveCursorTo(x, y)
	fmt.Print(a...)
	restoreCursorPosition()
}

func printAtPosition(x, y int, a ...interface{}) {
	moveCursorTo(x, y)
	fmt.Print(a...)
}
