package kernelSrc_test

import (
	"fmt"
	kernelSrc "github.com/ArkNX/ark-go/plugin/kernel/src"
	"testing"
)

func TestAFData(t *testing.T) {
	data := kernelSrc.NewAFData(nil)
	fmt.Println(data.GetType())

	data1 := kernelSrc.NewAFData([]int{1, 23})
	fmt.Println(data1.GetType())

}
