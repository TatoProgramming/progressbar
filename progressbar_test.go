package progressbar

import "testing"

func TestBar(t *testing.T) {
	want := 10
	got := Bar(want)
	if got.total != want {
		t.Errorf("The progress bars total was %q expecting %q", want, got.total)
	}
}

func TestAdd(t *testing.T) {
	bar := Bar(100)
	cases := []struct {
		in, want int
	}{
		{10, 10},
		{10, 20},
		{20, 40},
		{20, 60},
		{40, 100},
		{10, 100},
	}
	for _, c := range cases {
		bar.Add(c.in)
		if bar.current != c.want {
			t.Errorf("Adding %d resulted in %d expected %d", c.in, bar.current, c.want)
		}
	}
}


func TestChangeTheme(t *testing.T) {
	bar := Bar(100)
	defaultTheme := bar.theme
	theme := Theme{start: "(", fill: "¯", head: "-", space: "_", end: ")"}

	bar.Add(1)
	if bar.current != c.want {
		t.Errorf("Adding %d resulted in %d expected %d")
	}
}