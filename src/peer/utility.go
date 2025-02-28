package peer
// contains utility functions (metadata generation, chunking, reassembly, validation)


import (
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
)


type Metadata struct {
	FileName  string `json:"file_name"`
	FileSize  uint64 `json:"file_size"`
	ChunkSize uint64 `json:"chunk_size"`
}


func GenerateMetadata(filename string, metadata_path string) {

	// metadata for a particular file should inlude filename, file size, chunk size 

	// build the metadata into a json object; write the object to the metadata file


	fileInfo, err := os.Stat(filename)

	if err != nil {
		log.Fatal(err)
	}


	// Size of the file in bytes
	fileSize := fileInfo.Size()


	// build the struct

	metadata := new(Metadata)

	metadata.FileName  = filename
	metadata.FileSize  = uint64(fileSize)
	metadata.ChunkSize = CHUNK_SIZE

	
	metadataFile, err := os.OpenFile(metadata_path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644) // 0644 sets permissions (read/write for owner; read for everyone else)

	if err != nil {
		log.Println("Error opening metadata file: %v",err)
		return
	}

	defer metadataFile.Close()
	

	byteValue, err := ioutil.ReadFile(metadata_path)

	if err != nil {
		log.Println("Error reading metadata file: %v",  err)
		return
	}
	


	data := new([]Metadata)
	err = json.Unmarshal(byteValue, &data)

	if err != nil {

		if err.Error() == "unexpected end of JSON input" {
			// This means the file is empty so instead we will just write
			// append to the data
			*data = append(*data, *metadata)

			// encode the list
			encodedData, err := json.Marshal(data)

			if err != nil {
				log.Println("Error encoding updated metadata: %v", err)
			}

			// rewrite the data to the metadata json file
			_, err = metadataFile.Write(encodedData)

		} else {
			log.Printf("Error decoding metadata file for %s: %v , error code: %v", filename, err, err.Error() )
		}
		return
	}


	// append to the data
	*data = append(*data, *metadata)

	// encode the list
	encodedData, err := json.Marshal(data)

	if err != nil {
		log.Println("Error encoding updated metadata: %v", err)
	}

	// rewrite the data to the metadata json file
	_, err = metadataFile.Write(encodedData)

	if err != nil {
		log.Println("Error writing to metadata file: %v", err)
	}

	log.Printf("%s metadata successfully generated.")

}



func (peer *Peer) Chunk(filename string) {
	// takes a filename as input and modifies the DHT for this peer
}