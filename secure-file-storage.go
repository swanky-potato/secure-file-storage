package sfm

// ChecksumStore is the location where to store checksum files
var ChecksumStore = ""

// SecretStore is location where to store the encrypted files
var SecretStore = ""

// Store a file on storage and generates a checksum it will return if fails it will return nil, err
func Store(data []byte, filename string, passphrase string) ([]byte, error) {
	// create new checksum for encrypted file
	checksum, err := createChecksum(data)
	if err != nil {
		return nil, err
	}
	// create new file
	if err := writeFile(data, filename, passphrase); err != nil {
		return nil, err
	}
	// save the checksum
	if err := storeChecksum(checksum, filename); err != nil {
		return nil, err
	}
	return checksum, nil
}

// Exists checks if both the checksum and encrypted file exist on the configured storage.
func Exists(filename string) bool {
	if c := checkFileExists(filename); c != true {
		return c
	}
	if c := checksumFileExists(filename); c != true {
		return c
	}
	return true
}

// ValidateFileStored takes the stored checksum and dycrypts the file check if the checksum still matches the data
func ValidateFileStored(filename, passphrase string) error {
	d, err := readFile(filename, passphrase)
	if err != nil {
		return err
	}
	if err := checkChecksumOnStorage(d, filename); err != nil {
		return err
	}
	return nil
}

// ValidateFileWithChecksum takes a checksum string and checks if the decrypted data stored still matched with this checksum
func ValidateFileWithChecksum(checksum, filename, passphrase string) error {
	d, err := readFile(filename, passphrase)
	if err != nil {
		return err
	}
	if err := checkChecksumFromInput(d, checksum); err != nil {
		return err
	}
	return nil
}

// Read will return decrypted content of the file or return a error
func Read(filename, passphrase string) ([]byte, error) {
	d, err := readFile(filename, passphrase)
	if err != nil {
		return nil, err
	}
	if err := checkChecksumOnStorage(d, filename); err != nil {
		return nil, err
	}
	return d, nil
}

// Remove only deletes the file of storage but only if you are able to decrypt it correctly
func Remove(filename, passphrase string) error {
	d, err := readFile(filename, passphrase)
	if err != nil {
		return err
	}
	if err := checkChecksumOnStorage(d, filename); err != nil {
		return err
	}
	if err := removeFile(filename, passphrase); err != nil {
		return err
	}
	return nil
}
