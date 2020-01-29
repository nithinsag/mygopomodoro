package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func playFileAndBlock(file string) {
	f, err := os.Open(file + ".mp3")
	if err != nil {
		log.Fatal(err)
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	// Callback creates a streamer which calls our callback
	// Seq stiches the streamers together
	// And Play plays the
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func main() {
	var countdown int64
	countdown = 30
	if len(os.Args) > 1 {
		var err error
		countdown, err = strconv.ParseInt(os.Args[1], 10, 64)
		if err != nil {
			log.Fatal("optional argument is the number of minutes for your pomodoro: eg. 20")
		}

	}
	fmt.Printf("Starting countdown with %d minute\n", countdown)
	// donetimer := make(chan bool)
	time.Sleep(time.Duration(countdown) * time.Minute)
	playFileAndBlock("ketchup")
}
