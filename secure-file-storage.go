package sfm

// ChecksumStore is the location where to store checksum files
var ChecksumStore = ""

// SecretStore is location where to store the encrypted files
var SecretStore = ""

// Store will create and write a encrypted file to filesystem to the given location/filemane return error on fail else nil
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

// Exists trys to find the file checksum and encrypted file exists when file is not found it returns false
func Exists(filename string) bool {
	if c := checkFileExists(filename); c != true {
		return c
	}
	if c := checksumFileExists(filename); c != true {
		return c
	}
	return true
}

// CheckFileIsValid takes the stored checksum and dycrypts the file check if the checksum still matches the file.
// return a error when there is a missmatch
func CheckFileIsValid(filename, passphrase string) error {
	d, err := readFile(filename, passphrase)
	if err != nil {
		return err
	}
	if err := checkChecksum(d, filename); err != nil {
		return err
	}
	return nil
}

// Read will return decrypted content of the file or returns a error
func Read(filename, passphrase string) ([]byte, error) {
	d, err := readFile(filename, passphrase)
	if err != nil {
		return nil, err
	}
	if err := checkChecksum(d, filename); err != nil {
		return nil, err
	}
	return d, nil
}

// Remove only deletes the file of storage after decyption passes else return err
func Remove(filename, passphrase string) error {
	d, err := readFile(filename, passphrase)
	if err != nil {
		return err
	}
	if err := checkChecksum(d, filename); err != nil {
		return err
	}
	if err := removeFile(filename, passphrase); err != nil {
		return err
	}
	return nil
}
