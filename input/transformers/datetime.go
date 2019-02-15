package transformers

import "time"

type DateTimeTransformer struct {
}

func (dt DateTimeTransformer) Transform(value interface{}) interface{} {
    t, _ := time.Parse("2006-01-02 15:04:05", value.(string))

    return t
}