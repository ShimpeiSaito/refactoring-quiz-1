package main

import (
	"fmt"
	"time"
)

type Portion int // 麺の量

const (
	Regular Portion = 0
	Small   Portion = 1
	Large   Portion = 2
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(p Portion, aburaage bool, ebiten uint) *Udon {
	flag1 := 0 == ebiten
	flag2 := time.Now().Hour() < 10
	if flag1 && flag2 {
		return &Udon{
			men:      p,
			aburaage: aburaage,
			ebiten:   ebiten + 1,
		}
	} else {
		return &Udon{
			men:      p,
			aburaage: aburaage,
			ebiten:   ebiten,
		}
	}
}

func main() {
	fmt.Println(NewUdon(Large, true, 0))
}

// go run question3.go