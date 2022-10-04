package main

import (
	"fmt"
	"time"
)

type Portion int // 麺の量

const (
	Regular Portion = iota
	Small
	Large
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(p Portion, aburaage bool, ebiten uint) *Udon {
	if ebiten == 0 && time.Now().Hour() < 10 {
		ebiten += 1
	}
	return &Udon{
		men:      p,
		aburaage: aburaage,
		ebiten:   ebiten,
	}
}

func main() {
	fmt.Println(NewUdon(Large, true, 0))
}

// // go run question3_ans.go