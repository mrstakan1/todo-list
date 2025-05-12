package service

import "testing"

// seed-корпус — реальные строки
var seeds = []string{
	"Купить молоко", "Fix bug #42",
}

func FuzzValidateTitle(f *testing.F) {
	for _, s := range seeds {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, in string) {
		defer func() { // ловим паники, если они будут
			if r := recover(); r != nil {
				t.Fatalf("panic: %v\ninput: %q", r, in)
			}
		}()
		_ = ValidateTitle(in) // проверяем, что не упадёт
	})
}
