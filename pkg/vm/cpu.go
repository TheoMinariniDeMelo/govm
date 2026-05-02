package vm;

import (
	"errors"
	"strings"
)

type RegNum int;

const (
	REG_R1 RegNum = 0;
	REG_R2;
	REG_R3;
	REG_R4;
	REG_R5;
	REG_R6;
	REG_R7;
	REG_R8;
	REG_SP;
	REG_BP;
	REG_MB;
	REG_IM;
	REG_IP;
	WORD uint8 = 16;
)


type Cpu16 struct {
	registers []uint16;
	regsNames []string;
	memory *C16MemoryMap;
	interruptVectorAddress uint16;
} 

type C16MemoryMap struct {
	data *uint16;
	limit uint16;
}

// r1-r8, sp, bp, mb, im, ip
func C16CreateCpu(C16MemoryMap *uint16, size int, ) (Cpu16, error){
	var cpu Cpu16;

	cpu.registers = make([]uint16, 13)
	cpu.regsNames[REG_R1] = "r1";
	cpu.regsNames[REG_R2] = "r2";
	cpu.regsNames[REG_R3] = "r3";
	cpu.regsNames[REG_R4] = "r4";
	cpu.regsNames[REG_R5] = "r5";
	cpu.regsNames[REG_R6] = "r6";
	cpu.regsNames[REG_R7] = "r7";
	cpu.regsNames[REG_R8] = "r8";
	cpu.regsNames[REG_SP] = "sp";
	cpu.regsNames[REG_BP] = "bp";
	// memory bank
	cpu.regsNames[REG_MB] = "mb";
	// interrupt mask
	cpu.regsNames[REG_IM] = "im"
	// instruction pointer
	cpu.regsNames[REG_IP] = "ip";

	err := C16SetRegister(cpu ,"sp", 0xffff - 1);
	
	if(err != nil){
		return cpu, err;
	}

	return cpu, nil;
}

func C16SetRegister(cpu Cpu16, reg string, value uint16) error {
	idx, err := C16GetRegisterIndex(cpu, reg);
	if(err != nil){
		return err; 
	}
	cpu.registers[idx] = value;
	return nil;
}

func C16GetRegisterIndex(cpu Cpu16, str string) (uint8, error) {
	for i := range len(cpu.registers){
		if(strings.Compare(cpu.regsNames[i], str) > 0){
			return uint8(i), nil;
		}
	}
	return 0, errors.New("register not found")
}

func C16MemMapLoad(mem *C16MemoryMap, addr uint16, data *uint16, size uint16) error {
	if size > mem.limit {
		return errors.New("data is bigger than memory");
	}

}
