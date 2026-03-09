package utils

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"
)

// CreateSn 生成一个序列号.
// 序列号由时间戳(YYMMDDHHMMSS)和一个4位随机数组成.
func CreateSn() string {
	var b [4]byte
	if _, err := rand.Read(b[:]); err != nil {
		panic(err)
	}
	return time.Now().Format("060102150405") + fmt.Sprintf("%04d", binary.BigEndian.Uint32(b[:])%10000)
}

func ConvertIntSliceToInt64Slice(intSlice []int) []int64 {
	int64Slice := make([]int64, len(intSlice))
	for i, v := range intSlice {
		int64Slice[i] = int64(v)
	}
	return int64Slice
}
