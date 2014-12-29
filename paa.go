package paa

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

//go:generate stringer -type=PaaType
type PaaType uint16

const (
	DXT1          PaaType = 0xFF01
	DXT2          PaaType = 0xFF02
	DXT3          PaaType = 0xFF03
	DXT4          PaaType = 0xFF04
	DXT5          PaaType = 0xFF05
	RGBA4444      PaaType = 0x4444
	RGBA5551      PaaType = 0x5551
	RGBA8888      PaaType = 0x8888
	GrayWithAlpha PaaType = 0x8080
)

func (p PaaType) IsValid() bool {
	switch p {
	case DXT1:
		return true
	case DXT2:
		return true
	case DXT3:
		return true
	case DXT4:
		return true
	case DXT5:
		return true
	case RGBA4444:
		return true
	case RGBA5551:
		return true
	case RGBA8888:
		return true
	case GrayWithAlpha:
		return true
	}

	return false
}

func ReadPaa(r io.Reader) ([]Tagg, error) {
	reader := bufio.NewReader(r)

	// Read the file magic
	var magic PaaType
	err := binary.Read(reader, binary.LittleEndian, &magic)

	if err != nil {
		return nil, fmt.Errorf("An error occurred: %s", err)
	}

	if !magic.IsValid() {
		return nil, fmt.Errorf("Invalid or unknown PAA type 0x%X", uint16(magic))
	}

	taggs := []Tagg{}
	for {
		// Read in a new tagg
		tagg := Tagg{}

		signature := make([]byte, 4)
		_, err := reader.Read(signature)

		if err != nil {
			return nil, err
		}

		reverse(&signature)

		tagg.signature = string(signature)

		if tagg.Signature() != TaggSignature {
			return nil, fmt.Errorf("Invalid TAGG signature (expected TAGG, got \"%s\"", tagg.Signature())
		}

		// Read the tag name
		taggType := make([]byte, 4)
		_, err = reader.Read(taggType)
		if err != nil {
			return nil, err
		}

		reverse(&taggType)

		tagg.name = string(taggType)

		// Read the length of the data
		err = binary.Read(reader, binary.LittleEndian, &tagg.dataLength)
		if err != nil {
			return nil, err
		}

		// Read in the data. The data is a bunch of 32-bit integers, and the length is the length
		// of the raw byte buffer I guess
		tagg.data = make([]int32, tagg.dataLength/4)

		var data int32
		for i := range tagg.data {
			err := binary.Read(reader, binary.LittleEndian, &data)
			if err != nil {
				return nil, err
			}

			tagg.data[i] = data
		}

		taggs = append(taggs, tagg)

		if tagg.Name() == OFFS {
			break
		}
	}

	return taggs, nil
}
