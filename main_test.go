package main

import (
	"testing"
	"unsafe"

	"github.com/go-ole/go-ole"
)

func BenchmarkJVGets(b *testing.B) {
	dispatch := JVOpen()
	defer JVClose(dispatch)
	bound := &ole.SafeArrayBound{Elements: 102890, LowerBound: 0}
	safeArray, _ := safeArrayCreate(ole.VT_UI1, 1, bound)
	variant := ole.NewVariant(ole.VT_ARRAY|ole.VT_UI1, int64(uintptr(unsafe.Pointer(safeArray))))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JVGets(dispatch, &variant)
	}
	b.StopTimer()
}

// func BenchmarkWriteByteArray(b *testing.B) {
// 	f, err := os.Open("./data/test.dat")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
// 	bytes := make([]byte, 102890)
// 	buf := bufio.NewWriter(f)

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		buf.Write(bytes)
// 	}
// 	buf.Flush()
// 	b.StopTimer()
// }
