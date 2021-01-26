package main

import (
	"fmt"
	"os"

	"github.com/erikh/colorwriter"
	"github.com/fatih/color"
)

func main() {
	w := colorwriter.NewWriter(os.Stdout, 2, 2, 2, ' ', 0)
	w.Write([]byte("FOO\tBAR\tQUUX\n"))
	w.Write([]byte(fmt.Sprintf("%s\t%s\t%s\n", color.RedString("test"), color.BlueString("test2"), color.CyanString("blah"))))
	w.Write([]byte(fmt.Sprintf("%s\t%s\t%s\n", color.RedString("testing"), color.BlueString("test2 oh my"), color.CyanString("blah"))))
	w.Flush()
}
