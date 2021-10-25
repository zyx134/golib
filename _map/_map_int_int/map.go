package _map_int_int

func Copy(m map[int]int) map[int]int {
	dst := make(map[int]int)
	for k, v := range m {
		dst[k] = v
	}
	return dst
}
