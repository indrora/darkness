package lri

import (
	"bytes"
	"encoding/binary"
	"errors"
	"image/color"

	"image"

	"github.com/indrora/darkness/shine/lri/ltpb"
	"github.com/pektezol/bitreader"
)

type LIRImagePart interface {
	Decode(data []byte) error
	GetImages() ([]image.Image, error)
}

type PackedRawImage struct {
	Size     ltpb.Point2I
	Offset   ltpb.Point2I
	Stride   uint64
	DataSize uint64
	Data     []uint16
}

func (rImg *PackedRawImage) Decode(data []byte) error {
	reader := bitreader.NewReaderFromBytes(data, true)
	if reader == nil {
		return errors.New("??")
	}
	rImg.DataSize = uint64((*rImg.Size.X) * (*rImg.Size.Y))
	rImg.Data = make([]uint16, rImg.DataSize)
	for i := 0; i < int(rImg.DataSize); i++ {
		pixel, err := reader.ReadBits(10)
		if err != nil {
			return err
		} else {
			rImg.Data[i] = uint16(pixel)
		}
	}

	return nil
}

func (rImg *PackedRawImage) GetImages() ([]image.Image, error) {

	// Construct an image

	raw := image.NewGray16(image.Rect(0, 0, int(rImg.Size.GetX()), int(rImg.Size.GetY())))
	for y := 0; y < int(rImg.Size.GetY()); y++ {
		for x := 0; x < int(rImg.Size.GetX()); x++ {
			pxcolor := rImg.Data[rImg.Stride*uint64(y)+uint64(x)]
			raw.Set(x, y, color.Gray16{Y: pxcolor})
		}
	}

	return []image.Image{raw}, nil
}

type JpegRawImage struct {
	Size       ltpb.Point2I
	Offset     ltpb.Point2I
	Monochrome bool
	Lengths    [4]uint32
	ExtraData  [1552]byte
	Data       [][]byte
	SensorType ltpb.SensorType
}

func (jpImage *JpegRawImage) Decode(data []byte) error {

	// Check that the header is correct.

	magic := data[0:4]

	if !bytes.Equal(magic, []byte{66, 74, 80, 71}) {
		return errors.New("Not a BayerJPEG")
	}

	jpImage.Monochrome = binary.LittleEndian.Uint32(data[4:8]) == 1

	jpImage.Lengths[0] = binary.LittleEndian.Uint32(data[8:12])
	jpImage.Lengths[1] = binary.LittleEndian.Uint32(data[12:16])
	jpImage.Lengths[2] = binary.LittleEndian.Uint32(data[16:20])
	jpImage.Lengths[3] = binary.LittleEndian.Uint32(data[20:24])

	// Copy whatever the extra data is in there. It might be useful.
	for i := 0; i < 1552; i++ {
		jpImage.ExtraData[i] = data[24+i]
	}

	// Next, get the four JPEGs
	if jpImage.Monochrome {
		jpImage.Data = make([][]byte, 1)
		jpImage.Data[0] = data[24+jpImage.Lengths[0] : 24+jpImage.Lengths[0]+jpImage.Lengths[1]]
	} else {
		jpImage.Data = make([][]byte, 4)
		for i := 0; i < 4; i++ {
			jpImage.Data[i] = data[24+jpImage.Lengths[i] : 24+jpImage.Lengths[i]+jpImage.Lengths[i+1]]
		}
	}

	return nil
}

func (jpImage *JpegRawImage) GetImages() ([]image.Image, error) {

	images := make([]image.Image, len(jpImage.Data))
	for i := 0; i < len(jpImage.Data); i++ {
		images[i], _, _ = image.Decode(bytes.NewReader(jpImage.Data[i]))
	}

	return images, nil
}

type RawImageData struct {
	Module  *ltpb.CameraModule
	Surface *ltpb.CameraModule_Surface
	Content LIRImagePart
}

func NewRawImage(module *ltpb.CameraModule, surface *ltpb.CameraModule_Surface, data []byte) (*RawImageData, error) {
	rawImage := &RawImageData{
		Module:  module,
		Surface: surface,
	}
	if surface.GetFormat() == ltpb.CameraModule_Surface_RAW_BAYER_JPEG {
		rawImage.Content = &JpegRawImage{}
		err := rawImage.Content.Decode(data)
		if err != nil {
			return nil, err
		}
	} else if surface.GetFormat() == ltpb.CameraModule_Surface_RAW_PACKED_10BPP {
		rawImage.Content = &PackedRawImage{}
		err := rawImage.Content.Decode(data)
		if err != nil {
			return nil, err
		}

	}

	return rawImage, nil
}
