package main

import (
	"fmt"
	"io"
	"net/http"
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

	resp, err := http.Get("https://listen.mixlr.com/aaed1af462ac25481f18b15ed038912a")

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)

	defer resp.Body.Close()

	fmt.Println("Connected")

	fmt.Println("Sending...")

	w, err := shout.Connect(cfg)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	_, err = io.Copy(w, resp.Body)

	if err != nil {
		panic(err)
	}
}
