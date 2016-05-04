package prepend

func Prepend(v interface{}, slice []interface{}) []interface{}{
	return append([]interface{}{v}, slice...)
}