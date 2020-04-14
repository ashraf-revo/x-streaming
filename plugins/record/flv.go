package record

import (
	"io"
	"os"
	"path"

	. "github.com/micro-community/x-streaming/engine"
	"github.com/micro-community/x-streaming/engine/avformat"
	"github.com/micro-community/x-streaming/engine/util"
)

func getDuration(file *os.File) uint32 {
	_, err := file.Seek(-4, io.SeekEnd)
	if err == nil {
		var tagSize uint32
		if tagSize, err = util.ReadByteToUint32(file, true); err == nil {
			_, err = file.Seek(-int64(tagSize)-4, io.SeekEnd)
			if err == nil {
				var tag *avformat.AVPacket
				tag, err = avformat.ReadFLVTag(file)
				if err == nil {
					return tag.Timestamp
				}
			}
		}
	}
	return 0
}
func SaveFlv(streamPath string, append bool) error {
	flag := os.O_CREATE
	if append {
		flag = flag | os.O_RDWR | os.O_APPEND
	} else {
		flag = flag | os.O_TRUNC | os.O_WRONLY
	}
	filePath := config.Path + streamPath + ".flv"
	os.MkdirAll(path.Dir(filePath), 0666)
	file, err := os.OpenFile(filePath, flag, 0666)
	if err != nil {
		return err
	}
	p := OutputStream{SendHandler: func(packet *avformat.SendPacket) error {
		return avformat.WriteFLVTag(file, packet)
	}}
	p.ID = filePath
	p.Type = "FlvRecord"
	if append {
		p.OffsetTime = getDuration(file)
		file.Seek(0, io.SeekEnd)
	} else {
		_, err = file.Write(avformat.FLVHeader)
	}
	if err == nil {
		recordings.Store(filePath, &p)
		go func() {
			p.Play(streamPath)
			file.Close()
		}()
	} else {
		file.Close()
	}
	return err
}
