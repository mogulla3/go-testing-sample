package calc

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Write setup code here
	fmt.Println("Before running test")

	// Run test
	code := m.Run()

	// Write teardown code here
	fmt.Println("After running test")

	os.Exit(code)
}

func TestAdd(t *testing.T) {
	expected := 3
	actual := Add(1, 2)

	if expected != actual {
		t.Errorf("expected: %v , actual: %v", expected, actual)
	}
}

func TestSub(t *testing.T) {
	expected := 1
	actual := Sub(2, 1)

	if expected != actual {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestMul(t *testing.T) {
	expected := 6
	actual := Mul(2, 3)

	if expected != actual {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestDiv(t *testing.T) {
	expected := 3
	actual := Div(6, 2)

	if expected != actual {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
