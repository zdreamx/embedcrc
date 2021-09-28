package embedcrc

//CRC16Sum 通用嵌入式的CRC16函数
func CRC16Sum(data []byte) uint16 {
	var reg_crc uint16 = 0xFFFF
	blen := len(data)
	for i := 0; i < blen; i++ {
		reg_crc ^= uint16(data[i])
		for j := 0; j < 8; j++ {
			if (reg_crc & 0x01) > 0 {
				reg_crc = (reg_crc >> 1) ^ 0xA001
			} else {
				reg_crc = reg_crc >> 1
			}

		}
	}
	return reg_crc
}
