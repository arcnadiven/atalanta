package gotools

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"github.com/labstack/gommon/bytes"
	"io/ioutil"
	"os"
	"strings"
)

const (
	hash_block_size = bytes.MB * 1
)

func ReadFileInLine(path string) ([]string, error) {
	resp, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.Replace(string(resp), "\r", "", -1), "\n"), nil
}

func WriteFileInLine(path string, data []string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()
	for _, v := range data {
		if _, err := file.WriteString(v + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func getFileSHA1(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return ``, err
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		return ``, err
	}
	file_block_num := fi.Size()/hash_block_size + 1
	if fi.Size()%hash_block_size == 0 {
		file_block_num--
	}
	hSHA1 := sha1.New()
	for i := int64(0); i < file_block_num; i++ {
		buf := make([]byte, hash_block_size)
		if _, err := file.Seek(i*hash_block_size, os.SEEK_SET); err != nil {
			return ``, err
		}
		offset, err := file.Read(buf)
		if err != nil {
			return ``, err
		}
		if _, err := hSHA1.Write(buf[:offset]); err != nil {
			return ``, err
		}
	}
	return hex.EncodeToString(hSHA1.Sum(nil)), nil
}

func getFileMD5(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return ``, err
	}
	defer file.Close()
	fi, err := file.Stat()
	if err != nil {
		return ``, err
	}
	file_block_num := fi.Size()/hash_block_size + 1
	if fi.Size()%hash_block_size == 0 {
		file_block_num--
	}
	hMD5 := md5.New()
	for i := int64(0); i < file_block_num; i++ {
		buf := make([]byte, hash_block_size)
		if _, err := file.Seek(i*hash_block_size, os.SEEK_SET); err != nil {
			return ``, err
		}
		offset, err := file.Read(buf)
		if err != nil {
			return ``, err
		}
		if _, err := hMD5.Write(buf[:offset]); err != nil {
			return ``, err
		}
	}
	return hex.EncodeToString(hMD5.Sum(nil)), nil
}
