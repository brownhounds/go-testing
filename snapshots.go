package go_testing

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	UPDATE        = "UPDATE"
	SNAPSHOTS_DIR = "SNAPSHOTS_DIR"
)

var snapshotsPath = filepath.Join(os.Getenv(SNAPSHOTS_DIR)) // nolint

func validateSnapshotsDir() {
	_, defined := os.LookupEnv(SNAPSHOTS_DIR)
	if !defined {
		panic(SNAPSHOTS_DIR + " is not defined")
	}
}

func isUpdate() bool {
	_, defined := os.LookupEnv(UPDATE)
	return defined
}

func readSnapshot(filename string) string {
	path := filepath.Join(snapshotsPath, filename)
	c, err := os.ReadFile(path)
	if err != nil {
		panic("snapshot file does not exists: " + path)
	}
	return string(c)
}

func createSnapshotDirectory() {
	if err := os.MkdirAll(snapshotsPath, 0o744); err != nil {
		panic(err)
	}
}

func writeSnapshot(filename, contents string) {
	if err := os.WriteFile(filepath.Join(snapshotsPath, filename), []byte(contents), 0o644); err != nil {
		panic(err)
	}
}

func fileName(tName, name string) string {
	var fileName string
	if strings.TrimSpace(name) != "" {
		fileName = tName + "_" + name
	} else {
		fileName = tName
	}

	return strings.ReplaceAll(strings.ReplaceAll(fileName, "/", "-"), " ", "")
}

func assert(t *testing.T, expected, received string) {
	if expected != received {
		ErrorMessage("Expected", expected)
		SuccessMessage("Received", received)
		t.Fatalf("values do not matching")
	}
}

func ToMatchSnapshot(t *testing.T, contents, name string) {
	validateSnapshotsDir()

	if isUpdate() {
		createSnapshotDirectory()
		writeSnapshot(fileName(t.Name(), name), contents)
	}

	c := readSnapshot(fileName(t.Name(), name))
	assert(t, contents, c)
}
