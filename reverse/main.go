package main

import "os"

func main() {
	if err := encode("flag.txt"); err != nil {
		panic(err)
	}
}

func encode(filename string) error {
	byteFile, err := openFile(filename)
	if err != nil {
		return err
	}
	newByteFile := []byte{}
	for _, b := range byteFile {
		newByteFile = append(newByteFile, b+12)
	}
	if err := os.WriteFile(filename, newByteFile, 0644); err != nil {
		return err
	}
	return nil
}

func openFile(filename string) ([]byte, error) {
	result, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return result, nil
}
