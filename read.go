package main

import (
	"bufio"
	"bytes"
	"log"
	"os"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"github.com/t3ta/jvlink-go/jv"
)

func read() {
	f, err := os.OpenFile("./data/RAVM2022069920220701134618.jvd.dat", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		jvrr := jv.NewJvRaRace()
		stream := kaitai.NewStream(bytes.NewReader(scanner.Bytes()))
		err = jvrr.Read(stream, nil, jvrr)
		if err != nil {
			log.Fatal(err)
		}
		print(jvrr.Records.RaceId)
	}
}
