package main

import "fmt"

type TV struct{}

func (t *TV) On() { fmt.Println("TV is now ON") }

func (t *TV) Off() { fmt.Println("TV is now OFF") }

type SoundSystem struct{}

func (s *SoundSystem) On()  { fmt.Println("Sound system is ON") }
func (s *SoundSystem) Off() { fmt.Println("Sound system is OFF") }

type StreamingService struct{}

func (ss *StreamingService) Start() { fmt.Println("Streaming service started") }
func (ss *StreamingService) Stop()  { fmt.Println("Streaming service stopped") }

type HomeTheaterFacade struct {
	tv       *TV
	sound    *SoundSystem
	streamer *StreamingService
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		tv:       &TV{},
		sound:    &SoundSystem{},
		streamer: &StreamingService{},
	}
}

func (h *HomeTheaterFacade) WatchMovie() {
	fmt.Println("Preparing to watch a movie...")
	h.tv.On()
	h.sound.On()
	h.streamer.Start()
}

func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting down movie experience...")
	h.streamer.Stop()
	h.sound.Off()
	h.tv.Off()
}
func main() {
	fmt.Println("----- Start Movie -----")
	
	var home HomeTheaterFacade
	home.WatchMovie()

	fmt.Println("\n----- End Movie -----")
	home.EndMovie()
}
