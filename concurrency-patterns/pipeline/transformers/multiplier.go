package transformers

func Multiplier(k int) Transformer {
	return func(v interface{}) interface{} {
		return v.(int) * k
	}
}
