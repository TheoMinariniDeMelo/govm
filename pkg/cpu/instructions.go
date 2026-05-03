package cpu;
type Inst uint8;


// k = constant pool
const (
	// memory management
	MOV_K Inst = 0x10; // r[A] = k[B]
	MOV_REG Inst = 0x13; // r[A] = r[B]

	LOAD Inst = 0x14
	LOAD_MEM_FUNC Inst = 0x14;
	LOAD_INT Inst = 0x16;
	LOAD_FLOAT Inst = 0x17;
	LOAD_STRING Inst = 0x18;

	// math
	ADD_REG_REG Inst = 0x40; // r[A] += r[B]
	ADD_REG_K Inst = 0x41; // r[A] += k[B]
	ADD_REG_MEMORY Inst = 0x42; // r[A] += memory[B]
	ADD_REG_REG_REG Inst = 0x43; // r[A] = r[B] + r[C]
	ADD_REG_REG_K Inst = 0x43; // r[A] = r[B] + k[C]
	ADD_REG_REG_MEMORY Inst = 0x43; // r[A] = r[B] + memory[C]

	SUB_REG_REG Inst = 0x50; // r[A] -= r[B]
	SUB_REG_K Inst = 0x51; // r[A] -= k[B]
	SUB_REG_MEMORY Inst = 0x52; // r[A] -= memory[B]

	// stack
	PUSH Inst = 0x60;
	POP Inst = 0x61;

	// interruption
	INT Inst = 0x70;
	RET_INT Inst = 0x71;

	CMP_REG_REG Inst = 0x80;
	CMP_REG_K Inst = 0x81;
	CMP_REG_MEMORY Inst = 0x82;

	// bitmask
	AND_REG_REG Inst = 0x90;
	AND_REG_K Inst = 0x91;
	AND_REG_MEMORY Inst = 0x92;
	
	SHL_REG_REG Inst = 0x93;
	SHL_REG_K Inst = 0x94;
	SHL_REG_MEMORY Inst = 0x95;
	SHR_REG_REG Inst = 0x96;
	SHR_REG_K Inst = 0x97;
	SHR_REG_MEMORY Inst = 0x98;

	SAL_REG_REG Inst = 0x99;
	SAL_REG_K Inst = 0x9A;
	SAL_REG_MEMORY Inst = 0x9B;
	SAR_REG_REG Inst = 0x9C;
	SAR_REG_K Inst = 0x9D;
	SAR_REG_MEMORY Inst = 0x9E;

	// non linear exection
	JMP Inst = 0xA0;
	JZ Inst = 0xA1;
	JE Inst = 0xA2;
	JEG Inst = 0xA3;
	JNE Inst = 0xA4;
	JBE Inst = 0xA5;
	JNA Inst = 0xA6;
	JMP_TRUE Inst = 0xA7; // used in 'IS' functions

	// validation
	IS_INTEGER Inst = 0xC0; // if type(r[A]) == IntegerObject then flag 
	IS_STRING Inst = 0xC1;
	IS_FLOAT Inst = 0xC2;
	IS_FUNCTION Inst = 0xC3;

	// call functions

)
