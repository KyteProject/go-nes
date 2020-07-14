package nes

type Memory interface {
	Read(address uint16, readOnly bool) uint8
	Write(address uint16, data uint8)
}

type Mem struct {
	nes *NES
}

func NewMem(nes *NES) Memory {
	return &Mem{nes:nes}
}

func (m *Mem) Read (addr uint16, readOnly bool) uint8 {
	if addr >= 0x0000 && addr <= 0xFFFF {
		return m.nes.RAM[addr]
	}
	return 0x00
}

func (m *Mem) Write (addr uint16, data uint8) {
	if addr >= 0x0000 && addr <= 0xFFFF {
		m.nes.RAM[addr] = data
	}
}