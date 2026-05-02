package vm

import (
	"io"
	"os"
	"strings"
)

func Load(cpu Cpu16, addr uint16, code []byte){
	C16MemMapLoad(cpu.memory, addr, code, len(code));
}

func Fetch(cpu *Cpu16) ([]byte, error) {
	region := cpu.registers[REG_IP];
	instruction := cpu.memory.data[region:region + 3];

}
