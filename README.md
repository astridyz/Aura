# Aura Logger
Aura is a lightweight package to help in Logging.
Being developed by myself.

## Info
![image](https://github.com/astridyz/Aura/assets/163058589/95587326-6c31-4546-bb2f-2885a6d252e8)

## Usage

```go
func main() {
    	log := NewLogger("Astrid")

	log.SetPrefix(&Prefix{
		Structure: "Astrid: ",
		Color:     colors.BrightPink,
	})

	log.Print("Hey, keep going! Don't give up.")
	log.Printf("Hey, %v, how are you?", "Kame")

	log.Error("That's an error!")
	log.Errorf("Hey, %v, help me with this problem...", "Kame")

	log.Panic("And that's a panic.")
}
```
