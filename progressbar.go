package progressbar

// Bar 
type Progressbar struct {
	total int
}

func Bar(total int) Progressbar {
	return Progressbar{total}
}