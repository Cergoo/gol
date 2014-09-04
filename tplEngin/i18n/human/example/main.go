package main

import (
	"fmt"
	"github.com/Cergoo/gol/tplEngin/i18n/human"
)

func main() {
	vbyten_En := human.GetBytenHumanize([]string{"B", "KB", "MB", "GB", "TB", "PB", "EB"})
	vbyten_Ru := human.GetBytenHumanize([]string{"Б", "КБ", "МБ", "ГБ", "ТБ", "ПБ", "ЕБ"})
	vbiten_En := human.GetBytenHumanize([]string{"b", "Kb", "Mb", "Gb", "Tb", "Pb", "Eb"})
	v := uint64(10000000)
	fmt.Print(vbyten_En(v), "\n", vbyten_Ru(v), "\n", vbiten_En(v))
}
