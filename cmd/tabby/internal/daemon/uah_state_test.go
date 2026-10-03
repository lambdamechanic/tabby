package daemon

import (
	"testing"

	"github.com/brendandebeasi/tabby/pkg/tmux"
)

func TestUAHWorkingTitle(t *testing.T) {
	tests := []struct {
		name    string
		command string
		title   string
		working bool
	}{
		{"working", "uah", "uah · working · repo", true},
		{"yolo working", "uah", "uah · yolo · working · repo", true},
		{"idle workspace named working", "uah", "uah · working", false},
		{"idle yolo workspace named working", "uah", "uah · yolo · working", false},
		{"approval", "uah", "uah · approve? · repo", false},
		{"missing workspace", "uah", "uah · working · ", false},
		{"retained title after exit", "zsh", "uah · working · repo", false},
		{"unrelated title", "uah", "repo · working · task", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uahWorkingTitle(tmux.Pane{Command: tt.command, Title: tt.title})
			if got != tt.working {
				t.Fatalf("uahWorkingTitle = %t, want %t", got, tt.working)
			}
		})
	}
}
