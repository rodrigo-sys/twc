package programfile

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ProgramFile struct {
	Path              string
	Example_file_path string
	Content           string
	Name              string
	Placeholders      map[string]func() (string, error)
}

func (f ProgramFile) GetPath() string {
	var default_path string
	var path string

	if f.Path == "" {
		config_dir, _ := os.UserConfigDir()
		default_path = filepath.Join(config_dir, "twc", f.Name)
		// default_path = filepath.Join(config_dir, filepath.Base(os.Args[0]), f.Name)

		path = default_path
	} else {
		path = f.Path
	}

	/* td: check if config_path is a fullpath
	or think what to do if it is relative */

	return path
}

func (f ProgramFile) CreateFile() {
	var source_file io.Reader
	example_file_path := f.Example_file_path

	if example_file_path == "" {
		file_content := f.Content

		// replace placeholders with actual values
		for placeholder, value := range f.Placeholders {
			value, _ := value()
			file_content = strings.ReplaceAll(file_content, placeholder, value)
		}

		// create a reader with the modified content
		source_file = strings.NewReader(file_content)

	} else {
		source_file, _ = os.Open(example_file_path)
		// defer source_file.Close()
	}

	destination_file, _ := os.Create(f.GetPath())
	defer destination_file.Close()
	io.Copy(destination_file, source_file)
}
