package daemon

import (
	"strings"

	"github.com/brendandebeasi/tabby/pkg/tmux"
)

// uahWorkingTitle recognizes UAH's native title, not a project name containing
// "working". Require the current command too: tmux retains titles after exit.
// UAH's idle title has no state field; approvals use "approve?" instead.
func uahWorkingTitle(pane tmux.Pane) bool {
	if pane.Command != "uah" {
		return false
	}
	fields := strings.Split(pane.Title, " · ")
	if len(fields) < 3 || fields[0] != "uah" {
		return false
	}
	state := 1
	if fields[state] == "yolo" {
		state++
	}
	// A workspace must follow the state, so an idle workspace named "working"
	// (including in yolo mode) cannot keep a stale busy hook alive.
	return len(fields) > state+1 && fields[state] == "working" && fields[state+1] != ""
}
