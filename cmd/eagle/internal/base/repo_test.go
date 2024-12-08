package base

import (
	"context"
	"testing"
)

func TestRepo(t *testing.T) {
	r := NewRepo("https://github.com/linqining/eagle-layout.git", "main")
	if err := r.Clone(context.Background()); err != nil {
		t.Fatal(err)
	}
	if err := r.CopyTo(context.Background(), "/tmp/test_eagle_repo", "github.com/linqining/eagle-layout", nil); err != nil {
		t.Fatal(err)
	}
}
