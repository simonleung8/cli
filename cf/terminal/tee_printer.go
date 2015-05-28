package terminal

import (
	"fmt"
)

type Printer interface {
	Print(a ...interface{}) (n int, err error)
	Printf(format string, a ...interface{}) (n int, err error)
	Println(a ...interface{}) (n int, err error)
	ForcePrint(a ...interface{}) (n int, err error)
	ForcePrintf(format string, a ...interface{}) (n int, err error)
	ForcePrintln(a ...interface{}) (n int, err error)
	DisableTerminalOutput(bool)
}

type OutputCapture interface {
	GetOutputAndReset() []string
}

type TerminalOutputSwitch interface {
	DisableTerminalOutput(bool)
}

type TeePrinter struct {
	disableTerminalOutput bool
	output                []string
}

func NewTeePrinter() *TeePrinter {
	return &TeePrinter{
		output: []string{},
	}
}

func (t *TeePrinter) GetOutputAndReset() []string {
	currentOutput := t.output
	t.output = []string{}
	return currentOutput
}

func (t *TeePrinter) Print(values ...interface{}) (n int, err error) {
	str := fmt.Sprint(values...)
	t.output = append(t.output, Decolorize(str))
	if !t.disableTerminalOutput {
		return fmt.Print(str)
	}
	return
}

func (t *TeePrinter) Printf(format string, a ...interface{}) (n int, err error) {
	str := fmt.Sprintf(format, a...)
	t.output = append(t.output, Decolorize(str))
	if !t.disableTerminalOutput {
		return fmt.Print(str)
	}
	return
}

func (t *TeePrinter) Println(values ...interface{}) (n int, err error) {
	str := fmt.Sprint(values...)
	t.output = append(t.output, Decolorize(str))
	if !t.disableTerminalOutput {
		return fmt.Println(str)
	}
	return
}

func (t *TeePrinter) ForcePrint(values ...interface{}) (n int, err error) {
	str := fmt.Sprint(values...)
	t.output = append(t.output, Decolorize(str))
	return fmt.Print(str)
}

func (t *TeePrinter) ForcePrintf(format string, a ...interface{}) (n int, err error) {
	str := fmt.Sprintf(format, a...)
	t.output = append(t.output, Decolorize(str))
	return fmt.Print(str)
}

func (t *TeePrinter) ForcePrintln(values ...interface{}) (n int, err error) {
	str := fmt.Sprint(values...)
	t.output = append(t.output, Decolorize(str))
	return fmt.Println(str)
}

func (t *TeePrinter) DisableTerminalOutput(disable bool) {
	t.disableTerminalOutput = disable
}
