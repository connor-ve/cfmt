# cfmt

_add small badges_

---

`cfmt` aka colorfmt is your way to easily add color to your Strings and Outputs using a similar library to your beloved `fmt`

Under the hood, we take care of the annoy syntax to allow for ANSI Escape Codes formatting within your strings and outputs. For the user all functions and usecases are the same but we have added an extra parameter to your traditional fmt.

## Install

```bash
go get github.com/connor-ve/cfmt
```

## Formatting String Syntax

To allow for easy implementation the first parameter of your fmt functions is a string to help represent the color and additional formats you would like to pass to your string.

In our snippets we will be using `cfmt.Print()`, but this will hold true throughout all.

#### Color Format

Within `cfmt` we currently accept Hex Codes and semantic names for colors.

```go
// Hex
cfmt.Print("ff0000", "Hello World") // Output will be in red
cfmt.Print("00ff00", "Hello World") // Output will be in green

// Semantic
cfmt.Print("red", "Hello World")    // Output will be in red
cfmt.Print("green", "Hello World")  // Output will be in green
```

> **NOTE** : Not all terminals will support true colors, in this case your hex codes will be set in factions of 51. I am unsure why this is the number.

#### Text Format

Using ANSI we are also able to achieve some fun text decorations. Current support will be shown in examples below.

These modifiers should either be added to the front or back of string respectively but logic is written to work wherever they are added. It is recommended that no spaces are added, but again we have handled this use case.

```go
// Output will be BOLD
cfmt.Print("!", "Hello World")

// Output will be Italic
cfmt.Print("/", "Hello World")

// Output will be Underlined
cfmt.Print("_", "Hello World")

// Output will have Strikethrough
cfmt.Print("!", "Hello World")

// All can be used at once, or any combination
cfmt.Print("!_-/", "Hello World")

// All can and should be used with color (both semantic or hex)
cfmt.Print("!red", "Hello World")
cfmt.Print("!FF0000", "Hello World")

// When used with color one can also set the background color using "->"
cfmt.Print("->red", "Hello World")

// Technically this is valid but please dont
cfmt.Print("->bl !u/e", "Hello World")
```

## Current `fmt` Functions Implemented

#### Print

```go
cfmt.Print("!red", "Hello World")

// compared to

fmt.Print("Hello World")
```

#### Printf

```go
cfmt.Printf("!red", "Hello %s", "World")

// compared to

fmt.Printf("Hello %s", "World")
```

#### Println

```go
fmt.Println("Hello World")

// compared to

fmt.Println("Hello World")
```

#### Sprint

```go
str := fmt.Sprint("orange", "Hello World")

// compared to

str := fmt.Sprint("Hello World")
```

#### Sprintf

```go
str := cfmt.Sprintf("!red", "Hello %s", "World")

// compared to

str := fmt.Sprintf("Hello %s", "World")
```

## Contributing

`To Be Made`

## FAQ

`To Be Made`
