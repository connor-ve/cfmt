package cfmt

import (
	"cfmt/models"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

func checkInput(hexString string) (int, int, int, error) {

	if hexValue, ok := models.Colors[hexString]; ok {
		return hexToRGB(hexValue)
	} else {
		return hexToRGB(hexString)
	}
}

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

func rgbToAnsi(r, g, b int) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func brightAnsi(a bool) string {
	if a {
		return "\033[1m" 
	} else {
		return ""
	}
}

func resetAnsi() string {
	return "\033[0m"
}

// Implementation of print from fmt using color inputs
func Print(hex string, a ...any) {
	hex, bold := checkBold(hex)
	r, g, b, err := checkInput(hex)
	if err != nil {
		log.Fatalf("Failed to convert hex to RGB: %s\n", err)
	}

	ansiCode := rgbToAnsi(r, g, b)
	var a_string string
	if len(a) == 1 {
		a_string = fmt.Sprint(a[0])
	} else {
		a_string = fmt.Sprint(a...)
	}
	fmt.Print(brightAnsi(bold) + ansiCode + a_string + resetAnsi())
}

func Sprint(hex string, a ...any) string {
	hex, bold := checkBold(hex)
	r, g, b, err := checkInput(hex)
	if err != nil {
		log.Fatalf("Failed to convert hex to RGB: %s\n", err)
	}

	ansiCode := rgbToAnsi(r, g, b)
	var a_string string
	if len(a) == 1 {
		a_string = fmt.Sprint(a[0])
	} else {
		a_string = fmt.Sprint(a...)
	}
	return fmt.Sprint(brightAnsi(bold) + ansiCode + a_string + resetAnsi())
}

func Printf(hex string, format string, a ...any) {
	// hex, bold := checkBold(hex)
	// r, g, b, err := hexToRGB(hex)
	// if err != nil {
	// 	log.Fatalf("Failed to convert hex to RGB: %s\n", err)
	// }
	// ansiCode := rgbToAnsi(r, g, b)
	// a_string := fmt.Sprint(a)
	// fmt.Println(ansiCode + a_string + resetAnsi())
}

func Println(hex string, a ...any) {
	// hex, bold := checkBold(hex)
	r, g, b, err := checkInput(hex)
	if err != nil {
		log.Fatalf("Failed to convert hex to RGB: %s\n", err)
	}
	ansiCode := rgbToAnsi(r, g, b)
	var a_string string
	if len(a) == 1 {
		a_string = fmt.Sprint(a[0])
	} else {
		a_string = fmt.Sprint(a...)
	}
	fmt.Println(ansiCode + a_string + resetAnsi())
}

func Printrc(hex string, a ...any) {
	// hex, bold := checkBold(hex)
	r, g, b, err := checkInput(hex)
	if err != nil {
		log.Fatalf("Failed to convert hex to RGB: %s\n", err)
	}
	ansiCode := rgbToAnsi(r, g, b)
	a_string := fmt.Sprint(a...)
	fmt.Print("\r" + ansiCode + a_string + resetAnsi())
}

func checkBold(hex string) (string, bool) {
	bold := false
	if boldChck := strings.Index(hex, "!"); boldChck >= 0 {
		hex = hex[:len(hex)-1]
		bold = true
	}
	return hex, bold
}
