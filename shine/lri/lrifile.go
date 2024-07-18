// shine/lri: LRI file format decoding

package lri

import (
	"io"

	"github.com/indrora/darkness/shine/lri/ltpb"
	"google.golang.org/protobuf/proto"
)

type LRIFile struct {
	reader io.ReadSeeker
	//header *ltpb.LightHeader

	Images []RawImageData
}

func NewLRI(reader io.ReadSeeker) (*LRIFile, error) {

	_, err := reader.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	file := &LRIFile{
		reader: reader,
	}

	block, err := NewBlock(reader)

	if err != nil {
		return nil, err
	}

	if block.BlockType == BType_LightHeader {
		header := &ltpb.LightHeader{}
		err = proto.Unmarshal(block.Message, header)
		if err != nil {
			return nil, err
		}

		for _, module := range header.GetModules() {
			surface := module.GetSensorDataSurface()

			format := surface.GetFormat()
			if format == ltpb.CameraModule_Surface_RAW_BAYER_JPEG {

				data := block.Body[*surface.DataOffset-32:]

				image, err := NewRawImage(module, surface, data)
				if err != nil {
					return nil, err
				}
				file.Images = append(file.Images, *image)
			}
		}

	}
	return file, nil

}
