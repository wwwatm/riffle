package lib

import (
	"bytes"

	"github.com/dedis/crypto/abstract"
)

func SetBit(n_int int, b bool, bs []byte) {
	n := uint(n_int)
	if b {
		bs[n/8] |= 1 << (n % 8)
	} else {
		bs[n/8] &= ^(1 << (n % 8))
	}
}

func Xor(r []byte, w []byte) {
	for i, b := range r {
		w[i] ^= b
	}
}

func Xors(bss [][]byte) []byte {
	n := len(bss[0])
	x := make([]byte, n)
	for _, bs := range bss {
		for i, b := range bs {
			x[i] ^= b
		}
	}
	return x
}

func AllZero(xs []byte) bool {
	for _, x := range xs {
		if x != 0 {
			return false
		}
	}
	return true
}

func MarshalPoint(pt abstract.Point) []byte {
	buf := new(bytes.Buffer)
	ptByte := make([]byte, SecretSize)
	pt.MarshalTo(buf)
	buf.Read(ptByte)
	return ptByte
}

func UnmarshalPoint(ptByte []byte) abstract.Point {
	buf := bytes.NewBuffer(ptByte)
	pt := Suite.Point()
	pt.UnmarshalFrom(buf)
	return pt
}
