package utils

import (
	"log"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func MkdirNotExits(path string) {
	exist, err := PathExists(path)
	if err != nil {
		log.Println(err.Error())
		return
	}
	if !exist {
		os.Mkdir(path, os.ModePerm)
	}
}
func GetFilePath() string {
	path := "./upload-video/"
	MkdirNotExits(path)
	return path
}
