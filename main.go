package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

const LINE_SIZE = 102890

func main() {
	logFile, _ := os.OpenFile("./logs/output.log", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	log.SetOutput(io.Writer(logFile))

	dispatch := JVOpen()
	defer JVClose(dispatch)

	var (
		file *os.File
		buf  *bufio.Writer
	)

	newFile := true

	for i := 0; true; i++ {
		var (
			filename string
			line     string
		)

		if newFile {
			_, err := oleutil.CallMethod(dispatch, "JVRead", &line, 102890, &filename)
			if err != nil {
				log.Fatal(err)
			}
			filepath := fmt.Sprintf("%s/rawdata/%s.dat", os.Getenv("JVDataPath"), filename)

			if !(strings.HasPrefix(filename, "RA") || strings.HasPrefix(filename, "SE")) {
				oleutil.MustCallMethod(dispatch, "JVSkip")
				log.Printf("Skipped %s", filename)
				continue
			}

			if fileExists(filepath) {
				oleutil.MustCallMethod(dispatch, "JVSkip")
				log.Printf("Skipped %s", filename)
				continue
			}

			file = fileOpen(filepath)
			if err != nil {
				log.Fatal(err)
			}
			buf = bufio.NewWriter(file)
			defer file.Close()
			log.Printf("Open %s", filename)

			_, err = buf.Write([]byte(line))
			if err != nil {
				log.Fatal(err)
			}

			newFile = false
			continue
		}

		res, err := oleutil.CallMethod(dispatch, "JVRead", &line, 102890, &filename)
		if err != nil {
			log.Fatal(err)
		}

		status := res.Value().(int32)
		if status == 0 {
			log.Print("Completed")
			break
		}
		if status == -1 {
			log.Printf("Get %s finished.", filename)
			buf.Flush()
			newFile = true
			continue
		}

		_, err = buf.Write([]byte(line))
		if err != nil {
			log.Fatal(err)
		}
	}
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

	res = oleutil.MustCallMethod(dispatch, "JVOpen", "RACE", os.Getenv("JVLastUpdate"), 1, 0, 0, "")
	code = int(res.Val)
	if code != 0 {
		log.Fatalf("JVOpen failed with code: %d", code)
	}

	log.Print("JVOpen succeed")

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
