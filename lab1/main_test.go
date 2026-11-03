package lab1

import (
	"fmt"
	"os"
	"testing"

	"github.com/quickfeed/quickfeed/kit/score"
)

var scores = score.NewRegistry()

func TestMain(m *testing.M) {
	scores.PrintTestInfo()
	code := m.Run()
	if err := scores.Validate(); err != nil {
		fmt.Println(err)
	}
	os.Exit(code)
}