package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath string
	OutputFilePath string
}



func (fm FileManager)ReadLines()([]string, error){
	file, err := os.Open(fm.InputFilePath)

	if err!= nil{
		return nil, errors.New("Failed to read")
	}

	//executes when surrounding function finished execution (error or success)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		// file.Close()
		return nil, errors.New("Failed")
	}

		// file.Close()

		return lines, nil
}

func (fm FileManager)WriteResult(data interface{}) error{
		file, err := os.Create(fm.OutputFilePath)

		fmt.Println(file)

		defer file.Close()

		if err != nil {
			return errors.New("Failed to create file")
		}

		time.Sleep(3 * time.Second)

		encoder := json.NewEncoder(file)

		err = encoder.Encode(data)
		if err != nil {
			// file.Close()
			return errors.New("Failed to convert data to JSON")
		}

		// file.Close()

		return nil
}


func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}

}
