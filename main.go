package main

import (
	"bufio"
	"fmt"
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

	var (
		f   *os.File
		buf *bufio.Writer
	)

	for i := 0; true; i++ { // each file
		var filename string
		res := oleutil.MustCallMethod(dispatch, "JVGets", &variant, LINE_SIZE, &filename)
		filepath := fmt.Sprintf("%s/rawdata/%s.dat", os.Getenv("JVDataPath"), filename)

		if fileExists(filepath) {
			oleutil.MustCallMethod(dispatch, "JVSkip")
			log.Printf("Skipped %s", filename)
			continue
		}

		f = fileOpen(filepath)
		buf = bufio.NewWriter(f)
		defer f.Close()
		log.Printf("Open %s", filename)

		for j := 0; true; j++ { // each line
			switch res.Value().(int32) {
			case 0:
				log.Printf("Completed")
			case -1:
				log.Printf("Get %s finished.", filename)
			default:
				bytes := variant.ToArray().ToByteArray()
				_, err := buf.Write(bytes)
				if err != nil {
					log.Fatal(err)
				}
			}
		}

		buf.Flush()
	}
}

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

	oleutil.MustCallMethod(dispatch, "JVSetSavePath", fmt.Sprintf("%s/jvdata", os.Getenv("JVDataPath")))
	oleutil.MustCallMethod(dispatch, "JVSetSaveFlag", 1)

	res = oleutil.MustCallMethod(dispatch, "JVOpen", "RACE", os.Getenv("JVLastUpdate"), 4, 0, 0, "")
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

func fileOpen(name string) *os.File {
	f, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
