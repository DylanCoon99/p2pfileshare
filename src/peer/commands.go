package peer



// These are cmds that are called by the peer request handler

func (peer Peer) Connect(req *Request) {


}


func (peer Peer) DHTInfo(req *Request) {


}


func (peer Peer) MetadataShare(req *Request) {

	// this handles a metadata share request (when another node is sharing metadata for it's files)
	

}


func (peer Peer) FileChunk(req *Request) {


}


func (peer Peer) Ping(req *Request) {


}


func (peer Peer) Disconnect(req *Request) {


}


func (peer Peer) BadRequest(req *Request) {


}