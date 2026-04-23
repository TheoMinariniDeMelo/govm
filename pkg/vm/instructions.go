package vm;
type Inst int;

const (
	MOV Inst = 0;
	ADD;
	IADD;
	SUB;
	ISUB;
	PUSH;
	POP;
	MUL;
	IMUL;
	JMP;
	JNZ;
	JGE;
	JLE;
	JA;
	JAE;
	JB;
	JBE;
	CMP;
	Test;
)
