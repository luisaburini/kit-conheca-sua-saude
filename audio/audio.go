package audio

import (
	"bytes"
	"log"
	"strings"
	"time"
	"conheca/sua/saude/resources"
	"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/dialog"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

// func fileExists(path string) error {
// 	_, err := os.Stat(path)
// 	if err == nil {
// 		log.Println("File " + path + " exists.")
// 	} else {
// 		log.Println(err.Error())
// 	}
// 	return err
// }

func Play(sentence string, window fyne.Window) {
	log.Println("sentence " + sentence)
	words := strings.Split(sentence, " ")
	log.Println(words)
	for _, word := range words {
		if word == "" {
			continue
		}
		fileBytes := resources.GetAudioResource(word).StaticContent
		// Convert the pure bytes into a reader object that can be used with the mp3 decoder
		fileBytesReader := bytes.NewReader(fileBytes)
		// Decode file
		decodedMp3, err := mp3.NewDecoder(fileBytesReader)
		if err != nil {
			log.Println("mp3.NewDecoder failed: " + err.Error())
			continue
		} else {
			log.Println("Decoded file successfully")
		}

		// Prepare an Oto context (this will use your default audio device) that will
		// play all our sounds. Its configuration can't be changed later.

		// Usually 44100 or 48000. Other values might cause distortions in Oto
		samplingRate := 44100

		// Number of channels (aka locations) to play sounds from. Either 1 or 2.
		// 1 is mono sound, and 2 is stereo (most speakers are stereo).
		numOfChannels := 2

		// Bytes used by a channel to represent one sample. Either 1 or 2 (usually 2).
		audioBitDepth := 2

		// Remember that you should **not** create more than one context
		otoCtx, readyChan, err := oto.NewContext(samplingRate, numOfChannels, audioBitDepth)
		if err != nil {
			log.Println("oto.NewContext failed: " + err.Error())
			continue
		} else {
			log.Println("oto New Context succeeded")
		}
		// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
		<-readyChan

		// Create a new 'player' that will handle our sound. Paused by default.
		player := otoCtx.NewPlayer(decodedMp3)

		// Play starts playing the sound and returns without waiting for it (Play() is async).
		player.Play()

		// We can wait for the sound to finish playing using something like this
		for player.IsPlaying() {
			time.Sleep(time.Millisecond)
		}
		// If you don't want the player/sound anymore simply close
		err = player.Close()
		if err != nil {
			panic("player.Close failed: " + err.Error())
		} else {
			log.Println("player closed successfully")
		}
	}
}
