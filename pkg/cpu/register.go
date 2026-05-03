package cpu

type Procedure func();


type TObject interface {
	Type() string
}

type Integer32Object struct {
	value uint32
}

type StringObject struct {
	value *string
}

type Float32Object struct {
	value float32
}

type ProcedureObject struct {
	value *Procedure 
}

func (o *Integer32Object) Type() string {
	return "Integer32"
}

func (o *StringObject) Type() string {
	return "String"
}

func (o *Float32Object) Type() string {
	return "Float32"
}

func (o *ProcedureObject) Type() string {
	return "Procedure"
}

