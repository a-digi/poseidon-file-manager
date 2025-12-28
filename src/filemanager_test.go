package filemanager

import (
	"os"
	"testing"
)

func TestCreateReadDeleteFile(t *testing.T) {
	filename := "testfile.txt"
	content := []byte("Hello, Poseidon!")

	// Test CreateFile
	err := CreateFile(filename, content)
	if err != nil {
		t.Fatalf("CreateFile failed: %v", err)
	}

	// Test ReadFile
	readContent, err := ReadFile(filename)
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}
	if string(readContent) != string(content) {
		t.Fatalf("ReadFile content mismatch: got %s, want %s", string(readContent), string(content))
	}

	// Test DeleteFile
	err = DeleteFile(filename)
	if err != nil {
		t.Fatalf("DeleteFile failed: %v", err)
	}

	// Ensure file is deleted
	_, err = os.Stat(filename)
	if !os.IsNotExist(err) {
		t.Fatalf("File still exists after DeleteFile")
	}
}

func TestFileManagerClient_API(t *testing.T) {
	client := &FileManagerClient{}
	pluginName := "testplugin"
	filename := "apitestfile.txt"
	content := "{\"foo\":123}"

	// Test CreateFile (API)
	err := client.CreateFile(pluginName, filename, content)
	if err != nil {
		t.Fatalf("CreateFile (API) failed: %v", err)
	}

	// Test GetFilePath (API)
	path, err := client.GetFilePath(pluginName, filename)
	if err != nil {
		t.Fatalf("GetFilePath (API) failed: %v", err)
	}
	if path == "" {
		t.Fatalf("GetFilePath (API) returned empty path")
	}

	// Test DeleteFile (API)
	err = client.DeleteFile(pluginName, filename)
	if err != nil {
		t.Fatalf("DeleteFile (API) failed: %v", err)
	}
}
