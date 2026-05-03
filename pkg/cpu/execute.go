package cpu

import (
	"context"
	"fmt"
)

func Run(cpu *Cpu16) error {
	for{
		select {
			case <- cpu.ctx.Done():
				return fmt.Errorf("timeout during execution");
			default:
				// pass
		}

		instruction, err  := Fetch(cpu)

		if err != nil {
			return err;
		}
	
		switch instruction[0] >> 2 {

		}


	}


}
