package transformers

func Incrementor(v interface{}) interface{} {
	return v.(int) + 1
}
