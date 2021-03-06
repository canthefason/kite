package main

import (
	"fmt"

	"github.com/koding/kite"
)

func main() {
	// Create a kite
	k := kite.New("math", "1.0.0")

	// Add our handler method,authentication is disabled for this example
	k.HandleFunc("square", Square).DisableAuthentication()

	// Attach to a server and run it
	k.Config.Port = 3636
	k.Run()
}

func Square(r *kite.Request) (interface{}, error) {
	// Unmarshal method arguments
	a := r.Args.One().MustFloat64()

	result := a * a

	fmt.Printf("Call received, sending result %.0f back\n", result)

	// Print a log on remote Kite.
	// This message will be printed on client's console.
	r.Client.Go("kite.log", fmt.Sprintf("Message from %s: \"You have requested square of %.0f\"", r.LocalKite.Kite().Name, a))

	// You can return anything as result, as long as it is JSON marshalable.
	return result, nil
}
