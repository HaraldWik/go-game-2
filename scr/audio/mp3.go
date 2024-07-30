package audio

import (
	"log"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type MP3 struct {
	Audio  *mix.Music
	Volume uint32
}

func LoadMP3(path string) *MP3 {
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		log.Fatalf("Failed to initialize SDL: %v", err)
		sdl.Quit()
	}

	if err := mix.Init(mix.INIT_MP3); err != nil {
		log.Fatalf("Failed to initialize SDL_mixer: %v", err)
		mix.Quit()
	}

	if err := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Fatalf("Failed to open audio: %v", err)
		mix.CloseAudio()
	}

	audio, err := mix.LoadMUS(path)
	if err != nil {
		log.Fatalf("Failed to load music: %v", err)
		audio.Free()
	}

	return &MP3{
		Audio:  audio,
		Volume: mix.MAX_VOLUME,
	}
}

func (m *MP3) Play(loops int32) {
	if err := m.Audio.Play(int(loops)); err != nil {
		log.Fatalf("Failed to play music: %v", err)
	}
	m.SetVolume(m.Volume)
}

func (m *MP3) Stop() {
	if m.Audio != nil {
		if mix.PlayingMusic() {
			mix.HaltMusic()
		}
		m.Audio.Free()
		m.Audio = nil
	}
}

func (m *MP3) SetVolume(volume uint32) {
	if volume > mix.MAX_VOLUME {
		volume = mix.MAX_VOLUME
	}
	m.Volume = uint32(volume)
	mix.VolumeMusic(int(volume))
}

func (m *MP3) GetVolume() uint32 {
	return m.Volume
}
