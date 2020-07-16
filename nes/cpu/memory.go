package cpu

type Memory interface {
	Read(address uint16) (value uint8)
	Write(address uint16, value uint8)
}

type ProcessorMemory struct {
	RAM []byte
}

func NewProcessorMemory() Memory {
	return &ProcessorMemory{
		RAM: make([]byte, 2048), // 2KB memory
	}
}

func (p *ProcessorMemory) Read(address uint16) uint8 {
	return p.RAM[address]
}

func (p *ProcessorMemory) Write(address uint16, value uint8) {
	p.RAM[address] = value
}
