package dashf_test

import (
	"github.com/cdornsife/dashf"
	"reflect"
	"testing"
)

var controlTestStruct *TestStruct

func init() {
	controlTestStruct = &TestStruct{
		One: 1,
		Two: "two",
	}

	controlTestStruct.Three = append(controlTestStruct.Three, struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
	}{"a", "b", "c"})

}

// TestStruct represents the data in the file
type TestStruct struct {
	One   int    `json:"one"`
	Two   string `json:"two"`
	Three []struct {
		A string `json:"a"`
		B string `json:"b"`
		C string `json:"c"`
	} `json:"three"`
}

// TestDashF_JSON load and test tests/test.json
func TestDashF_JSON(t *testing.T) {
	doIt("tests/test.json", t)
}

// TestDashF_YAML load and test tests/test.yaml
func TestDashF_YAML(t *testing.T) {
	doIt("tests/test.yaml", t)
}

// TestDashF_YML load and test tests/test.yml
func TestDashF_YML(t *testing.T) {
	doIt("tests/test.yml", t)
}

// TestDashF_URL load and test from url
func TestDashF_URL(t *testing.T) {
	doIt("https://raw.githubusercontent.com/cdornsife/dashf/master/tests/test.yaml", t)
}

// TestDashF_STDIN load and test from stdin
func TestDashF_STDIN(t *testing.T) {
	// TODO:
	//doIt("https://raw.githubusercontent.com/cdornsife/dashf/master/tests/test.yaml", t)
}

// Nearly every test function will do the same thing.
func doIt(file string, t *testing.T) {

	var out = new(TestStruct)

	err := dashf.Unmarshal(file, &out)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(out, controlTestStruct) {
		t.Fatal("results don't match")
	}

}
