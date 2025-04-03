package transformers

func Adder(k int) Transformer {
	return func(v interface{}) interface{} {
		return v.(int) + k
	}
}
