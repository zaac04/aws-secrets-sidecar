package writer

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func WriteFile(FileName string, Content string) error {

	file, err := os.OpenFile(FileName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	n, err := file.Write([]byte(Content))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n, "number of bytes wrote")

	return err
}

func CreateDir(vol string, dir string) (err error) {
	path := filepath.Join(vol, dir)

	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
	}

	return err
}
