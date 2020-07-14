package nes

type NES struct {
	CPU uint16 // *OLC6502
	RAM	[]uint8
}

func NewBus() (*NES, error) {
	ram := make([]uint8, 64 * 1024)
	nes := NES{CPU: nil, RAM: ram}

	return &nes, nil
}