package termtools

import "errors"

const (
	// The following constants hold ANSI escape sequences setting color and style of output
	Esc string = "\x1b"

	//Basic 8 colors
	Black   string = Esc + "[30m"
	Red            = Esc + "[31m"
	Green          = Esc + "[32m"
	Yellow         = Esc + "[33m"
	Blue           = Esc + "[34m"
	Magenta        = Esc + "[35m"
	Cyan           = Esc + "[36m"
	White          = Esc + "[37m"

	//Additional 8 bright colors
	BrightBlack   string = Esc + "[30;1m"
	BrightRed            = Esc + "[31;1m"
	BrightGreen          = Esc + "[32;1m"
	BrightYellow         = Esc + "[33;1m"
	BrightBlue           = Esc + "[34;1m"
	BrightMagenta        = Esc + "[35;1m"
	BrightCyan           = Esc + "[36;1m"
	BrightWhite          = Esc + "[37;1m"

	//Basic 8 background colors
	BBlack   string = Esc + "[40m"
	BRed            = Esc + "[41m"
	BGreen          = Esc + "[42m"
	BYellow         = Esc + "[43m"
	BBlue           = Esc + "[44m"
	BMagenta        = Esc + "[45m"
	BCyan           = Esc + "[46m"
	BWhite          = Esc + "[47m"

	//Additional 8 bright background colors
	BBrightBlack   string = Esc + "[40;1m"
	BBrightRed            = Esc + "[41;1m"
	BBrightGreen          = Esc + "[42;1m"
	BBrightYellow         = Esc + "[43;1m"
	BBrightBlue           = Esc + "[44;1m"
	BBrightMagenta        = Esc + "[45;1m"
	BBrightCyan           = Esc + "[46;1m"
	BBrightWhite          = Esc + "[47;1m"

	// Color format string to use with 256 color codes. Needs int in range [0;255].
	ColorIDTemplate      string = Esc + "[38;5;%vm"
	BackgroundIDTemplate        = Esc + "[48;5;%vm"

	//Styles. Can be used separately or together with color and background codes.
	Bold      string = Esc + "[1m"
	Underline        = Esc + "[4m"
	Blinking         = Esc + "[5m"
	Reversed         = Esc + "[7m"

	// Reset escape sequence
	Reset string = Esc + "[0m"

	// Cursor maniputation
	CursorSave                  string = Esc + "[s"
	CursorRestore                      = Esc + "[u"
	CursorGotoTemplate                 = Esc + "[%v;%vH"
	CursorMoveUpTemplate               = Esc + "[%vA"
	CursorMoveDownTemplate             = Esc + "[%vB"
	CursorMoveRightTemplate            = Esc + "[%vC"
	CursorMoveLeftTemplate             = Esc + "[%vD"
	CursorMoveToNextRowTemplate        = Esc + "[E"
	CursorMoveToRowTemplate            = Esc + "[%vH"

	// Clear screen codes
	Clear     string = Esc + "[2J"
	ClearUp          = Esc + "[1J"
	ClearDown        = Esc + "[0J"

	// Clear line codes
	ClearL      string = Esc + "[2K"
	ClearLLeft         = Esc + "[1K"
	ClearLRight        = Esc + "[0K"
)

var (
	ErrUnknownColor    = errors.New("error: unknown color name or color id out of range [0;255]")
	ErrUnknownTermSize = errors.New("error: could not find out terminal size")
)

var (
	colorMap = map[string]string{
		"black":         Black,
		"red":           Red,
		"green":         Green,
		"yellow":        Yellow,
		"blue":          Blue,
		"magenta":       Magenta,
		"cyan":          Cyan,
		"white":         White,
		"brightblack":   BrightBlack,
		"brightred":     BrightRed,
		"brightgreen":   BrightGreen,
		"brightyellow":  BrightYellow,
		"brightblue":    BrightBlue,
		"brightmagenta": BrightMagenta,
		"brightcyan":    BrightCyan,
		"brightwhite":   BrightWhite}
	backgroundMap = map[string]string{
		"black":         BBlack,
		"red":           BRed,
		"green":         BGreen,
		"yellow":        BYellow,
		"blue":          BBlue,
		"magenta":       BMagenta,
		"cyan":          BCyan,
		"white":         BWhite,
		"brightblack":   BBrightBlack,
		"brightred":     BBrightRed,
		"brightgreen":   BBrightGreen,
		"brightyellow":  BBrightYellow,
		"brightblue":    BBrightBlue,
		"brightmagenta": BBrightMagenta,
		"brightcyan":    BBrightCyan,
		"brightwhite":   BBrightWhite}
	modeMap = map[string]string{
		"bold":      Bold,
		"underline": Underline,
		"blinking":  Blinking,
		"reversed":  Reversed}
)
