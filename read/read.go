package read

import (
	"bufio"
	"bytes"
	"log"
	"os"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
	"github.com/t3ta/jvlink-go/jv"
)

func main() {
	f, err := os.OpenFile("../data/RACE.dat", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		jvrr := jv.NewJvRaRace()
		stream := kaitai.NewStream(bytes.NewReader(scanner.Bytes()))
		jvrr.Read(stream, nil, jvrr)
		print(jvrr.Records.RaceId)
	}
}
