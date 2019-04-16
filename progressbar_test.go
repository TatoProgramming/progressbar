package progressbar

import "testing"

func TestBar(t *testing.T){
	want := 10
	got := Bar(want)
	if got.total != want  {
		t.Errorf("The progress bars total was not %q", want)
	}
}