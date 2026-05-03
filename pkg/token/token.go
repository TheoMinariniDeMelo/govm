package token

const (
	// memory management
	MOV string = "MOV" 

	// math
	ADD string = "ADD"
	SUB string = "SUB"

	// stack
	PUSH string = "PUSH"
	POP string = "POP"

	// interruption
	INT string = "INT"
	RET_INT string = "RET_INT"

	// non linear exection
	JMP string = "JMP"
	JZ string = "JZ"
	JE string = "JE"
	JEG string = "JEG"
	JNE string = "JNE"
	JBE string = "JBE"
	JNA string = "JNA"
)
