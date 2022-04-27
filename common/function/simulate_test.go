package function

import (
	"fmt"
	"testing"
)

func TestSimulate(t *testing.T) {
	result := GenerateUserData(1000)
	fmt.Println(result)
}
