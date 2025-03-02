package peer



// These are cmds that are called by the peer request handler

func (peer Peer) Connect(req *Request) {


}



func (peer Peer) MetadataShare(req *Request) {

	// this handles a metadata share request (when another node is sharing metadata for it's files)


	// decode the bytes into a list of metadata

	// we need to be able to read current metadata for this peer and then add the 

}


func (peer Peer) FileChunk(req *Request) {


}


func (peer Peer) Ping(req *Request) {


}


func (peer Peer) Disconnect(req *Request) {


}


func (peer Peer) BadRequest(req *Request) {


}