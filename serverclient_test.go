package GoServerDetect_test

import (
	"github.com/KairosSystems/GoServerDetect"
	"testing"
)

func ServerTest(t *testing.T) {
	response := []byte("Response from server")
	err := GoServerDetect.CreateServer(91838899, "testingPassword", response)
	if err != nil {
		t.Fail(err)
	}
}
