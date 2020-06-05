package log

import (
	"fmt"
	"testing"
)

func TestInitLogger(t *testing.T) {
	InitLogger(true)

	//var test uint64
	test := "hello"

	fmt.Printf("github.com/sony/sonyflake:      %x\n", test)

	Info("info test")
	Infof("info %x", test)
	InfoW("infow test", "infow", test)

	Error("Error test")
	Errorf("Error %s", test)
	Errorw("Error test", "Error", test)

	Warn("Error test")
	Warnf("Error %s", test)
	Warnw("Errorw test", "Error", test)
}
