package sfm

import "testing"

var filespath = "files/"
var checksumpath = "files/"
var filename = "unit-test.txt"
var passphrase = "popcorn"
var data = []byte("Some long byte slice")
var checksum = []byte("")

func TestStore(t *testing.T) {
	if ChecksumStore == "" || SecretStore == "" {
		ChecksumStore = checksumpath
		SecretStore = filespath
	}
	ch, err := Store(data, filename, passphrase)
	if err != nil {
		t.Error(err)
	}
	checksum = ch
}

func TestRead(t *testing.T) {
	if ChecksumStore == "" || SecretStore == "" {
		ChecksumStore = checksumpath
		SecretStore = filespath
	}
	d, err := Read(filename, passphrase)
	if err != nil {
		t.Error(err)
	}
	if string(d) != string(data) {
		t.Error("returned data does not match input data")
	}
}

func TestRemove(t *testing.T) {
	if ChecksumStore == "" || SecretStore == "" {
		ChecksumStore = checksumpath
		SecretStore = filespath
	}
	err := Remove(filename, passphrase)
	if err != nil {
		t.Error(err)
	}
}
