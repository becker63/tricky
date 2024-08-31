package Container

import (
	"fmt"

	"github.com/becker63/tricky/types"
)

type Container struct {
	settings      types.CommonBox
	extrasettings string
}

func ConstructContainer(c string) types.Machine {
	return Container{extrasettings: c}
}

func (d Container) Start_i() error {
	return fmt.Errorf("")
}

func (d Container) Stop_i() error {
	return fmt.Errorf("")
}

func (d Container) GetConfig_i() error {
	return fmt.Errorf("")
}

func (d Container) ChangeConfig_i() error {
	return fmt.Errorf("")
}
