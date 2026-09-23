package main

const (
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorBlue    = "\033[34m"
	colorMagenta = "\033[35m"
	colorCyan    = "\033[36m"
	colorWhite   = "\033[37m"
	colorBold    = "\033[1m"
)

func colorize(code, s string) string {
	return code + s + colorReset
}


func red(s string) string { return colorize(colorRed, s) }


func green(s string) string { return colorize(colorGreen, s) }


func yellow(s string) string { return colorize(colorYellow, s) }


func blue(s string) string { return colorize(colorBlue, s) }


func magenta(s string) string { return colorize(colorMagenta, s) }


func cyan(s string) string { return colorize(colorCyan, s) }


func bold(s string) string { return colorize(colorBold, s) }
