package product

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type Cloner interface {
	GetDescription() string
	SetDescription(string)
}

func Clone(src Cloner, dst Cloner) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(src)
	if err != nil {
		return fmt.Errorf("encoding failed: %s", err)
	}
	dec := gob.NewDecoder(&buf)
	err = dec.Decode(dst)
	if err != nil {
		return fmt.Errorf("decoding failed: %s", err)
	}
	return nil
}
