package cfmt

import (
	"cfmt/models"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

type AnsiFMT struct {
	Background    bool
	Bold          bool
	Italic        bool
	Underline     bool
	Strikethrough bool
}

// Default is set to foreground
func NewAnsiFMT() AnsiFMT {
	ansiFMT := AnsiFMT{}
	ansiFMT.Background = false
	ansiFMT.Bold = false
	ansiFMT.Italic = false
	ansiFMT.Underline = false
	ansiFMT.Strikethrough = false
	return ansiFMT
}

/*
Validates Hexidecimal versus Text input
*/
func inputValidate(hexString string) (int, int, int, error) {

	if hexValue, ok := models.Colors[hexString]; ok {
		return hexToRGB(hexValue)
	} else {
		return hexToRGB(hexString)
	}
}

/*
Ansi Escape Codes work for color with RGB not Hex
This converts to a value we can color the string with
*/
func hexToRGB(hexString string) (int, int, int, error) {
	if poundChck := strings.Index(hexString, "#"); poundChck >= 0 {
		hexString = hexString[1:]
	}

	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return 0, 0, 0, err
	}
	// fmt.Println(bytes)
	if len(bytes) != 3 {
		return 0, 0, 0, fmt.Errorf("invalid hex color length")
	}
	return int(bytes[0]), int(bytes[1]), int(bytes[2]), nil
}

/*
Output of hexToRGB is now a string that can dictacte color in ANSI
*/
func rgbToAnsi(r int, g int, b int, back bool) string {
	if back {
		return fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func resetAnsi() string {
	return "\033[0m"
}

// Implementation of print from fmt using color inputs
func Print(hex string, a ...any) {
	ansi := NewAnsiFMT()
	color := ansi.parseFormat(hex)
	ansiCode := ""
	if len(color) != 0 {
		r, g, b, err := inputValidate(strings.ToLower(color))
		if err != nil {
			log.Fatalf("Failed to convert hex to RGB: %s\n", err)
		}
		ansiCode = rgbToAnsi(r, g, b, ansi.Background)
	}

	var a_string string
	if len(a) == 1 {
		a_string = fmt.Sprint(a[0])
	} else {
		a_string = fmt.Sprint(a...)
	}

	fmt.Print(ansi.configFormat() + ansiCode + a_string + resetAnsi())
}

func Sprint(hex string, a ...any) string {
	ansi := NewAnsiFMT()
	color := ansi.parseFormat(hex)
	r, g, b, err := inputValidate(strings.ToLower(color))
	if err != nil {
		log.Fatalf("Failed to convert hex to RGB: %s\n", err)
	}
	ansiCode := rgbToAnsi(r, g, b, ansi.Background)
	var a_string string
	if len(a) == 1 {
		a_string = fmt.Sprint(a[0])
	} else {
		a_string = fmt.Sprint(a...)
	}
	return fmt.Sprint(ansi.configFormat() + ansiCode + a_string + resetAnsi())
}

func Printf(hex string, format string, a ...any) {
	ansi := NewAnsiFMT()
	color := ansi.parseFormat(hex)
	r, g, b, err := inputValidate(strings.ToLower(color))
	if err != nil {
		log.Fatalf("Failed to convert hex to RGB: %s\n", err)
	}
	ansiCode := rgbToAnsi(r, g, b, ansi.Background)
	a_string := fmt.Sprintf(format, a...)
	fmt.Println(ansi.configFormat() + ansiCode + a_string + resetAnsi())
}

func Sprintf(hex string, format string, a ...any) string {
	ansi := NewAnsiFMT()
	color := ansi.parseFormat(hex)
	r, g, b, err := inputValidate(strings.ToLower(color))
	if err != nil {
		log.Fatalf("Failed to convert hex to RGB: %s\n", err)
	}
	ansiCode := rgbToAnsi(r, g, b, ansi.Background)
	a_string := fmt.Sprintf(format, a...)
	return fmt.Sprint(ansi.configFormat() + ansiCode + a_string + resetAnsi())
}

func Println(hex string, a ...any) {
	ansi := NewAnsiFMT()
	color := ansi.parseFormat(hex)
	r, g, b, err := inputValidate(strings.ToLower(color))
	if err != nil {
		log.Fatalf("Failed to convert hex to RGB: %s\n", err)
	}
	ansiCode := rgbToAnsi(r, g, b, ansi.Background)
	var a_string string
	if len(a) == 1 {
		a_string = fmt.Sprint(a[0])
	} else {
		a_string = fmt.Sprint(a...)
	}

	fmt.Println(ansi.configFormat() + ansiCode + a_string + resetAnsi())
}

func (this *AnsiFMT) parseFormat(inputString string) string {
	inputString = strings.Replace(inputString, " ", "", -1)
	if strings.Contains(inputString, "!") {
		this.Bold = true
		inputString = strings.Replace(inputString, "!", "", -1)
	}
	if strings.Contains(inputString, "->") {
		this.Background = true
		inputString = strings.Replace(inputString, "->", "", -1)
	}

	if strings.Contains(inputString, "/") {
		this.Italic = true
		inputString = strings.Replace(inputString, "/", "", -1)
	}

	if strings.Contains(inputString, "-") {
		this.Strikethrough = true
		inputString = strings.Replace(inputString, "-", "", -1)
	}

	if strings.Contains(inputString, "_") {
		this.Underline = true
		inputString = strings.Replace(inputString, "_", "", -1)
	}

	return inputString
}

func (this *AnsiFMT) configFormat() string {
	format := ""
	if this.Bold {
		format += "\033[1m"
	}
	if this.Italic {
		format += "\033[3m"
	}
	if this.Underline {
		format += "\033[4m"
	}
	if this.Strikethrough {
		format += "\033[9m"
	}
	return format
}
