package wasm

import "fmt"

type EncoderRegistry struct {
	encoders map[string]Encoder
}

// NewEncoderRegistry creates a new registry for message encoders.
func NewEncoderRegistry() *EncoderRegistry {
	return &EncoderRegistry{
		encoders: make(map[string]Encoder),
	}
}

// RegisterEncoder adds a message encoder for the given route.
func (qr *EncoderRegistry) RegisterEncoder(route string, encoder Encoder) {
	if _, exists := qr.encoders[route]; exists {
		panic(fmt.Sprintf("wasm: encoder already registered for route: %s", route))
	}
	qr.encoders[route] = encoder
}
