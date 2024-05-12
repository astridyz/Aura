# Aura Logger
Aura is a lightweight module to help in Logging.
Being developed by myself.

## Info
![image](https://github.com/astridyz/Aura/assets/163058589/95587326-6c31-4546-bb2f-2885a6d252e8)

## Installation

Run:
`go get "github.com/astridyz/Aura"`

Import as:
```go
import aura "github.com/astridyz/Aura/src"
```

## Usage

```go
log := aura.New(&aura.Prefix{Body: "Astrid:", Level: aura.Debug})

log.Print("Hey, keep going! Don't give up.")
log.Printf("Hey, %v, how are you?", "Kame")

log.Warn("Warning! Server is closing!")

log.Error("That's an error!")
log.Errorf("Hey, %v, help me with this problem...", "Kame")

log.Panic("And that's a panic.")
```