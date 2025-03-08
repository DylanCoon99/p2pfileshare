package peer
// contains utility functions (metadata generation, chunking, reassembly, validation)


import (
	"os"
	"log"
	//"hash"
	"io/ioutil"
	"path/filepath"
	"encoding/json"
)


type Metadata struct {
	FileName  string `json:"file_name"`
	FileSize  uint64 `json:"file_size"`
	ChunkSize uint64 `json:"chunk_size"`
}


func (cfg *PeerCfg) GenerateMetadata() {

	// metadata for a particular file should inlude filename, file size, chunk size 

	// build the metadata into a json object; write the object to the metadata file

	dirPath := cfg.DirectoryPath
	fileList := new([]string)

	data := new([]Metadata)

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		
		if err != nil {
			return err
		}

		if !info.IsDir() {
			// add file to the list
			*fileList = append(*fileList, path)
		}
		return nil
	})


	if err != nil {
		log.Printf("Error obtaining list of files: %v", err)
	}

	

	for _, file := range *fileList {
		fileInfo, err := os.Stat(file)

		if err != nil {
			log.Fatal(err)
		}

		// Size of the file in bytes
		fileSize := fileInfo.Size()


		// build the struct

		metadata := new(Metadata)

		metadata.FileName  = file
		metadata.FileSize  = uint64(fileSize)
		metadata.ChunkSize = CHUNK_SIZE

		*data = append(*data, *metadata)

		cfg.Metadata.Add(metadata)  // add metadata to the metadata set
	}
	

	log.Printf("Here is the metadata list: %v", data)



	err = cfg.WriteMetadata(*data) // write metadata to the file

	if err != nil {
		log.Printf("Error writing to metadata file: %v", err)
	}
	
}



func (cfg *PeerCfg) WriteMetadata(metadataList []Metadata) error {


	metadataFile, err := os.OpenFile(cfg.MetadataPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644) // 0644 sets permissions (read/write for owner; read for everyone else)

	if err != nil {
		log.Println("Error opening metadata file: %v",err)
		return err
	}

	defer metadataFile.Close()
	

	// encode the list
	encodedData, err := json.Marshal(metadataList)

	if err != nil {
		log.Printf("Error encoding updated metadata: %v", err)
		return err
	}

	// rewrite the data to the metadata json file
	_, err = metadataFile.Write(encodedData)

	if err != nil {
		log.Printf("Error writing to metadata file: %v", err)
		return err
	}

	log.Printf("Metadata successfully generated.")
	return nil

}



func ExtractMetadata(metadataPath string) (*[]Metadata, error) {

	// extract the metadata from the json file into a Metadata array


	// Metadata array
	metadata := new([]Metadata)


	// Open metadata file
	metadataFile, err := os.OpenFile(metadataPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644) // 0644 sets permissions (read/write for owner; read for everyone else)

	if err != nil {
		log.Printf("Error opening metadata file: %v",err)
		return nil, err
	}

	defer metadataFile.Close()
	
	// read metadata file to bytes
	byteValue, err := ioutil.ReadFile(metadataPath)

	if err != nil {
		log.Printf("Error reading metadata file: %v",  err)
		return nil, err
	}

	// decode bytes to a struct
	err = json.Unmarshal(byteValue, metadata)
	
	if err != nil {

		if err.Error() == "unexpected end of JSON input" {
			// This means the file is empty so instead we will just return empty array
			return metadata, nil

		} else {
			log.Printf("Error decoding metadata file: %v , error code: %v",  err, err.Error() )
		}

		return nil, err
	}


	return metadata, nil


}



func (cfg *PeerCfg) ConstructMetadata(addMetadata []Metadata) error {


	metadataList, err := ExtractMetadata(cfg.MetadataPath)

	if err != nil {
		log.Printf("Error extracting metadata: %v", err)
		return err
	}


	for _, metadata := range addMetadata {
		// for each element in addMetadata, check if it exists in metadataList
		if cfg.Metadata.Exists(metadata.FileName) {
			// metadata already in list; do not add
		} else {
			cfg.Metadata.Add(&metadata)
			*metadataList = append(*metadataList, metadata)
		}
	}


	// Write metadata list to the metadata file

	err = cfg.WriteMetadata(*metadataList)

	if err != nil {
		log.Printf("Error writing metadata: %v", err)
		return err
	}

	return nil

}



func (cfg *PeerCfg) GenerateDHT() {

	/*
	metadataPath := cfg.MetadataPath

	// obtain array of metadata
	metadata, err := ExtractMetadata(metadataPath)

	if err != nil {
		// handle error
	}

	// apply the hash function to each metadata in the array

	// store it in the hash table

	*/
}




func (peer *Peer) Chunk(filename string) {
	// takes a filename as input and modifies the DHT for this peer
}




func CreateMetadataSet() *MetadataSet {

	return &MetadataSet{metadata: make(map[string]Metadata)}

}



func (m *MetadataSet) Add(item *Metadata) {
	m.metadata[item.FileName] = *item
}



func (m *MetadataSet) Exists(filename string) bool {

	_, ok := m.metadata[filename]

	return ok
}

