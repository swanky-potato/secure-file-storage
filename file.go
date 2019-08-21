package sfm

import (
	"io/ioutil"
	"os"
	"path"
)

func writeFile(data []byte, filename string, passphrase string) error {
	f, err := os.Create(path.Join(SecretStore, filename+".secret"))
	if err != nil {
		return err
	}
	defer f.Close()
	f.Write(encrypt(data, passphrase))
	return nil
}

func readFile(filename, passphrase string) ([]byte, error) {
	data, err := ioutil.ReadFile(path.Join(SecretStore, filename+".secret"))
	if err != nil {
		return nil, err
	}
	return decrypt(data, passphrase), nil
}

func removeFile(filename, passphrase string) error {
	if err := os.Remove(path.Join(SecretStore, filename+".secret")); err != nil {
		return err
	}
	return nil
}

func checkFileExists(filename string) bool {
	if _, err := os.Stat(path.Join(SecretStore, filename+".secret")); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
