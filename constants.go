package termtools

const (
	//The following constans hold ANSI escape sequences with font and background color codes
	//and basic style codes
	//Basic 8 colors
	Black   = "\u001b[30m"
	Red     = "\u001b[31m"
	Green   = "\u001b[32m"
	Yellow  = "\u001b[33m"
	Blue    = "\u001b[34m"
	Magenta = "\u001b[35m"
	Cyan    = "\u001b[36m"
	White   = "\u001b[37m"
	//Additional 8 bright colors
	BrightBlack   = "\u001b[30;1m"
	BrightRed     = "\u001b[31;1m"
	BrightGreen   = "\u001b[32;1m"
	BrightYellow  = "\u001b[33;1m"
	BrightBlue    = "\u001b[34;1m"
	BrightMagenta = "\u001b[35;1m"
	BrightCyan    = "\u001b[36;1m"
	BrightWhite   = "\u001b[37;1m"
	//Basic 8 background colors
	BBlack   = "\u001b[40m"
	BRed     = "\u001b[41m"
	BGreen   = "\u001b[42m"
	BYellow  = "\u001b[43m"
	BBlue    = "\u001b[44m"
	BMagenta = "\u001b[45m"
	BCyan    = "\u001b[46m"
	BWhite   = "\u001b[47m"
	//Additional 8 bright background colors
	BBrightBlack   = "\u001b[40;1m"
	BBrightRed     = "\u001b[41;1m"
	BBrightGreen   = "\u001b[42;1m"
	BBrightYellow  = "\u001b[43;1m"
	BBrightBlue    = "\u001b[44;1m"
	BBrightMagenta = "\u001b[45;1m"
	BBrightCyan    = "\u001b[46;1m"
	BBrightWhite   = "\u001b[47;1m"
	//Styles (can be used separately ot together)
	Bold      = "\u001b[1m"
	Underline = "\u001b[4m"
	Reversed  = "\u001b[7m"
	//Reset escape sequence
	ColorReset = "\u001b[0m"
	//Save cursor position
	SaveCursor = "\033[s"
	// restore cursor position
	RestoreCursor = "\033[u"
)
