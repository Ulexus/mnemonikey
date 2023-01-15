package mnemonikey

import (
	"bytes"
	"math/big"
	"time"
)

// EpochIncrement is the level of granularity available for the creation date of
// keys generated by mnemonikey.
const EpochIncrement = time.Second

// EpochStart is the start of the epoch after which key creation times are encoded
// in recovery phrases. It is exactly midnight in UTC time on the new year's eve
// between 2022 and 2023.
//
// In unix time, this epoch is exactly 1672531200 seconds after the unix epoch.
var EpochStart = time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

// checksumTable is the precomputed table used to create checksums of backup payloads.
const (
	// VersionLatest is the latest known mnemonikey version number. Backups encoded with versions higher
	// than this number will fail to decode.
	VersionLatest uint = 0

	// VersionBitCount is the number of bits in the backup payload reserved for the version number.
	VersionBitCount uint = 2

	// CreationOffsetBitCount is the number of bits used to represent a key creation offset.
	CreationOffsetBitCount uint = 30

	// ChecksumBitCount is the number of bits in the backup payload reserved for the checksum.
	ChecksumBitCount uint = 5

	// EntropyBitCount is the number of bits of entropy in the seed used to derive PGP keys.
	EntropyBitCount = 128

	// MnemonicSize is the number of mnemonic words needed to encode
	// both the key creation offset and 128 bits of seed entropy.
	MnemonicSize uint = 15
)

const checksumMask = (1 << ChecksumBitCount) - 1

var entropyMask = new(big.Int).SetBytes(bytes.Repeat([]byte{0xFF}, EntropyBitCount/8))

// MaxCreationTime is the farthest point in the future that the mnemonikey recovery phrase
// encoding algorithm can represent key creation timestamps for.
var MaxCreationTime = EpochStart.Add(EpochIncrement * (time.Duration(1<<CreationOffsetBitCount) - 1))
