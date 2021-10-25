package _slice

/*
 * 清空Slice，传入的slice对象地址发生变化。
 * params:
 *   s: slice对象指针，类型为*[]interface{}
 * return:
 *   无
 */
func Clear(s *[]interface{}) {
	*s = append([]interface{}{})
}

/*
 * 清空Slice，传入的slice对象地址不变。
 * params:
 *   s: slice对象指针，类型为*[]interface{}
 * return:
 *   无
 */
func Clear2(s *[]interface{}) {
	*s = (*s)[0:0]
}
