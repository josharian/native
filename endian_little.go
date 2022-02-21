//go:build amd64 || 386 || arm || arm64 || loong64 || mipsle || mips64le || ppc64le || riscv64 || wasm
// +build amd64 386 arm arm64 loong64 mipsle mips64le ppc64le riscv64 wasm

package native

import "encoding/binary"

var Endian = binary.LittleEndian
