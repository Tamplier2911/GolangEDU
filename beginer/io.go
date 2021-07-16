package main

import (
	"bufio"
	"io/ioutil"
	"os"
)

type InputOutput struct {
}

// Initialize and setup InputOutput abstraction layer, returns instance of InputOutput
func (io InputOutput) Setup() InputOutput {
	// init client
	return io
}

// make directory structure using provided path
func (io *InputOutput) MkDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// srite using iotil writer
func (io *InputOutput) WriteIOUtil(path string, bs []byte) error {
	err := ioutil.WriteFile(path, bs, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// write using os
func (io *InputOutput) WriteOS(path string, bs []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// write
	_, err = f.Write(bs)
	if err != nil {
		return err
	}

	/*
		// or write as string
		_, err = f.WriteString(string(bs[:]))
		if err != nil {
			log.Fatal("failed to write a file as string", err)
			return err
		}
	*/
	f.Sync()
	return nil
}

// using bufio
func (io *InputOutput) WriteBUFIO(path string, bs []byte) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(bs[:]))
	if err != nil {
		return err
	}
	w.Flush()

	return nil
}

// https://stackoverflow.com/questions/1821811/how-to-read-write-from-to-a-file-using-go
