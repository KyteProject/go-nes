/*

An emulation of the 6502 microprocessor used in the NES.

"The 6502 microprocessor is a relatively simple 8 bit CPU with only a few internal registers capable of addressing at most 64Kb of memory via its 16 bit address bus. The processor is little endian and expects addresses to be stored in memory least significant byte first.

The first 256 byte page of memory ($0000-$00FF) is referred to as 'Zero Page' and is the focus of a number of special addressing modes that result in shorter (and quicker) instructions or allow indirect access to the memory. The second page of memory ($0100-$01FF) is reserved for the system stack and which cannot be relocated.

The only other reserved locations in the memory map are the very last 6 bytes of memory $FFFA to $FFFF which must be programmed with the addresses of the non-maskable interrupt handler ($FFFA/B), the power on reset location ($FFFC/D) and the BRK/interrupt request handler ($FFFE/F) respectively.

The 6502 does not have any special support of hardware devices so they must be mapped to regions of memory in order to exchange data with the hardware latches."

- Copied from: http://www.obelisk.me.uk/6502/architecture.html

*/

package cpu

type Status uint8

// Enum Processor Status (P) flags
const (
	C Status = 1 << iota // Carry Bit
	Z                    // Zero
	I                    // Disable Interrupts
	D                    // Decimal Mode (unused in this implementation)
	B                    // Break
	U                    // Unused
	V                    // Overflow
	N                    // Negative
)

/*
	All 6502 registers are 8-bit, with the exception of the program counter which is 16-bit.

	Accumulator: 		Used for all arithmetic and logic instructions (with the exception of
						increments and decrements).
	Index X: 			Used to hold counters or offsets for accessing memory. Can be used to get
						a copy of the stack pointer or change its value.
	Index Y: 			Same as X, but without any special functions.
	Processor Status: 	As instructions are executed a set of processor flags are set or clear to
						record the results of the operation.
	Program Counter: 	Points to the next instruction to be executed. The value of program
						counter is modified automatically as instructions are executed.
	Stack Pointer: 		The processor supports a 256 byte stack located between $0100 and $01FF.
*/
type Registers struct {
	A  uint8  // Accumulator
	X  uint8  // Index X
	Y  uint8  // Index Y
	P  Status // Processor Status
	SP uint8  // Stack Pointer
	PC uint16 // Program Counter
}

//type AddressMode struct {
//}
//
//func (a *AddressMode) Implicit() {}
//
//func (a *AddressMode) Immediate() {}
//
//func (a *AddressMode) ZeroPage() {}
//
//func (a *AddressMode) ZeroPageX() {}
//
//func (a *AddressMode) ZeroPageY() {}
//
//func (a *AddressMode) Relative() {}
//
//func (a *AddressMode) Absolute() {}
//
//func (a *AddressMode) AbsoluteX() {}
//
//func (a *AddressMode) AbsoluteY() {}
//
//func (a *AddressMode) Indirect() {}
//
//func (a *AddressMode) IndexedIndirect() {}
//
//func (a *AddressMode) IndirectIndexed() {}

type CPU struct {
	RAM Memory
	Registers
}

func NewCPU() *CPU {
	cpu := &CPU{RAM: NewProcessorMemory()}
	return cpu
}

func (c *CPU) Clock() {
	return
}
func (c *CPU) Reset() {
	return
}
func (c *CPU) IRQ() {
	return
}
func (c *CPU) NMI() {
	return
}
