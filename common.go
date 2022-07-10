package main

import (
	"encoding/binary"
	"fmt"
)
const (
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
	_seqOffset    = _opOffset + _opSize
	protocolVersion = 2
	operation = 4
	sequence        = 9
)

func encode(body string) []byte {
	var (
		packLen = _rawHeaderSize + uint32(len(body))
		buf = make([]byte, packLen)
	)
	binary.BigEndian.PutUint32(buf[_packOffset:], packLen)
	binary.BigEndian.PutUint16(buf[_headerOffset:], uint16(_rawHeaderSize))
	binary.BigEndian.PutUint16(buf[_verOffset:], uint16(protocolVersion))
	binary.BigEndian.PutUint32(buf[_opOffset:], operation)
	binary.BigEndian.PutUint32(buf[_seqOffset:], sequence)
	byteBody := []byte(body)
	copy(buf[_rawHeaderSize:], byteBody)
	return buf
}


func decode(data []byte) {
	packetLen := binary.BigEndian.Uint32(data[_packOffset:_headerOffset])
	headerLen := binary.BigEndian.Uint16(data[_headerOffset:_verOffset])
	version := binary.BigEndian.Uint16(data[_verOffset:_opOffset])
	operation := binary.BigEndian.Uint32(data[_opOffset:_seqOffset])
	sequence := binary.BigEndian.Uint32(data[_seqOffset:_rawHeaderSize])
	body := string(data[_rawHeaderSize:])

	//输出
	fmt.Printf("packet length:%v\n", packetLen)
	fmt.Printf("header length:%v\n", headerLen)
	fmt.Printf("version:%v\n", version)
	fmt.Printf("operation:%v\n", operation)
	fmt.Printf("sequence:%v\n", sequence)
	fmt.Printf("body:%v\n", body)
}