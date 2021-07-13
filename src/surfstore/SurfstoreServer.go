package surfstore

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {
	BlockStore BlockStoreInterface
	MetaStore  MetaStoreInterface
}

func (s *Server) GetFileInfoMap(succ *bool, serverFileInfoMap *map[string]FileMetaData) error {
	//panic("todo")
	//The server should just populate serverFileInfoMap with the entries it has in its MetaStore.
	return s.MetaStore.GetFileInfoMap(succ,serverFileInfoMap)
}

func (s *Server) UpdateFile(fileMetaData *FileMetaData, latestVersion *int) error {
	//panic("todo")
	return s.MetaStore.UpdateFile(fileMetaData,latestVersion)
}

func (s *Server) GetBlock(blockHash string, blockData *Block) error {
	//panic("todo")
	return s.BlockStore.GetBlock(blockHash,blockData)
}

func (s *Server) PutBlock(blockData Block, succ *bool) error {
	//panic("todo")
	return s.BlockStore.PutBlock(blockData,succ)
}

func (s *Server) HasBlocks(blockHashesIn []string, blockHashesOut *[]string) error {
	//panic("todo")
	return s.BlockStore.HasBlocks(blockHashesIn,blockHashesOut)
}

// This line guarantees all method for surfstore are implemented
var _ Surfstore = new(Server)

func NewSurfstoreServer() Server {
	blockStore := BlockStore{BlockMap: map[string]Block{}}
	metaStore := MetaStore{FileMetaMap: map[string]FileMetaData{}}

	return Server{
		BlockStore: &blockStore,
		MetaStore:  &metaStore,
	}
}

func ServeSurfstoreServer(hostAddr string, surfstoreServer Server) error {
	//panic("todo")
	 var s *Server = &surfstoreServer
	// rpc.Register(s)
	rpc.RegisterName("Surfstore",s)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp",hostAddr)
	if e != nil {
		log.Panicln("listen error: ",e)
	}
	return http.Serve(l,nil)

}
