package main

import (
	"log"
	"os"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

//go:linkname safeArrayCreate github.com/go-ole/go-ole.safeArrayCreate
func safeArrayCreate(variantType ole.VT, dimensions uint32, bounds *ole.SafeArrayBound) (*ole.SafeArray, error)

//go:linkname safeArrayDestroyData github.com/go-ole/go-ole.safeArrayDestroyData
func safeArrayDestroyData(safearray *ole.SafeArray) error

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

	res = oleutil.MustCallMethod(agent, "JVOpen", "RACE", "20200801000000", 4, 0, 0, "")
	code = int(res.Val)
	if code != 0 {
		log.Fatalf("JVOpen failed with code: %d", code)
	}
	defer oleutil.MustCallMethod(agent, "JVClose")

	// bound := &ole.SafeArrayBound{Elements: 102890, LowerBound: 0}
	// safeArray, _ := safeArrayCreate(ole.VT_UI1, 1, bound)
	// variant := ole.NewVariant(ole.VT_ARRAY|ole.VT_UI1, int64(uintptr(unsafe.Pointer(safeArray))))
	// variant.Clear()

	f, err := os.OpenFile("./data/test.dat", os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for {
		bound := &ole.SafeArrayBound{Elements: 102890, LowerBound: 0}
		safeArray, _ := safeArrayCreate(ole.VT_UI1, 1, bound)
		variant := ole.NewVariant(ole.VT_ARRAY|ole.VT_UI1, int64(uintptr(unsafe.Pointer(safeArray))))
		defer safeArrayDestroyData(safeArray)
		defer variant.Clear()

		res, err = oleutil.CallMethod(agent, "JVGets", &variant, 102890, "")
		if err != nil {
			log.Fatal(err)
		}

		if res.Val == 0 {
			log.Print("Completed")
			return
		}

		byteArray := variant.ToArray().ToByteArray()
		str := string(byteArray)
		if str[0:2] == "RA" {
			f.Write(byteArray)

			// jvrr := jv.NewJvRaRace()
			// reader := bytes.NewReader(byteArray)
			// stream := kaitai.NewStream(reader)

			// jvrr.Read(stream, nil, jvrr)

			// print(jvrr.Records.Hondai)
		}
	}
}
