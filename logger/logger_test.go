package logger

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestNoInitialization(t *testing.T) {
	if Trace != nil {
		t.Errorf("Trace should not yet be Initialized!")
	}

	if Info != nil {
		t.Errorf("Info should not yet be Initialized!")
	}

	if Warning != nil {
		t.Errorf("Warning should not yet be Initialized!")
	}

	if Error != nil {
		t.Errorf("Error should not yet be Initialized!")
	}

	if Initialized != false {
		t.Errorf("Initialized should not yet be Initialized!")
	}
}

func TestInitialization(t *testing.T) {
	Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	if Trace == nil {
		t.Errorf("Trace should be Initialized!")
	}

	if Info == nil {
		t.Errorf("Info should be Initialized!")
	}

	if Warning == nil {
		t.Errorf("Warning should be Initialized!")
	}

	if Error == nil {
		t.Errorf("Error should be Initialized!")
	}

	if Initialized == false {
		t.Errorf("Initialized should be Initialized!")
	}
}
