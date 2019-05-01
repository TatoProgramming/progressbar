package progressbar

import (
	"bytes"
	"fmt"
)

// Bar
type Progressbar struct {
	current, total int
	currentTime    float64
	Width          int
	theme          Theme
}

type Theme struct {
	Start, Fill, Head, Space, End string
}

func Bar(total int) Progressbar {
	return Progressbar{
		total: total,
		theme: Theme{Start: "|", Fill: "=", Head: ">", Space: " ", End: "|"},
	}
}

func (b *Progressbar) ChangeTheme(theme Theme) {
	b.theme = theme
}

func (b *Progressbar) addInt(increment int) {
	if increment < 0 {
		return
	}
	if b.current+increment <= b.total {
		b.current += increment
	} else {
		b.current = b.total
	}
}

func (b *Progressbar) Add(increment int){
	b.addInt(increment)
	go b.Display()
}

func (b *Progressbar) Display() {
	if b.Width == 0 {
		b.Width = 50
	}
	fmt.Printf("\r %s", b.DisplayString())
}

func (b *Progressbar) DisplayString() string{
	percentage := (float64(b.current) / float64(b.total))

	var buffer bytes.Buffer
	amount := int(float64(b.Width) * percentage)
	for i := 0; i < amount; i++ {
		buffer.WriteString(b.theme.Fill)
	}
	if percentage < 1.0 {
		buffer.WriteString(b.theme.Head)
	}
	for i := 0; i < b.Width-amount-1; i++ {
		buffer.WriteString(b.theme.Space)
	}
	return fmt.Sprintf("%s%s%s %3.0f%%",
		b.theme.Start,
		buffer.String(),
		b.theme.End,
		percentage*100.0,
	)
}
