package main

import "fmt"

// Implementation
type Device interface {
	SetVolume(int)
}

// Abstraction
type Remote struct {
	device Device
}

func (r *Remote) VolumeUp() {
	r.device.SetVolume(10)
}

type SamsungTV struct {}

func (s *SamsungTV) SetVolume(v int) {
	fmt.Printf("Samsung volume: %d\n", v)
}

