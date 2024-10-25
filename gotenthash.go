package gotenthash

import (
	"encoding/binary"
)

var (
	defaultState = [4]uint64{
		0x5d6daffc4411a967,
		0xe22d4dea68577f34,
		0xca50864d814cbc2e,
		0x894e29b9611eb173,
	}
	defaultRotations = [][2]uint{
		{16, 28},
		{14, 57},
		{11, 22},
		{35, 34},
		{57, 16},
		{59, 40},
		{44, 13},
	}
)

const (
	DigestSize = 160 / 8 // Digest size, in bytes.
	BlockSize  = 256 / 8 // Internal block size of the hash, in bytes.
)

type TentHasher struct {
	state         [4]uint64
	buf           [BlockSize]byte
	bufLength     int
	messageLength uint64
}

func New() *TentHasher {
	return &TentHasher{
		state: defaultState,
	}
}

func (t *TentHasher) Reset() {
	t.state = defaultState
	t.bufLength = 0
	t.messageLength = 0
}

func (t *TentHasher) Write(data []byte) (int, error) {
	n := len(data)
	t.messageLength += uint64(n)

	for len(data) > 0 {
		if t.bufLength == 0 && len(data) >= BlockSize {
			xorDataIntoState(&t.state, data)
			mixState(&t.state)
			data = data[BlockSize:]
		} else if t.bufLength == BlockSize {
			xorDataIntoState(&t.state, t.buf[:])
			mixState(&t.state)
			t.bufLength = 0
		} else {
			toCopy := BlockSize - t.bufLength
			if toCopy > len(data) {
				toCopy = len(data)
			}
			copy(t.buf[t.bufLength:], data[:toCopy])
			data = data[toCopy:]
			t.bufLength += toCopy
		}
	}

	return n, nil
}

func (t *TentHasher) Sum(b []byte) []byte {
	// Create a copy of the current state
	clone := *t

	// Hash the remaining bytes if there are any
	if clone.bufLength > 0 {
		for i := clone.bufLength; i < BlockSize; i++ {
			clone.buf[i] = 0
		}
		xorDataIntoState(&clone.state, clone.buf[:])
		mixState(&clone.state)
	}

	// Incorporate the message length (in bits) and do the final mixing
	clone.state[0] ^= clone.messageLength * 8
	mixState(&clone.state)
	mixState(&clone.state)

	// Get the digest as a byte array
	digest := make([]byte, DigestSize)
	binary.LittleEndian.PutUint64(digest[0:8], clone.state[0])
	binary.LittleEndian.PutUint64(digest[8:16], clone.state[1])
	binary.LittleEndian.PutUint32(digest[16:20], uint32(clone.state[2]))

	return append(b, digest...)
}

func xorDataIntoState(state *[4]uint64, data []byte) {
	state[0] ^= binary.LittleEndian.Uint64(data[0:8])
	state[1] ^= binary.LittleEndian.Uint64(data[8:16])
	state[2] ^= binary.LittleEndian.Uint64(data[16:24])
	state[3] ^= binary.LittleEndian.Uint64(data[24:32])
}

func mixState(state *[4]uint64) {
	rotations := defaultRotations
	for _, rot := range rotations {
		state[0] = state[0] + state[2]
		state[2] = (state[2] << rot[0]) | (state[2] >> (64 - rot[0]))
		state[2] ^= state[0]
		state[1] = state[1] + state[3]
		state[3] = (state[3] << rot[1]) | (state[3] >> (64 - rot[1]))
		state[3] ^= state[1]

		state[0], state[1] = state[1], state[0]
	}
}

func Hash(data []byte) [DigestSize]byte {
	h := New()
	h.Write(data)
	var digest [DigestSize]byte
	copy(digest[:], h.Sum(nil))
	return digest
}
