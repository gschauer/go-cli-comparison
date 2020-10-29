package chat

import "fmt"

// Version is the version string of this application.
const Version = "0.1.1"

// Connect establishes a connection to the specified server and opens the given channels.
func Connect(svr string, chs []string) error {
	_, err := fmt.Printf("Connecting to %v: %v\n", svr, chs)
	return err
}
