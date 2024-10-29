// Package acdc asks if you are ready to rock
package acdc

// Sum adds an unlimited number of of type int
func Sum(x1 ... int) int {
	s := 0
	for _, v := range x1 {
		s += v
	}
	return s
}