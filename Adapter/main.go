package main

func main() {
	player := newAudioPlayer()

	player.play("mp3", "music.mp3")
	player.play("", "music.mp3")
	player.play("mp4", "video.mp4")
	player.play("vlc", "music.vlc")
	player.play("VLC", "music.vlc")
}
