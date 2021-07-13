package surfstore

// import {
// 	"fmt"
// }

type MetaStore struct {
	FileMetaMap map[string]FileMetaData
}

func (m *MetaStore) GetFileInfoMap(_ignore *bool, serverFileInfoMap *map[string]FileMetaData) error {
	//panic("todo")
	*serverFileInfoMap = m.FileMetaMap
	return nil
}

func (m *MetaStore) UpdateFile(fileMetaData *FileMetaData, latestVersion *int) (err error) {
	//panic("todo")
	filename := (*fileMetaData).Filename
	if curr_file_metadata, ok := m.FileMetaMap[filename]; ok{
		curr_version := curr_file_metadata.Version
		new_version := (*fileMetaData).Version
		var temp FileMetaData
		if curr_version == new_version - 1 {
			temp.BlockHashList = (*fileMetaData).BlockHashList
			temp.Version = new_version
			temp.Filename = filename
			*latestVersion = new_version
			m.FileMetaMap[filename] = temp
			return nil
		}
		*latestVersion = -1
		return nil //fmt.Errorf("SERVER ERROR: File version outdated!! Unable to update file in MetaStore")
	}else{
		//newly created file
		m.FileMetaMap[filename] = *fileMetaData
		new_version := (*fileMetaData).Version
		*latestVersion = new_version
		return nil
	}
}

var _ MetaStoreInterface = new(MetaStore)
