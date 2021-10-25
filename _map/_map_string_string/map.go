package _map_string_string

// 取得map的key
func Keys(m map[string]string) []string {
	s := make([]string, 0, len(m))

	for k := range m {
		s = append(s, k)
	}

	return s
}

// 取得map的值
func Values(m map[string]string) []string {
	s := make([]string, 0, len(m))

	for k := range m {
		s = append(s, m[k])
	}

	return s
}

// 合并map
func Merge(m map[string]string, more ...map[string]string) map[string]string {
	if nil == m {
		m = make(map[string]string)
	}
	for _, v := range more {
		for k_, v_ := range v {
			m[k_] = v_
		}
	}

	return m
}

func Copy(m map[string]string) map[string]string {
	dst := make(map[string]string)
	for k, v := range m {
		dst[k] = v
	}
	return dst
}

//根据key删除map的value
func DelValue(m map[string]string, key string) map[string]string {
	delete(m, key)
	return m
}

//清空整个map
func Clear(m map[string]string) map[string]string {
	m = nil //清空map 元素
	return m
}

//判断Key是否存在
func KeyExist(m map[string]string, key string) bool {
	if _, ok := m[key]; ok {
		return true
	}
	return false
}
