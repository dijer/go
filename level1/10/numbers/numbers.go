package numbers

import "strconv"

func translateNumber(number int) string {
	units := number % 10
	tens := int((number % 100) / 10)
	hundreds := int((number % 1000) / 100)

	return strconv.Itoa(hundreds) + " сотен, " + strconv.Itoa(tens) + " десятков, " + strconv.Itoa(units) + " единиц"
}
