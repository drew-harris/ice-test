package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	shout "github.com/jdarnold/go-libshout"
)

// Setup some command line flags

func main() {
	fmt.Println("WE MADE IT DUDE!!")
	flag.Parse()

	// Setup libshout parameters
	s := shout.Shout{
		Host:     "localhost",
		Port:     8000,
		User:     "source",
		Password: "hackme",
		Mount:    "test-mount",
		Format:   shout.FORMAT_MP3,
		Protocol: shout.PROTOCOL_HTTP,
	}

	// Open the file
	//
	file, err := os.Open("ageis.mp3")
	if err != nil {
		panic(err)
	}

	// Create a channel where we can send the data
	//
	stream, err := s.Open()
	if err != nil {
		panic(err)
	}
	// /opt/homebrew/Cellar/libshout/2.4.6_1/include/shout/shout.h

	defer s.Close()

	buffer := make([]byte, shout.BUFFER_SIZE)
	fmt.Println("BUFFER SIZE: ", shout.BUFFER_SIZE)
	for {
		// Read from file
		n, err := file.Read(buffer)

		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0 {
			break
		}

		// Send to shoutcast server
		stream <- buffer
	}
}
