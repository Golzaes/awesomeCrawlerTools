package Head

import (
	"fmt"
	"testing"
)

func TestRandUa(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(RandomUserAgent())
	}
}
