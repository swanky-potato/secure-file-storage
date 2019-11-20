package sfm

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Namespace contains the path and metadata about the namespace and is used to represent the folder on disk
type Namespace struct {
	Path  string   `json:"path"`              // Path location within storage directory and namespace name
	UID   int      `json:"userid,omitempty"`  // User ID of file owner
	GID   int      `json:"groupid,omitempty"` // Group ID of access group
	Extra struct{} `json:"extra,omitempty"`   // To store extra custom metadata
}

// GetNamespace reads the metadata from the namespace directory and return struct containing the data
func GetNamespace(nspath string) (Namespace, error) {
	spaces := Namespace{}
	// check if space exists
	if _, err := os.Stat(path.Join(StorageLocation, nspath)); err != nil {
		return Namespace{}, err
	}
	data, err := ioutil.ReadFile(path.Join(StorageLocation, nspath, ".meta-ns"))
	if err != nil {
		return Namespace{}, err
	}
	json.Unmarshal(data, &spaces)
	return spaces, nil
}

// Create a new namespace
func (ns *Namespace) Create() error {
	// check if space already exists
	if _, err := os.Stat(path.Join(StorageLocation, ns.Path)); err == nil {
		return errors.New("namespace already exists")
	}
	// create space
	if err := os.MkdirAll(path.Join(StorageLocation, ns.Path), os.ModePerm); err != nil {
		return err
	}
	// create metadata file for namespace
	f, err := os.Create(path.Join(StorageLocation, ns.Path, ".meta-ns"))
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := json.Marshal(ns)
	if err != nil {
		return err
	}
	if _, err := f.Write(b); err != nil {
		return err
	}
	return nil
}

// // AddNamespace creates a new namespace on storage based on Namspace Struct and stores the metadata
// func AddNamespace(space Namespace) error {
// 	// check if space already exists
// 	if _, err := os.Stat(path.Join(StorageLocation, space.Path)); err == nil {
// 		return errors.New("namespace already exists")
// 	}
// 	// create space
// 	if err := os.MkdirAll(path.Join(StorageLocation, space.Path), os.ModePerm); err != nil {
// 		return err
// 	}
// 	// create metadata file for namespace
// 	f, err := os.Create(path.Join(StorageLocation, space.Path, ".meta-ns"))
// 	if err != nil {
// 		return err
// 	}
// 	defer f.Close()
// 	b, err := json.Marshal(space)
// 	if err != nil {
// 		return err
// 	}
// 	if _, err := f.Write(b); err != nil {
// 		return err
// 	}
// 	return nil
// }

// Update that already exists with new metadata
func (ns *Namespace) Update() error {
	if _, err := os.Stat(path.Join(StorageLocation, ns.Path, ".meta-ns")); err != nil {
		return err
	}
	// remove old metadata file
	if err := os.Remove(path.Join(StorageLocation, ns.Path, ".meta-ns")); err != nil {
		return err
	}
	// create metadata file for namespace
	f, err := os.Create(path.Join(StorageLocation, ns.Path, ".meta-ns"))
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := json.Marshal(ns)
	if err != nil {
		return err
	}
	if _, err := f.Write(b); err != nil {
		return err
	}
	return nil
}

// Delete the namespaces and its content
func (ns *Namespace) Delete() error {
	// check if space already exists
	if _, err := os.Stat(path.Join(StorageLocation, ns.Path)); err != nil {
		if os.IsNotExist(err) {
			return errors.New("namespace does not exists")
		}
	}
	if err := os.RemoveAll(path.Join(StorageLocation, ns.Path)); err != nil {
		return err
	}
	return nil
}

// GetNamespaceMapBelow returns a slice of namespaces
func GetNamespaceMapBelow(nspath string) ([]Namespace, error) {
	if _, err := os.Stat(path.Join(StorageLocation, nspath)); err != nil {
		if os.IsNotExist(err) {
			return nil, errors.New("namespace does not exists")
		}
	}
	spaces := []Namespace{}
	err := filepath.Walk(path.Join(StorageLocation, nspath), func(pwd string, info os.FileInfo, err error) error {
		if info.IsDir() == true {
			splitpath := strings.Split(pwd, nspath)
			if _, err := os.Stat(path.Join(StorageLocation, nspath, splitpath[len(splitpath)-1], ".meta-ns")); err != nil {
				if os.IsNotExist(err) {
					// log.Println(errors.New("no namespace metadata file found in " + info.Name() + ", skipping"))
				} else {
					return err
				}
			} else {
				ns, err := GetNamespace(path.Join(nspath, splitpath[len(splitpath)-1]))
				if err != nil {
					return err
				}
				spaces = append(spaces, ns)
			}
		}
		return nil
	})

	if err != nil {
		return spaces, err
	}
	return spaces, nil

}
