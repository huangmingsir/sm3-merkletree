package merkletree

import "fmt"

// Formatter formats a []byte in to a string.
// It is used by DOT() to provide users with the required format for the graphical display of their Merkle trees.
type Formatter interface {
	// Format
	Format([]byte) string
}

// TruncatedHexFormatter shows only the first and last two bytes of the value
type TruncatedHexFormatter struct{}

func (f *TruncatedHexFormatter) Format(data []byte) string {
	return fmt.Sprintf("%4x…%4x", data[0:2], data[len(data)-2:len(data)])
}

// HexFormatter shows the entire value
type HexFormatter struct{}

func (f *HexFormatter) Format(data []byte) string {
	return fmt.Sprintf("%0x", data)
}

// StringFormatter shows the entire value as a string
type StringFormatter struct{}

func (f *StringFormatter) Format(data []byte) string {
	return fmt.Sprintf("%s", string(data))
}
