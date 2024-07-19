package cfmt

import (
	"testing"
)

func TestHexToRGB(t *testing.T) {
	r, g, b, err := hexToRGB("#FF5733")
	if err != nil {
		t.Fatalf("Expected no error, got %s", err)
	}
	if r != 255 || g != 87 || b != 51 {
		t.Fatalf("Expected RGB (255, 87, 51), got (%d, %d, %d)", r, g, b)
	}
}

func TestRGBToAnsi(t *testing.T) {
	ansiCode := rgbToAnsi(255, 87, 51, false)
	expected := "\033[38;2;255;87;51m"
	if ansiCode != expected {
		t.Fatalf("Expected %s, got %s", expected, ansiCode)
	}
}

func TestPrint(t *testing.T) {
	Print("#FF0000", "Pass if RED")
}

func TestSprint(t *testing.T) {
	result := Sprint("#FF5733", "Hello, World!")
	expected := "\033[38;2;255;87;51mHello, World!\033[0m"
	if result != expected {
		t.Fatalf("Expected %s, got %s", expected, result)
	}
}

func TestPrintf(t *testing.T) {
	// Note: This test will print to the console
	Printf("red", "Pass if, %s!", "RED")
}

func TestSprintf(t *testing.T) {
	result := Sprintf("#FF5733", "Hello, %s!", "World")
	expected := "\033[38;2;255;87;51mHello, World!\033[0m"
	if result != expected {
		t.Fatalf("Expected %s, got %s", expected, result)
	}
}

func TestPrintln(t *testing.T) {
	// Note: This test will print to the console
	Println("red", "Pass if RED and NL")
}

func TestParseFormat(t *testing.T) {
	ansi := NewAnsiFMT()
	result := ansi.parseFormat("->!/bold_color")
	if result != "boldcolor" {
		t.Fatalf("Expected 'bold_color', got '%s'", result)
	}
	if !ansi.Background {
		t.Fatalf("Expected Background to be true, got false")
	}
	if !ansi.Bold {
		t.Fatalf("Expected Bold to be true, got false")
	}
}
func TestConfigFormat(t *testing.T) {
	ansi := NewAnsiFMT()
	ansi.Bold = true
	ansi.Italic = true
	result := ansi.configFormat()
	expected := "\033[1m\033[3m"
	if result != expected {
		t.Fatalf("Expected %s, got %s", expected, result)
	}
}
