package kafka

import "errors"

// parseHeader extracts the correlation ID from the header.
func parseHeader(header []byte) ([]byte, error) {
	if len(header) < 8 {
		return nil, errors.New("header must have at least 8 bytes")
	}

	correlationId := header[4:8]
	return correlationId, nil
}
