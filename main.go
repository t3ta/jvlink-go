package main

import (
	"log"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

//go:linkname safeArrayCreate github.com/go-ole/go-ole.safeArrayCreate
func safeArrayCreate(variantType ole.VT, dimensions uint32, bounds *ole.SafeArrayBound) (*ole.SafeArray, error)

func main() {
	ole.CoInitialize(0)
	defer ole.CoUninitialize()

	unknown, err := oleutil.CreateObject("JVDTLab.JVLink")
	if err != nil {
		log.Fatal(err)
		return
	}

	agent := unknown.MustQueryInterface(ole.IID_IDispatch)
	res := oleutil.MustCallMethod(agent, "JVInit", "JVLinkGo")
	code := int(res.Val)
	if code != 0 {
		log.Fatalf("JVInit failed with code: %d", code)
	}

	res = oleutil.MustCallMethod(agent, "JVOpen", "RACE", "20220801000000", 2, 0, 0, "")
	code = int(res.Val)
	if code != 0 {
		log.Fatalf("JVOpen failed with code: %d", code)
	}
	defer oleutil.MustCallMethod(agent, "JVClose")

	var buff [102890]byte
	bound := &ole.SafeArrayBound{Elements: uint32(len(buff)), LowerBound: 0}
	safeArray, _ := safeArrayCreate(ole.VT_UI1, 1, bound)
	variant := ole.NewVariant(ole.VT_ARRAY|ole.VT_UI1, int64(uintptr(unsafe.Pointer(safeArray))))

	_, err = oleutil.CallMethod(agent, "JVGets", &variant, 102890, "")
	if err != nil {
		log.Fatal(err)
	}

	byteArray := variant.ToArray().ToByteArray()
	// encoder := japanese.ShiftJIS.NewEncoder()
	// sjisStr, _, err := transform.Bytes(encoder, byteArray)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	print(string(byteArray))
}
