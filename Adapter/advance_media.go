package main

import "fmt"

type advanceMediaPlayer interface {
	playMP4(filename string)
	playVLC(filename string)
}

func newMP4Player() advanceMediaPlayer {
	return &MP4Player{}
}

func newVLCPlayer() advanceMediaPlayer {
	return &VLCPlayer{}
}

type MP4Player struct{}

func (M MP4Player) playMP4(filename string) {
	fmt.Println("playing mp4 filename:", filename)
}

func (M MP4Player) playVLC(filename string) {}

type VLCPlayer struct{}

func (V VLCPlayer) playMP4(filename string) {}

func (V VLCPlayer) playVLC(filename string) {
	fmt.Println("playing vlc filename:", filename)
}
