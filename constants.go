package termtools

const (
	// The following constants hold ANSI escape sequences with font and background color codes
	// and basic style escapes

	//Basic 8 colors
	Black   string = "\u001b[30m"
	Red            = "\u001b[31m"
	Green          = "\u001b[32m"
	Yellow         = "\u001b[33m"
	Blue           = "\u001b[34m"
	Magenta        = "\u001b[35m"
	Cyan           = "\u001b[36m"
	White          = "\u001b[37m"

	//Additional 8 bright colors
	BrightBlack   string = "\u001b[30;1m"
	BrightRed            = "\u001b[31;1m"
	BrightGreen          = "\u001b[32;1m"
	BrightYellow         = "\u001b[33;1m"
	BrightBlue           = "\u001b[34;1m"
	BrightMagenta        = "\u001b[35;1m"
	BrightCyan           = "\u001b[36;1m"
	BrightWhite          = "\u001b[37;1m"

	//Basic 8 background colors
	BBlack   string = "\u001b[40m"
	BRed            = "\u001b[41m"
	BGreen          = "\u001b[42m"
	BYellow         = "\u001b[43m"
	BBlue           = "\u001b[44m"
	BMagenta        = "\u001b[45m"
	BCyan           = "\u001b[46m"
	BWhite          = "\u001b[47m"

	//Additional 8 bright background colors
	BBrightBlack   string = "\u001b[40;1m"
	BBrightRed            = "\u001b[41;1m"
	BBrightGreen          = "\u001b[42;1m"
	BBrightYellow         = "\u001b[43;1m"
	BBrightBlue           = "\u001b[44;1m"
	BBrightMagenta        = "\u001b[45;1m"
	BBrightCyan           = "\u001b[46;1m"
	BBrightWhite          = "\u001b[47;1m"

	//Styles (can be used separately or together with color and background codes)
	Bold      string = "\u001b[1m"
	Underline        = "\u001b[4m"
	Reversed         = "\u001b[7m"

	//Reset escape sequence
	Reset string = "\u001b[0m"

	// SaveCursor - code to save cursor position
	SaveCursor string = "\033[s"

	//RestoreCursor code to restore cursor position
	RestoreCursor string = "\033[u"

	// Clear screen Codes
	Clear     string = "\u001b[2J"
	ClearUp          = "\u001b[1J"
	ClearDown        = "\u001b[0J"

	// Clear line codes
	ClearL      string = "\u001b[2K"
	ClearLLeft         = "\u001b[1K"
	ClearLRight        = "\u001b[0K"

	// Format string to move cursor. Needs two values: y and x (in this order).
	MoveTemplate = "\033[%v;%vH"
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
)
