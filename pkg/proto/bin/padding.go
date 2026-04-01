package bin

import "fmt"

func validatePaddingZeros(b []byte, dataLen, paddedLen int) error {
	if paddedLen < dataLen || len(b) < paddedLen {
		return ErrUnexpectedEOF
	}
	for i := dataLen; i < paddedLen; i++ {
		if b[i] != 0 {
			return fmt.Errorf("non-zero tl padding at offset %d", i)
		}
	}
	return nil
}
