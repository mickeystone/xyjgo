package mp4

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
)

// Track Header Box (tkhd - mandatory)
//
// Status : only version 0 is decoded. version 1 is not supported
//
// This box describes the track. Duration is measured in time units (according to the time scale
// defined in the movie header box).
//
// Volume (relevant for audio tracks) is a fixed point number (8 bits + 8 bits). Full volume is 1.0.
// Width and Height (relevant for video tracks) are fixed point numbers (16 bits + 16 bits).
// Video pixels are not necessarily square.

// https://blog.csdn.net/PirateLeo/article/details/7597892
/*
每个trak都包含了一个track header atom. Thetrack header atom 定义了一个track的特性，例如时间，空间和音量信息，它的类型是('tkhd').
*/
type TkhdBox struct {
	Version          byte
	Flags            [3]byte
	CreationTime     uint32
	ModificationTime uint32
	TrackId          uint32
	reserved1        uint32
	Duration         uint32
	reserved2        [2]uint32
	Layer            uint16 // 类似photoshop中图层的概念，数值小的在播放时更贴近用户（上层图层）
	AlternateGroup   uint16 // should be int16
	Volume           Fixed16
	reserved3        uint16
	Matrix           []byte // 36
	Width, Height    Fixed32
}

func DecodeTkhd(r io.Reader) (Box, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	t := &TkhdBox{
		Version:          data[0],
		Flags:            [3]byte{data[1], data[2], data[3]},
		CreationTime:     binary.BigEndian.Uint32(data[4:8]),
		ModificationTime: binary.BigEndian.Uint32(data[8:12]),
		TrackId:          binary.BigEndian.Uint32(data[12:16]),
		// reserved1   16:20
		Duration: binary.BigEndian.Uint32(data[20:24]),
		// reserved2	24:32
		Layer:          binary.BigEndian.Uint16(data[32:34]),
		AlternateGroup: binary.BigEndian.Uint16(data[34:36]),
		Volume:         fixed16(data[36:38]),
		// reserved3	38:40
		Matrix: data[40:76],
		Width:  fixed32(data[76:80]),
		Height: fixed32(data[80:84]),
	}

	if t.Version == 1 {
		return nil, ErrOlnyDecodedVer0
	}
	return t, nil
}

func (b *TkhdBox) Type() string {
	return "tkhd"
}

func (b *TkhdBox) Size() int {
	return BoxHeaderSize + 84
}

func (b *TkhdBox) Encode(w io.Writer) error {
	err := EncodeHeader(b, w)
	if err != nil {
		return err
	}
	buf := makebuf(b)
	buf[0] = b.Version
	buf[1], buf[2], buf[3] = b.Flags[0], b.Flags[1], b.Flags[2]
	binary.BigEndian.PutUint32(buf[4:], b.CreationTime)
	binary.BigEndian.PutUint32(buf[8:], b.ModificationTime)
	binary.BigEndian.PutUint32(buf[12:], b.TrackId)
	binary.BigEndian.PutUint32(buf[20:], b.Duration)
	binary.BigEndian.PutUint16(buf[32:], b.Layer)
	binary.BigEndian.PutUint16(buf[34:], b.AlternateGroup)
	putFixed16(buf[36:], b.Volume)
	copy(buf[40:], b.Matrix)
	putFixed32(buf[76:], b.Width)
	putFixed32(buf[80:], b.Height)
	_, err = w.Write(buf)
	return err
}

func (b *TkhdBox) Dump() {
	fmt.Println("Track Header:")
	fmt.Printf(" Duration: %d units\n WxH: %sx%s\n", b.Duration, b.Width, b.Height)
}
