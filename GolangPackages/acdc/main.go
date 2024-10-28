package acdc

func Sum(x1 ... int) int {
	s := 0
	for _, v := range x1 {
		s += v
	}
	return s
}