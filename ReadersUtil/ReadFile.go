package ReadersUtil

import (
	"log"
	"os"
)

func ReadFile() {

	file, err := os.OpenFile("sloth.txt")

	if err != nil {
		log.Fatalf("error occured while loading the file : %v", err)
	}

	defer file.Close()

	bytesRead := make([]bytes, 33)

	n, err := file.Read(bytesRead)

	if err != nil {
		log.Fatalf("Error opening File : %v", err)
	}

	log.Printf("We read \"%s\" into bytesRead (%d bytes)", string(bytesRead), n)

}
