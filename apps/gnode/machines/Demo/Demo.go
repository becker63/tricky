package Demo

import (
	"fmt"

	"github.com/becker63/tricky/types"
)

type Demo struct {
	settings types.CommonBox
}

func ConstructDemo(c types.CommonBox) types.Machine {
	return Demo{settings: c}
}

// i is for internal
func (d Demo) Start_i() error {
	return fmt.Errorf("")
}

func (d Demo) Stop_i() error {
	return fmt.Errorf("")
}

func (d Demo) GetConfig_i() error {
	return fmt.Errorf("")
}

func (d Demo) ChangeConfig_i() error {
	return fmt.Errorf("")
}
