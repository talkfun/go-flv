//
// Copyright (c) 2018- yutopp (yutopp@gmail.com)
//
// Distributed under the Boost Software License, Version 1.0. (See accompanying
// file LICENSE_1_0.txt or copy at  https://www.boost.org/LICENSE_1_0.txt)
//

package tag

import (
	"io"
	"io/ioutil"
)

// ========================================
// FLV tags

type TagType uint8

const (
	TagTypeAudio      TagType = 8
	TagTypeVideo      TagType = 9
	TagTypeScriptData TagType = 18
)

type FlvTag struct {
	TagType
	DataSize          uint32
	Timestamp         uint32
	TimestampExtended uint8
	StreamID          uint32      // 24bit
	Data              interface{} // *AudioData | *VideoData | *ScriptData
}

// Close
func (t *FlvTag) Close() {
	// TODO: wrap an error?
	switch data := t.Data.(type) {
	case *AudioData:
		data.Close()
	case *VideoData:
		data.Close()
	}
}

// ========================================
// Audio tags

type SoundFormat uint8

const (
	SoundFormatLinearPCMPlatformEndian SoundFormat = 0
	SoundFormatADPCM                               = 1
	SoundFormatMP3                                 = 2
	SoundFormatLinearPCMLittleEndian               = 3
	SoundFormatNellymoser16kHzMono                 = 4
	SoundFormatNellymoser8kHzMono                  = 5
	SoundFormatNellymoser                          = 6
	SoundFormatG711ALawLogarithmicPCM              = 7
	SoundFormatG711muLawLogarithmicPCM             = 8
	SoundFormatReserved                            = 9
	SoundFormatAAC                                 = 10
	SoundFormatSpeex                               = 11
	SoundFormatMP3_8kHz                            = 14
	SoundFormatDeviceSpecificSound                 = 15
)

var SoundFormatName = map[SoundFormat]string{
	SoundFormatLinearPCMPlatformEndian: "LinearPCMPlatformEndian",
	SoundFormatADPCM:                   "ADPCM",
	SoundFormatMP3:                     "MP3",
	SoundFormatLinearPCMLittleEndian:   "LinearPCMLittleEndian",
	SoundFormatNellymoser16kHzMono:     "Nellymoser16kHzMono",
	SoundFormatNellymoser8kHzMono:      "Nellymoser8kHzMono",
	SoundFormatNellymoser:              "Nellymoser",
	SoundFormatG711ALawLogarithmicPCM:  "G711ALawLogarithmicPCM",
	SoundFormatG711muLawLogarithmicPCM: "G711muLawLogarithmicPCM",
	SoundFormatReserved:                "Reserved",
	SoundFormatAAC:                     "AAC",
	SoundFormatSpeex:                   "Speex",
	SoundFormatMP3_8kHz:                "MP3_8kHz",
	SoundFormatDeviceSpecificSound:     "DeviceSpecificSound",
}

type SoundRate uint8

const (
	SoundRate5_5kHz SoundRate = 0
	SoundRate11kHz            = 1
	SoundRate22kHz            = 2
	SoundRate44kHz            = 3
)

var SoundRateName = map[SoundRate]string{
	SoundRate5_5kHz: "5_5kHz",
	SoundRate11kHz:  "11kHz",
	SoundRate22kHz:  "22kHz",
	SoundRate44kHz:  "44kHz",
}

type SoundSize uint8

const (
	SoundSize8Bit  SoundSize = 0
	SoundSize16Bit           = 1
)

var SoundSizeName = map[SoundSize]string{
	SoundSize8Bit:  "8Bit",
	SoundSize16Bit: "16Bit",
}

type SoundType uint8

const (
	SoundTypeMono   SoundType = 0
	SoundTypeStereo           = 1
)

var SoundTypeName = map[SoundType]string{
	SoundTypeMono:   "Mono",
	SoundTypeStereo: "Stereo",
}

type AudioData struct {
	SoundFormat   SoundFormat
	SoundRate     SoundRate
	SoundSize     SoundSize
	SoundType     SoundType
	AACPacketType AACPacketType
	Data          io.Reader
}

func (d *AudioData) SoundFormatName() string {
	return SoundFormatName[d.SoundFormat]
}

func (d *AudioData) SoundRateName() string {
	return SoundRateName[d.SoundRate]
}

func (d *AudioData) SoundSizeName() string {
	return SoundSizeName[d.SoundSize]
}

func (d *AudioData) SoundTypeName() string {
	return SoundTypeName[d.SoundType]
}

func (d *AudioData) AACPacketTypeName() string {
	return AACPacketTypeName[d.AACPacketType]
}

func (d *AudioData) Read(buf []byte) (int, error) {
	return d.Read(buf)
}

func (d *AudioData) Close() {
	_, _ = io.Copy(ioutil.Discard, d.Data) //  // TODO: wrap an error?
}

type AACPacketType uint8

