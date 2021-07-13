# Surfstore

Surfstore is a cloud-based file storage service.
It is a networked file storage application that is based on Dropbox, and lets you sync
files to and from the “cloud”.

Multiple clients can concurrently connect to the SurfStore service (via RPC) to access a common, shared
set of files. Clients accessing SurfStore “see” a consistent set of updates to files, but SurfStore
does not offer any guarantees about operations across files, meaning that it does not support
multi-file transactions (such as atomic move).

The SurfStore service is composed of the following two services:
1. BlockStore : The content of each file in SurfStore is divided up into chunks, or blocks,
each of which has a unique identifier. This service stores these blocks, and when given
an identifier, retrieves and returns the appropriate block.
2. MetadataStore : The MetadataStore service holds the mapping of filenames to blocks.

## Data Types

Recall from the module write-up the following things:

1. The SurfStore service is composed of two services: BlockStore and MetadataStore 
2. A file in SurfStore is broken into an ordered sequence of one or more blocks which are stored in the BlockStore.
3. The MetadataStore maintains the mapping of filenames to hashes of these blocks (and versions) in a map.

The starter code defines the following types for your usage in `SurfstoreInterfaces.go`:

```go
type Block struct {
	BlockData []byte
	BlockSize int
}

type FileMetaData struct {
	Filename      string
	Version       int
	BlockHashList []string
}
```

## Surfstore Interface

`SurfstoreInterfaces.go` also contains interfaces for the BlockStore and the MetadataStore:

```go
type MetaStoreInterface interface {
	// Retrieves the server's FileInfoMap
	GetFileInfoMap(_ignore *bool, serverFileInfoMap *map[string]FileMetaData) error
	
	// Update a file's fileinfo entry
	UpdateFile(fileMetaData *FileMetaData, latestVersion *int) (err error)
}

type BlockStoreInterface interface {

	// Get a block based on its hash
	GetBlock(blockHash string, block *Block) error

	// Put a block
	PutBlock(block Block, succ *bool) error

	// Check if certain blocks are alredy present on the server
	HasBlocks(blockHashesIn []string, blockHashesOut *[]string) error
}
```

The `Surfstore` interface then glues these two together and is also present in `SurfstoreInterfaces.go`.

```go
type Surfstore interface {
	MetaStoreInterface
	BlockStoreInterface
}
```


## Setup

You will need to setup your runtime environment variables so that you can build
your code and also use the executables that will be generated.

1. If you are using a Mac, open `~/.bash_profile` or if you are using a
unix/linux machine, open `~/.bashrc`. Then add the following:

```
export GOPATH=<path to starter code>
export PATH=$PATH:$GOPATH/bin
```

2. Run `source ~/.bash_profile` or `source ~/.bashrc`

## Usage

1. After you have completed the `Setup`
steps, run the `build.sh` script provided with the starter code. This should
create 2 executables in the `bin` folder inside your code directory.

```shell
> ./build.sh
> ls bin
SurfstoreClientExec SurfstoreServerExec
```

2. Run your server using the script provided in the code.

```shell
./run-server.sh
```

3. From a new terminal (or a new node), run the client using the script
provided in the starter code (if using a new node, build using step 1 first).
Use a base directory with some files in it.

```shell
> mkdir dataA
> cp ~/pic.jpg dataA/ 
> ./run-client.sh server_addr:port dataA 4096
```

This would sync pic.jpg to the server hosted on `server_addr:port`, using
`dataA` as the base directory, with a block size of 4096 bytes.

4. From another terminal (or a new node), run the client to sync with the
server. (if using a new node, build using step 1 first)

```shell
> ls dataB/
> ./run-client.sh server_addr:port dataB 4096
> ls dataB/
pic.jpg index.txt
```

We observe that pic.jpg has been synced to this client.

Starter code provided by George Porter, UCSD.
