package area

func CalculateArea(length, width int) (area int, e bool) {
	if length > 0 && width > 0 {
		area = length * width
		e = true
		return
	} else {
		area = 0
		e = true
		return
	}
}
