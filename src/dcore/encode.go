/**
* @Author: Jason
* @Date: 2020/9/29 14:31
* @File : encode
* @Software: GoLand
**/

package dcore

import "encoding/binary"

/**
编码消息
*/
func EncodeMsg(msgType uint32, data []byte) []byte {
	msgLen := 8 + len(data)
	msg := make([]byte, msgLen)
	binary.BigEndian.PutUint32(msg, uint32(msgLen))
	binary.BigEndian.PutUint32(msg[4:], msgType)
	copy(msg[8:], data)
	return msg
}

/**
解码消息
*/
func DecodeMsg(data []byte) (uint32, uint32, []byte) {
	msgLen := binary.BigEndian.Uint32(data)
	msgType := binary.BigEndian.Uint32(data[4:])
	return msgType, msgLen, data[8:]
}
