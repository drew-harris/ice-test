package main

import (
	"fmt"
	"io"
	"os"

	"github.com/viert/shout"
)

func main() {
	f, err := os.Open("./ageis.mp3")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	cfg := &shout.Config{
		Host:     "localhost",
		Port:     8000,
		User:     "source",
		Password: "hackme",
		Mount:    "/shuffle",
		Proto:    shout.ProtocolHTTP,
		Format:   shout.ShoutFormatMP3,
	}

	w, err := shout.Connect(cfg)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	io.Copy(w, f)

	fmt.Println("Done")
}
