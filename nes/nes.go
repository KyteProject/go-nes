package nes

import "github.com/kyteproject/NESticle/nes/cpu"

type NES struct {
	CPU cpu.CPU
}

func NewBus() (*NES, error) {
	nes := NES{CPU: nil}

	return &nes, nil
}
