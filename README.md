# Aura Logger
Aura is a lightweight package to help in Logging.
Being developed by myself.

## Info
![image](https://github.com/astridyz/Aura/assets/163058589/95587326-6c31-4546-bb2f-2885a6d252e8)

## Installation

Run:
`go get "github.com/astridyz/Aura`

Import as:
```go
import aura "github.com/astridyz/Aura/src"
```

Colors package:
```go
import "github.com/astridyz/Aura/src/colors"
```

## Usage
```go
log := aura.NewLogger("Main")

log.SetPrefix(&aura.Prefix{ // --> Prefix aren't needed
	Structure: "Astrid: ",
	Color:     colors.BrightGreen,
})

log.Print("Hey, keep going! Don't give up.")
log.Printf("Hey, %v, how are you?", "Kame")

log.Error("That's an error!")
log.Errorf("Hey, %v, help me with this problem...")

log.Panic("And that's a panic.")
```