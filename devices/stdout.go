package devices

import (
	"fmt"

	"github.com/vkstack/expense-tracker/interfaces/iodevice"
)

type stdoutdev struct{}

func (stdo *stdoutdev) Write(message interface{}) {
	fmt.Println(message)
}

func NewStdOutDevice() iodevice.IOut {
	return &stdoutdev{}
}
