package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

//定义协议 数据包的前四个字节作为包头 存储数据长度

//将待发送的消息编码
func Encode(msg string) ([]byte, error) {
	//读取消息长度 转换成四个字节长度 int32
	var length = int32(len(msg))

	var pkg = new(bytes.Buffer)
	//消息头
	//writer  字节序（大端、小端） 数据（长度）
	err := binary.Write(pkg,binary.LittleEndian, length)
	if err!=nil {
		return nil,err
	}  //封装失败
	//写入消息体
	err = binary.Write(pkg,binary.LittleEndian, []byte(msg))
	if err!=nil {
		return nil,err
	}
	return pkg.Bytes(),nil
}


//消息解码
func Decode(reader *bufio.Reader) (string,error) {
	//首先获取消息长度
	lengthByte,_ := reader.Peek(4)  //获取前四个
	lengthBuff := bytes.NewBuffer(lengthByte)
}