package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/erikh/colorwriter"
	"github.com/fatih/color"
)

func main() {
	fmt.Println("This prints two tables with colors in them: the first is text/tabwriter, the second is colorwriter.")
	fmt.Println("Outside of this, they are no different. Enjoy!")

	fmt.Println()
	fmt.Println("text/tabwriter")
	fmt.Println()

	tab := tabwriter.NewWriter(os.Stdout, 2, 2, 2, ' ', 0)
	tab.Write([]byte("FOO\tBAR\tQUUX\n"))
	tab.Write([]byte(fmt.Sprintf("%s\t%s\t%s\n", color.RedString("test"), color.BlueString("test2"), color.CyanString("blah"))))
	tab.Write([]byte(fmt.Sprintf("%s\t%s\t%s\n", color.RedString("testing"), color.BlueString("test2 oh my"), color.CyanString("blah"))))
	tab.Flush()

	fmt.Println()
	fmt.Println("colorwriter")
	fmt.Println()

	w := colorwriter.NewWriter(os.Stdout, 2, 2, 2, ' ', 0)
	w.Write([]byte("FOO\tBAR\tQUUX\n"))
	w.Write([]byte(fmt.Sprintf("%s\t%s\t%s\n", color.RedString("test"), color.BlueString("test2"), color.CyanString("blah"))))
	w.Write([]byte(fmt.Sprintf("%s\t%s\t%s\n", color.RedString("testing"), color.BlueString("test2 oh my"), color.CyanString("blah"))))
	w.Flush()
}
