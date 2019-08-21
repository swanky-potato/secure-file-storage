package sfm

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
	"path"
)

func createChecksum(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	chsm := []byte(base64.StdEncoding.EncodeToString(h.Sum(nil)))
	return chsm, nil
}

func checkChecksumOnStorage(data []byte, filename string) error {
	cs, err := createChecksum(data)
	if err != nil {
		return err
	}
	fs, err := readChecksumFile(filename)
	if err != nil {
		return err
	}
	if string(cs) == string(fs) {
		return nil
	}
	return errors.New("Checksum stored does not match data")
}

func checkChecksumFromInput(data []byte, checksum string) error {
	cs, err := createChecksum(data)
	if err != nil {
		return err
	}
	if string(cs) == checksum {
		return nil
	}
	return errors.New("Checksum given does not match the data")
}

func checksumFileExists(filename string) bool {
	if _, err := os.Stat(path.Join(ChecksumStore + filename + ".checksum")); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func readChecksumFile(filename string) ([]byte, error) {
	p := path.Join(ChecksumStore + filename + ".checksum")
	d, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func removeChecksum(filename string) error {
	p := path.Join(ChecksumStore + filename + ".checksum")
	if err := os.Remove(p); err != nil {
		return err
	}
	return nil
}

func updateChecksum(checksum []byte, filename string) error {
	if err := removeChecksum(filename); err != nil {
		return err
	}
	if err := storeChecksum(checksum, filename); err != nil {
		return err
	}
	return nil
}

func storeChecksum(checksum []byte, filename string) error {
	path := path.Join(ChecksumStore + filename + ".checksum")
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err = f.Write(checksum); err != nil {
		return err
	}
	return nil
}
