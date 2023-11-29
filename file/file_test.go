package file

import (
	"os"
	"testing"
)

func TestFile(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		if err := os.WriteFile("path", []byte("hello world"), 493); err != nil {
			panic(err)
		}
	})
}
