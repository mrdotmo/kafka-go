package kafka

import "errors"

type Header struct {
	correlationId [4]byte
}

// parseHeader extracts the correlation ID from the header.
//
//lint:ignore U1000 Disable unused lint check until later.
func parseHeader(header []byte) (Header, error) {
	var rv Header
	if len(header) < 8 {
		return rv, errors.New("header must have at least 8 bytes")
	}

	copy(rv.correlationId[:], header[4:8])
	return rv, nil
}
