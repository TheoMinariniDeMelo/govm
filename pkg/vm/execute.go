package vm;

var (
	regs []int16;
	stack []int8 = make([]int8, 4*1024);
	memory []int8 = make([]int8, 64*1024);
)
