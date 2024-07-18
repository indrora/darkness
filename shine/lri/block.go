package lri

import (
	"encoding/binary"
	"errors"
	"io"
)

type BType uint8

const (
	BType_LightHeader BType = iota
	BType_ViewPreferences
	BType_GPSData
)

var (
	ErrBodyMismatch    = errors.New("body mismatch")
	ErrMessageMismatch = errors.New("message mismatch")
)

type BlockPreamble struct {
	Magic         [4]byte
	BlockLen      uint64
	MessageOffset uint64
	MessageLength uint32
	BlockType     BType
	reserved      [7]byte
}

type Block struct {
	Body      []byte
	Message   []byte
	BlockType BType
}

func NewBlock(reader io.ReadSeeker) (*Block, error) {

	err := errors.New("empty")
	preamble := &BlockPreamble{}

	if err = binary.Read(reader, binary.LittleEndian, &preamble); err != nil {
		return nil, err
	}

	block := &Block{
		BlockType: preamble.BlockType,
	}

	bodyLength := preamble.BlockLen - uint64(binary.Size(preamble)) - uint64(preamble.MessageLength)
	bodyBuff := make([]byte, bodyLength)
	messageBuff := make([]byte, preamble.MessageLength)
	if blen, err := reader.Read(bodyBuff); blen != int(bodyLength) || err != nil {
		return nil, errors.Join(ErrBodyMismatch, err)
	}
	if mlen, err := reader.Read(messageBuff); mlen != int(preamble.MessageLength) || err != nil {
		return nil, errors.Join(ErrMessageMismatch, err)
	}
	block.Body = bodyBuff
	block.Message = messageBuff

	return block, nil

}
