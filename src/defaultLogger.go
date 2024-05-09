package aura

import (
	"fmt"
	"os"

	"github.com/astridyz/Aura/src/colors"
)

var defaultLogger *logger = NewLogger()

// --> Print

// Is equivalent to print() with color and prefix if defined.
func Print(message ...any) string {
	result := defaultLogger.log(nil, colors.Cyan, message...)

	print(result)
	return result
}

// Is equivalent to logger.Print() with formatting
func Printf(format string, arguments ...any) string {
	text := fmt.Sprintf(format, arguments...)
	return defaultLogger.Print(text)
}

// --> Warn

// Is equivalent to print() with color and prefix if defined.
func Warn(message ...any) string {
	subPrefix := &Prefix{
		Structure: "WARN ",
		Color:     colors.BrightYellow,
	}
	result := defaultLogger.log(subPrefix, colors.BrightYellow, message...)

	print(result)
	return result
}

// Is equivalent to logger.Print() with formatting
func Warnf(format string, arguments ...any) string {
	text := fmt.Sprintf(format, arguments...)
	return defaultLogger.Warn(text)
}

// --> Panic

// It prints a message and call panic() with color
func Panic(message ...any) {
	subPrefix := &Prefix{
		Structure: "PANIC",
		Color:     colors.BrightPurple,
	}
	result := defaultLogger.log(subPrefix, colors.BrightYellow, message...)

	print(result)
	panic(fmt.Sprint(message...))
}

// Equivalent to logger.Panic() with formatting
func Panicf(format string, arguments ...any) {
	text := fmt.Sprintf(format, arguments...)
	defaultLogger.Panic(text)
}

// --> Error

// Prints an error message with color and prefix
func Error(message ...any) {
	subPrefix := &Prefix{
		Structure: "ERROR",
		Color:     colors.BrightRed,
	}
	result := defaultLogger.log(subPrefix, colors.BrightYellow, message...)

	print(result)
}

// Equivalent to logger.Error() with formatting
func Errorf(format string, arguments ...any) {
	text := fmt.Sprintf(format, arguments...)
	defaultLogger.Error(text)
}

// --> Fatal

// Prints an error message with color and prefix and calls os.Exit(1)
func Fatal(message ...any) {
	subPrefix := &Prefix{
		Structure: "FATAL",
		Color:     colors.BrightOrange,
	}
	result := defaultLogger.log(subPrefix, colors.BrightYellow, message...)

	print(result)
	os.Exit(1)
}

// Equivalent to logger.Fatal() with formatting
func Fatalf(format string, arguments ...any) {
	text := fmt.Sprintf(format, arguments...)
	defaultLogger.Fatal(text)
}
