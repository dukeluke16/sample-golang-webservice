package main

import (
	"errors"
	"testing"
)

func TestMainSuccess(t *testing.T) {
	start = func() error {
		return nil
	}

	defer func() {
		if recover() != nil {
			t.Error("Failure launching main.")
		}
	}()
	main()
}

func TestMainFailure(t *testing.T) {
	start = func() error {
		return errors.New("triggering failure to launch main")
	}

	fatal = func(v ...interface{}) {
	}

	defer func() {
		if recover() != nil {
			t.Error("Expected a failure launching main!")
		}
	}()
	main()
}
