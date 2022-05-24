package main

import (
	"fmt"
	"strings"
)

type MediaPlayer interface {
	play(mediaType, filename string)
}

func newAudioPlayer() MediaPlayer {
	return &audioPlayer{}
}

type audioPlayer struct{}

func (a audioPlayer) play(mediaType, filename string) {
	mediaType = strings.TrimSpace(mediaType)
	mediaType = strings.ToLower(mediaType)
	if mediaType == "mp3" {
		// 支持基本的类型播放
		fmt.Println("playing mp3 filename:", filename)
	} else if mediaType == "mp4" || mediaType == "vlc" {
		// 如果有更高级的类型，则使用适配器
		adapter := newMediaPlayerAdapter()
		adapter.play(mediaType, filename)
	} else {
		fmt.Println("unsupported media type")
	}
}
