package daemon

import "testing"

func TestWindowFocusRestoreTarget(t *testing.T) {
	tests := []struct {
		name      string
		active    string
		pending   string
		wantFocus string
	}{
		{
			name:      "no pending window",
			active:    "@1",
			pending:   "",
			wantFocus: "",
		},
		{
			name:      "pending matches active window",
			active:    "@1",
			pending:   "@1",
			wantFocus: "",
		},
		{
			name:      "pending new window differs from active window",
			active:    "@1",
			pending:   "@2",
			wantFocus: "@2",
		},
		{
			name:      "no active window but pending new window exists",
			active:    "",
			pending:   "@2",
			wantFocus: "@2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := windowFocusRestoreTarget(tt.active, tt.pending)
			if got != tt.wantFocus {
				t.Fatalf("windowFocusRestoreTarget(%q, %q) = %q, want %q", tt.active, tt.pending, got, tt.wantFocus)
			}
		})
	}
}

func TestRefreshFocusRestoreTarget(t *testing.T) {
	tests := []struct {
		name          string
		before, after string
		moves         int
		want          string
	}{
		{name: "grouped moves changed focus", before: "@1", after: "@2", moves: 3, want: "@1"},
		{name: "focus stayed put", before: "@1", after: "@1", moves: 3, want: ""},
		{name: "refresh did not renumber", before: "@1", after: "@2", moves: 0, want: ""},
		{name: "initial focus unavailable", before: "", after: "@2", moves: 3, want: ""},
		{name: "final focus unavailable", before: "@1", after: "", moves: 3, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := refreshFocusRestoreTarget(tt.before, tt.after, tt.moves); got != tt.want {
				t.Fatalf("refreshFocusRestoreTarget(%q, %q, %d) = %q, want %q", tt.before, tt.after, tt.moves, got, tt.want)
			}
		})
	}
}
