package native_test

import (
	"encoding/binary"
	"testing"

	"github.com/josharian/native"
)

func TestPrintEndianness(t *testing.T) {
	t.Logf("native endianness is %v", native.Endian)

	var want binary.ByteOrder = binary.BigEndian
	if !native.IsBigEndian {
		want = binary.LittleEndian
	}
	if native.Endian != want {
		t.Errorf("IsBigEndian = %v not consistent with native.Endian = %T", native.IsBigEndian, want)
	}
}
