# 说明

嵌入式领域通用CRC16校验的golang语言实现

# 用法

```go
import ("github.com/zdreamx/embedcrc")

var buf [10]byte
buf[0]=0x01
buf[1]=0x02
//buf[...]=0x..
crc := embedcrc.CRC16Sum(buf[:8])
```

