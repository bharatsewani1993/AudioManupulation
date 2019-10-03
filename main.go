package main

import (
	"fmt"
	"os"
	"github.com/hajimehoshi/go-mp3"
)

func main() {
	fmt.Println("Execution Started")
	file, err := os.Open("./abc.mp3")
	filed, err := mp3.NewDecoder(file)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//Get Filesize and SampleRate
	filesize := filed.Length()
	samplerate := filed.SampleRate()
	
	//Print FileSize and SampleRate
	fmt.Println("FileSize is ==>", filesize)
	fmt.Println("SampleRate is ==>", samplerate)
}
