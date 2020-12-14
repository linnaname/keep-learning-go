package main

import (
	"reflect"
	"testing"
	"unsafe"
)

const RUN_CNT = 1000000000

func BenchmarkString2bytesReflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		string2bytesReflect("aaaaaadsaabasfdsafdsafdasfdafdasfdasfdsafdsafdsafdasfdasfdasfdas这段代码是在找到了 key 要插入的位置后，进行“赋值”操作。insertk 和 val 分别表示 key 和 value 所要“放置”的地址。如果 t.indirectkey 为真，说明 bucket 中存储的是 key 的指针，因此需要将 insertk 看成指针的指针，这样才能将 bucket 中的相应位置的值设置成指向真实 key 的地址值，也就是说 key 存放的是指针。")
	}
}

func BenchmarkString2bytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		string2bytes("aaaaaadsaabasfdsafdsafdasfdafdasfdasfdsafdsafdsafdasfdasfdasfdas这段代码是在找到了 key 要插入的位置后，进行“赋值”操作。insertk 和 val 分别表示 key 和 value 所要“放置”的地址。如果 t.indirectkey 为真，说明 bucket 中存储的是 key 的指针，因此需要将 insertk 看成指针的指针，这样才能将 bucket 中的相应位置的值设置成指向真实 key 的地址值，也就是说 key 存放的是指针。")
	}
}

func string2bytesReflect(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bh))
}

func string2bytes(s string) []byte {
	sp := unsafe.Pointer(&s)
	return *(*[]byte)(sp)
}
