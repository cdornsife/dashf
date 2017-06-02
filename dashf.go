package dashf

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/ghodss/yaml"
)

// Unmarshal ...
func Unmarshal(input string, result interface{}) error {

	if input == "" {
		return errors.New("Missing input")
	}

	bytes, err := getBytesFromInput(input)

	if err != nil {
		return err
	}

	if bytes == nil || len(bytes) == 0 {
		return errors.New("File contained no data")
	}

	if input == "-" {
		if yaml.Unmarshal(bytes, &result) != nil {
			if json.Unmarshal(bytes, &result) != nil {
				return errors.New("failed to decode from stdin")
			}
		}
		return nil
	}

	ext := path.Ext(input)

	switch strings.ToLower(ext) {

	case ".yaml", ".yml":
		err = yaml.Unmarshal(bytes, &result)

	case ".json":
		err = json.Unmarshal(bytes, &result)

	default:
		err = fmt.Errorf("Unsupported file type: %s", ext)
	}

	return err
}

// getBytesFromInput read info from STDIN, URL or from a file.
func getBytesFromInput(input string) ([]byte, error) {

	switch {

	case input == "-":
		fi, err := os.Stdin.Stat()
		if err != nil {
			return nil, err
		}

		if fi.Mode()&os.ModeNamedPipe == 0 {
			return nil, errors.New("Not a stdin pipe")
		}

		return ioutil.ReadAll(os.Stdin)

	case strings.Index(input, "http://") == 0 || strings.Index(input, "https://") == 0:
		_, err := url.Parse(input)
		if err != nil {
			return nil, err
		}

		resp, err := http.Get(input)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		return ioutil.ReadAll(resp.Body)

	default:

		fi, err := os.Stat(input)
		switch {
		case err != nil:
			return nil, errors.New("No such file or directory")

		case fi.IsDir():
			// it's a dir!
			return nil, errors.New("Directories not supported yet")

		default:
			// it's a file!
			bytes, err := ioutil.ReadFile(input)
			if err != nil {
				return nil, err
			}
			if bytes == nil {
				return nil, errors.New("File is empty")
			}

			return bytes, err
		}

	}

}