const (
	AACPacketTypeSequenceHeader AACPacketType = 0
	AACPacketTypeRaw                          = 1
)

var AACPacketTypeName = map[AACPacketType]string{
	AACPacketTypeSequenceHeader: "SequenceHeader",
	AACPacketTypeRaw:            "Raw",
}

type AACAudioData struct {
	AACPacketType AACPacketType
	Data          io.Reader
}

// ========================================
// Video Tags

type FrameType uint8

const (
	FrameTypeKeyFrame              FrameType = 1
	FrameTypeInterFrame                      = 2
	FrameTypeDisposableInterFrame            = 3
	FrameTypeGeneratedKeyFrame               = 4
	FrameTypeVideoInfoCommandFrame           = 5
)

type CodecID uint8

const (
	CodecIDJPEG                   CodecID = 1
	CodecIDSorensonH263                   = 2
	CodecIDScreenVideo                    = 3
	CodecIDOn2VP6                         = 4
	CodecIDOn2VP6WithAlphaChannel         = 5
	CodecIDScreenVideoVersion2            = 6
	CodecIDAVC                            = 7
)

var (
	FrameTypeName = map[FrameType]string{
		FrameTypeKeyFrame:              "KeyFrame",
		FrameTypeInterFrame:            "InterFrame",
		FrameTypeDisposableInterFrame:  "DisposableInterFrame",
		FrameTypeGeneratedKeyFrame:     "GeneratedKeyFrame",
		FrameTypeVideoInfoCommandFrame: "VideoInfoCommandFrame",
	}

	CodecIDName = map[CodecID]string{
		CodecIDJPEG:                   "JPEG",
		CodecIDSorensonH263:           "SorensonH263",
		CodecIDScreenVideo:            "ScreenVideo",
		CodecIDOn2VP6:                 "On2VP6",
		CodecIDOn2VP6WithAlphaChannel: "On2VP6WithAlphaChannel",
		CodecIDScreenVideoVersion2:    "ScreenVideoVersion2",
		CodecIDAVC:                    "AVC",
	}
)

type VideoData struct {
	FrameType       FrameType
	CodecID         CodecID
	AVCPacketType   AVCPacketType
	CompositionTime int32
	Data            io.Reader
}

func (d *VideoData) FrameTypeName() string {
	return FrameTypeName[d.FrameType]
}

func (d *VideoData) CodecIDName() string {
	return CodecIDName[d.CodecID]
}

func (d *VideoData) AVCPacketTypeName() string {
	return AVCPacketTypeName[d.AVCPacketType]
}

func (d *VideoData) Read(buf []byte) (int, error) {
	return d.Read(buf)
}

func (d *VideoData) Close() {
	_, _ = io.Copy(ioutil.Discard, d.Data) //  // TODO: wrap an error?
}

type AVCPacketType uint8

const (
	AVCPacketTypeSequenceHeader AVCPacketType = 0
	AVCPacketTypeNALU                         = 1
	AVCPacketTypeEOS                          = 2
)

var AVCPacketTypeName = map[AVCPacketType]string{
	AVCPacketTypeSequenceHeader: "SequenceHeader",
	AVCPacketTypeNALU:           "NALU",
	AVCPacketTypeEOS:            "EOS",
}

type AVCVideoPacket struct {
	AVCPacketType   AVCPacketType
	CompositionTime int32
	Data            io.Reader
}

type AVCNaluType uint8

const (
	AVCNaluTypeSLICE    AVCNaluType = 1
	AVCNaluTypeDPA                  = 2
	AVCNaluTypeDPB                  = 3
	AVCNaluTypeDPC                  = 4
	AVCNaluTypeIDR                  = 5
	AVCNaluTypeSEI                  = 6
	AVCNaluTypeSPS                  = 7
	AVCNaluTypePPS                  = 8
	AVCNaluTypeAUD                  = 9
	AVCNaluTypeEOSEQ                = 10
	AVCNaluTypeEOSTREAM             = 11
	AVCNaluTypeFILL                 = 12
)

var AVCNaluTypeName = map[AVCNaluType]string{
	AVCNaluTypeSLICE:    "SLICE",
	AVCNaluTypeDPA:      "DPA",
	AVCNaluTypeDPB:      "DPB",
	AVCNaluTypeDPC:      "DPC",
	AVCNaluTypeIDR:      "IDR",
	AVCNaluTypeSEI:      "SEI",
	AVCNaluTypeSPS:      "SPS",
	AVCNaluTypePPS:      "PPS",
	AVCNaluTypeAUD:      "AUD",
	AVCNaluTypeEOSEQ:    "EOSEQ",
	AVCNaluTypeEOSTREAM: "EOSTREAM",
	AVCNaluTypeFILL:     "FILL",
}

func GetAVCNaluTypeName(naluType AVCNaluType) string {
	return AVCNaluTypeName[naluType]
}

// ========================================
// Data tags

type ScriptData struct {
	// all values are represented as subset of AMF0
	Objects map[string]interface{}
}
