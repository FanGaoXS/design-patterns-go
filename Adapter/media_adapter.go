package main

import "strings"

func newMediaPlayerAdapter() MediaPlayer {
	return &mediaPlayerAdapter{}
}

type mediaPlayerAdapter struct{}

func (m mediaPlayerAdapter) play(mediaType, filename string) {
	mediaType = strings.TrimSpace(mediaType)
	mediaType = strings.ToLower(mediaType)
	// 适配器匹配不同类型进行更高级的播放
	if mediaType == "mp4" {
		player := newMP4Player()
		player.playMP4(filename)
	} else if mediaType == "vlc" {
		player := newVLCPlayer()
		player.playVLC(filename)
	}
}
