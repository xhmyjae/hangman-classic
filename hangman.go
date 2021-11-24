package main

type hangManData struct {
	HiddenWord []string
	ToFind string
	Attempts int
	Tried []string
	Position []string
}

func (w *hangManData) Init(HiddenWord []string, ToFind string, Attempts int, Tried []string, Position []string) {
	w.HiddenWord = HiddenWord
	w.ToFind = ToFind
	w.Attempts = Attempts
	w.Tried = Tried
	w.Position = Position
}