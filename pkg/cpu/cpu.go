package cpu

import (
	"context"
	"errors"
	"strings"
)

type RegNum int;

const (
	REG_R1 RegNum = iota;
	REG_R2;
	REG_R3;
	REG_R4;
	REG_R5;
	REG_R6;
	REG_R7;
	REG_R8;
	REG_SP;
	REG_BP;
	REG_IM;
	REG_IP;
	WORD uint8 = 32;
)



type Cpu32 struct {
	registers [] TObject;

	regsNames []string;

	memory *C32MemoryMap;

	interruptVectorAddress uint32;

	ctx context.Context;
} 

type C32MemoryMap struct {
	data *uint32;
	limit uint32;
}

// r1-r8, sp, bp, mb, im, ip
func C32CreateCpu(C32MemoryMap *uint32, size int, ) (*Cpu32, error){
	var cpu *Cpu32 = &Cpu32{ ctx: context.Background()};


	cpu.registers = make([]TObject, 13);
	cpu.regsNames = make([]string, 13);

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
	// interrupt mask
	cpu.regsNames[REG_IM] = "im"
	// instruction pointer
	cpu.regsNames[REG_IP] = "ip";

	err := C32SetRegister(cpu ,"sp", &Integer32Object{ value: 0xffff - 1 });

	if(err != nil){
		return cpu, err;
	}


	return cpu, nil;
}

func C32SetRegister(cpu *Cpu32, reg string, object TObject) error {
	idx, err := C32GetRegisterIndex(cpu, reg);
	if(err != nil){
		return err; 
	}
	cpu.registers[idx] = object;
	return nil;
}

func C32GetRegisterIndex(cpu *Cpu32, str string) (uint8, error) {
	for i := range len(cpu.registers){
		if(strings.Compare(cpu.regsNames[i], str) > 0){
			return uint8(i), nil;
		}
	}
	return 0, errors.New("register not found")
}

func C32MemMapLoad(mem *C32MemoryMap, addr uint32, data *uint32, size uint32) error {
	if size > mem.limit {
		return errors.New("data is bigger than memory");
	}
	return nil;
}
