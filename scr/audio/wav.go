package audio

/*
#include <SDL2/SDL.h>
typedef unsigned char Uint8;
void OnAudioPlayback(void *userdata, Uint8 *stream, int len);
*/
import "C"
import (
	"log"
	"runtime"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

type Wav struct {
	Data   []byte
	Offset int
	dev    sdl.AudioDeviceID
}

//export OnAudioPlayback
func OnAudioPlayback(userdata unsafe.Pointer, stream *C.Uint8, length C.int) {
	audio := (*Wav)(userdata)
	n := int(length)
	// Convert C buffer to Go slice
	buf := ((*[1 << 30]byte)(unsafe.Pointer(stream)))[:n:n]
	for i := 0; i < n; i++ {
		buf[i] = audio.Data[audio.Offset]
		audio.Offset = (audio.Offset + 1) % len(audio.Data)
	}
	// Ensure that 'audio' is not garbage collected
	runtime.KeepAlive(audio)
}

func NewWavAudio(filePath string) *Wav {
	var err error
	audio := &Wav{}

	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		log.Fatalf("Failed to initialize SDL audio: %v", err)
	}

	tmpaudio, spec := sdl.LoadWAV(filePath)
	if spec == nil {
		log.Fatalf("Failed to load WAV file '%s': %v", filePath, sdl.GetError())
		sdl.Quit()
	}
	audio.Data = tmpaudio

	spec.Callback = sdl.AudioCallback(C.OnAudioPlayback)
	spec.UserData = unsafe.Pointer(audio)

	if audio.dev, err = sdl.OpenAudioDevice("", false, spec, nil, 0); err != nil {
		sdl.Quit()
		log.Fatalf("Failed to open audio device: %v", err)
	}

	return audio
}

func (a *Wav) Play() {
	sdl.PauseAudioDevice(a.dev, false)
}

func (a *Wav) Stop() {
	sdl.CloseAudioDevice(a.dev)
	sdl.Quit()
}
