package sfm

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path"
)

// File Is a data struct of a encrypted file on storage  including metadata
type File struct {
	ID        string      `json:"id"`                  // ID to track file and used as filename on the server storage
	Name      string      `json:"filename"`            // File name
	Extention string      `json:"extention"`           // File name
	Namespace string      `json:"namespace,omitempty"` // location on filesystem to group files and it can be used for limiting access to files
	Checksum  []byte      `json:"checksum"`            // checksum hash of decoded base64 string
	Content   []byte      `json:"content,omitempty"`   // string data stored byte slice
	UID       int         `json:"userid,omitempty"`    // Owner/Creator of the file
	GID       int         `json:"groupid,omitempty"`   // Group ID of access group
	Extra     interface{} `json:"extra,omitempty"`     // To store extra custom metadata prefer to be used with struct type
}

// ReadFile file from storage
func ReadFile(ns Namespace, fn string, pass string) (File, error) {
	fn = generateID(fn)
	e, err := ioutil.ReadFile(path.Join(StorageLocation, ns.Path, fn))
	if err != nil {
		return File{}, err
	}
	b64 := decrypt(e, pass)
	d, err := base64.StdEncoding.DecodeString(string(b64))
	if err != nil {
		return File{}, err
	}
	f := File{}
	json.Unmarshal(d, &f)
	return f, nil
}

// Store file on filesystem
func (f *File) Store(ns Namespace, passphrase string) error {
	f.ID = generateID(f.Name)
	if f.GID == 0 {
		f.GID = ns.GID
	}
	// always generate new checksum when storing
	f.Checksum = createChecksum(f.Content)
	// store file on storage as encrypted
	d, err := json.Marshal(f)
	if err != nil {
		return err
	}
	// b64 hash and encrypt write to file
	e := encrypt([]byte(base64.StdEncoding.EncodeToString(d)), passphrase)
	fl, err := os.Create(path.Join(StorageLocation, ns.Path, f.ID))
	if err != nil {
		return err
	}
	defer fl.Close()
	if _, err := fl.Write(e); err != nil {
		return err
	}
	return nil
}

// Update Content of stored data
func (f *File) Update(pass string) error {
	ns, err := GetNamespace(f.Namespace)
	if err != nil {
		return err
	}
	if err := f.Delete(); err != nil {
		return err
	}
	if err := f.Store(ns, pass); err != nil {
		return err
	}
	return nil
}

// Delete removes a file from storage
func (f *File) Delete() error {
	if err := os.Remove(path.Join(StorageLocation, f.Namespace, f.ID)); err != nil {
		return err
	}
	return nil
}

// generateID returns a md5 hash based on filename and time, used as ID for file
func generateID(fn string) string {
	h := md5.New()
	io.WriteString(h, fn)
	return hex.EncodeToString(h.Sum(nil))
}

// create checksum of input data
func createChecksum(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	chsm := []byte(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	return chsm
}
