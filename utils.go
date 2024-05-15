package go_testing

import (
	"fmt"

	"github.com/logrusorgru/aurora/v4"
)

func SuccessMessage(h, m string) {
	hf := aurora.Green(h + ":").Bold()
	mf := aurora.Green(m)
	fmt.Printf("%s %s\n", hf, mf)
}

func ErrorMessage(h, m string) {
	hf := aurora.Red(h + ":").Bold()
	mf := aurora.Red(m)
	fmt.Printf("%s %s\n", hf, mf)
}
