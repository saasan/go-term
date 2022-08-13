package term

import (
	"os"

	"golang.org/x/term"
)

func WaitAnyKey() error {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	_, err = os.Stdin.Read(make([]byte, 1))
	if err != nil {
		return err
	}

	return nil
}
