package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := parseFlags()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		// run the tests
		os.Exit(m.Run())
	}
}

func TestParsing(t *testing.T) {
	if len(bc.Receptors) != 2 {
		t.Log(fmt.Sprintf("expected 2 receptors in the config, got %d instead", len(bc.Receptors)))
		t.Fail()
	} else if len(bc.Processors) != 2 {
		t.Log(fmt.Sprintf("expected 2 processors in the config, got %d instead", len(bc.Processors)))
		t.Fail()
	} else if len(bc.Routes) != 2 {
		t.Log(fmt.Sprintf("expected 2 routes in the config, got %d instead", len(bc.Routes)))
		t.Fail()
	}

	if t.Failed() {
		return
	}
}
