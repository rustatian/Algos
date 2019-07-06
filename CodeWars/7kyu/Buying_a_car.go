package main

func main() {
	NbMonths(7500, 32000, 300, 1.55) //[2]int{6, 766}
}

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	//var saved int
	var spo float64 = float64(startPriceOld)
	var spn float64 = float64(startPriceNew)
	var spm float64 = 0
	var plbm float64 = percentLossByMonth / 100
	var monts int

	var res [2]int
	for monts = 0; spm+spo < spn; monts++ {

		spm += float64(savingperMonth)
		spo = spo * (1 - plbm)
		spn = spn * (1 - plbm)

		if (monts % 2) == 0 {
			plbm += 0.005
		}
	}

	res[0] = monts
	val := Round(spm + spo - spn)
	res[1] = int(val)

	return res
}

func Round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}
