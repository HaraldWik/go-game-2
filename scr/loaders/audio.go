package load

import (
	"log"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type Audio struct {
	Audio   *mix.Music
	Volume  uint32
	Channel uint32
}

func NewAudio(path string) *Audio {
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		log.Fatalf("Failed to init SDL audio: %v", err)
		sdl.Quit()
	}

	if err := mix.Init(mix.INIT_MP3); err != nil {
		log.Fatalf("Failed to init SDL mixer: %v", err)
		mix.Quit()
	}

	if err := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Fatalf("Failed to open audio: %v", err)
		mix.CloseAudio()
	}

	if audio, err := mix.LoadMUS(path); err != nil {
		log.Fatalf("Failed to load audio: %v", err)
		audio.Free()
	} else {
		return &Audio{
			Audio:  audio,
			Volume: mix.MAX_VOLUME,
		}
	}
	return nil
}

func (a *Audio) Play(loops int32) {
	if err := a.Audio.Play(int(loops)); err != nil {
		log.Fatalf("Failed to play music: %v", err)
	}
	a.SetVolume(a.Volume)
}

func (a *Audio) Stop() {
	if a.Audio != nil {
		if mix.PlayingMusic() {
			mix.HaltMusic()
		}
		a.Audio.Free()
		a.Audio = nil
	}
}

func (a *Audio) SetVolume(volume uint32) {
	if volume > mix.MAX_VOLUME {
		volume = mix.MAX_VOLUME
	}
	a.Volume = uint32(volume)
	mix.VolumeMusic(int(volume))
}

func (a *Audio) GetVolume() uint32 {
	return a.Volume
}

func (a *Audio) IsPlaying() bool {
	return mix.PlayingMusic()
}
