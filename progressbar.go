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
}

func Bar(total int) Progressbar {
	return Progressbar{total: total}
}

func (b *Progressbar) Add(increment int) {
	if increment < 0 { return }
	if b.current+increment <= b.total {
		b.current += increment
	} else {
		b.current = b.total
	}
	go b.Display()
}

func (b *Progressbar) Display() {
	if b.Width == 0 {
		b.Width = 50
	}

	percentage := (float64(b.current) / float64(b.total))
	whole := "="

	var buffer bytes.Buffer
	amount := int(float64(b.Width) * percentage)
	for i := 0; i < amount; i++ {
		buffer.WriteString(whole)
	}
	if percentage < 1.0 {
		buffer.WriteString(">")
	}
	for i := 0; i < b.Width-amount-1; i++ {
		buffer.WriteString(" ")
	}
	fmt.Printf("\r |%s| %3.0f%% - 5m 10s 30ms", buffer.String(), percentage*100.0)
}
