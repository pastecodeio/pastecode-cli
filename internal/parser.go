package internal

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func GetContents(fileFlag *string) ([]byte, error) {
	if *fileFlag != "" {
		fileContents, err := os.ReadFile(*fileFlag)
		if err != nil {
			err = fmt.Errorf("cannot read file: %s\n", err.Error())
			return nil, err
		}
		return fileContents, nil
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		stdinContents, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("cannot read redirected contents: %s\n", err.Error())
		}
		return stdinContents, nil
	}
	return nil, errors.New("you should either pipe contents to program or read a file directly")
}
