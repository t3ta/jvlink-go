package main

// #cgo LDFLAGS: -largeaddressaware

import (
	"bufio"
	"log"
	"os"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

const LINE_SIZE = 102890

//go:linkname safeArrayCreate github.com/go-ole/go-ole.safeArrayCreate
func safeArrayCreate(variantType ole.VT, dimensions uint32, bounds *ole.SafeArrayBound) (*ole.SafeArray, error)

func main() {
	dispatch := JVOpen()
	defer JVClose(dispatch)
	bound := &ole.SafeArrayBound{Elements: 102890, LowerBound: 0}
	safeArray, _ := safeArrayCreate(ole.VT_UI1, 1, bound)
	variant := ole.NewVariant(ole.VT_ARRAY|ole.VT_UI1, int64(uintptr(unsafe.Pointer(safeArray))))

	f, err := os.OpenFile("./data/RACE.dat", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := bufio.NewWriter(f)

	for i := 0; true; i++ {
		var filename string
		res := oleutil.MustCallMethod(dispatch, "JVGets", &variant, LINE_SIZE, &filename)

		switch res.Value().(int32) {
		case 0:
			log.Printf("Completed")
		case -1:
			log.Printf("Get %s finished.", filename)
		default:
			bytes := variant.ToArray().ToByteArray()
			_, err = buf.Write(bytes)
			if err != nil {
				log.Fatal(err)
			}
		}

		// if str[0:2] == "RA" {
		// 	f.Write(byteArray)
		// 	log.Print(byteArray)
		// 	jvrr := jv.NewJvRaRace()
		// 	reader := bytes.NewReader(byteArray)
		// 	stream := kaitai.NewStream(reader)

		// 	jvrr.Read(stream, nil, jvrr)

		// 	print(jvrr.Records.Hondai)
		// }
	}
	buf.Flush()
}

// var mem runtime.MemStats

// func PrintMemory() {
// 	runtime.ReadMemStats(&mem)
// 	fmt.Println(mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
// }

func JVGets(disp *ole.IDispatch, variant *ole.VARIANT) bool {
	res, err := oleutil.CallMethod(disp, "JVGets", variant, 102890, "")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Clear()
	return res.Val == 0
}

func JVOpen() (dispatch *ole.IDispatch) {
	ole.CoInitialize(0)

	unknown, err := oleutil.CreateObject("JVDTLab.JVLink")
	if err != nil {
		log.Fatal(err)
	}

	dispatch = unknown.MustQueryInterface(ole.IID_IDispatch)
	res := oleutil.MustCallMethod(dispatch, "JVInit", "JVLinkGo")
	code := int(res.Val)
	if code != 0 {
		log.Fatalf("JVInit failed with code: %d", code)
	}

	oleutil.MustCallMethod(dispatch, "JVSetSavePath", "C:/Keiba/jvdata")
	oleutil.MustCallMethod(dispatch, "JVSetSaveFlag", 1)

	res = oleutil.MustCallMethod(dispatch, "JVOpen", "RACE", "20000101000000", 4, 0, 0, "")
	code = int(res.Val)
	if code != 0 {
		log.Fatalf("JVOpen failed with code: %d", code)
	}

	return dispatch
}

func JVClose(dispatch *ole.IDispatch) {
	oleutil.MustCallMethod(dispatch, "JVClose")
	ole.CoUninitialize()
}
