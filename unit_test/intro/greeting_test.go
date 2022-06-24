package intro

import (
	"testing"
)

func TestGreeting(t *testing.T) {
	expected := "Selamat Pagi"
	actual := Greeting()
	if expected != actual {
		t.Error("Seharusnya mencetak \"Selamat Pagi\"")
	}
}

func TestSayHi(t *testing.T) {
	expected := "Hai salam kenal"
	actual := SayHi()
	if expected != actual {
		t.Error("Seharusnya mencetak \"Hai salam kenal\"")
	}
}

// CEKCOVERAGE go test ./... -cover -coverprofile=c.out && go tool cover -html=c.out -o coverage.html

// func BenchmarkGreeting2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Greeting2()
// 	}
// }
// // go test ./... -v -bench=.
