package filter_test

import (
	f "mini/functions/filter"
	"testing"
)

func TestFilter(t *testing.T) {
	var s = "fuck you"
	var excpted = "* you"
	var sed = f.Filter(s)
	if excpted != sed {
		t.Errorf("s = %s , excpted = %s , sed = %s", s, excpted, sed)
	}
}
