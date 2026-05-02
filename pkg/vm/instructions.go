package vm;
type Inst uint8;

const (
	MOV Inst = 0x10;
	ADD Inst = 0x11;
	SUB Inst = 0x12;
	PUSH Inst = 0x20;
	POP Inst = 0x21;
	INT Inst = 0x30;
	RET_INT Inst = 0x30;
	JMP Inst = 0xA0;
	JZ Inst = 0xA1;
	JE Inst = 0xA2;
	JEG Inst = 0xA3;
	JNE Inst = 0xA4;
	JBE Inst = 0xA5;
	JNA Inst = 0xA6;
)
