package Test

import (
	"fmt"

	"github.com/becker63/tricky/types"
)

type Test struct {
	settings      types.CommonBox
	extrasettings string
}

func ConstructTest(c string) types.Machine {
	return Test{extrasettings: c}
}

func (d Test) Start_i() error {
	return fmt.Errorf("")
}

func (d Test) Stop_i() error {
	return fmt.Errorf("")
}

func (d Test) GetConfig_i() error {
	return fmt.Errorf("")
}

func (d Test) ChangeConfig_i() error {
	return fmt.Errorf("")
}
