package std

import (
  "fmt"
  "os"
)

func Error(err error, code int) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(code)
	}
}
