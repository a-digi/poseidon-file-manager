// Package filemanager provides file management utilities for Poseidon plugins.
package filemanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// FileManagerClient interacts with the Poseidon file manager HTTP API.
type FileManagerClient struct{}

const baseURL = "http://localhost:2014"

type createFileRequest struct {
	PluginName string `json:"pluginName"`
	Filename   string `json:"filename"`
	Content    string `json:"content"`
}

type deleteFileRequest struct {
	PluginName string `json:"pluginName"`
	Filename   string `json:"filename"`
}

type getPathRequest struct {
	PluginName string `json:"pluginName"`
	Filename   string `json:"filename"`
}

type getPathResponse struct {
	Path string `json:"path"`
}

// CreateFile creates a new file with the given name and content.
func CreateFile(filename string, content []byte) error {
	return ioutil.WriteFile(filename, content, 0644)
}

// ReadFile reads the content of the specified file.
func ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// DeleteFile deletes the specified file.
func DeleteFile(filename string) error {
	return os.Remove(filename)
}

// CreateFileViaAPI creates a file via the HTTP API.
func (c *FileManagerClient) CreateFile(pluginName, filename, content string) error {
	url := baseURL + "/api/plugins/file-manager/create"
	body, _ := json.Marshal(createFileRequest{pluginName, filename, content})
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("create file failed: %s", string(respBody))
	}

	return nil
}

// DeleteFileViaAPI deletes a file via the HTTP API.
func (c *FileManagerClient) DeleteFile(pluginName, filename string) error {
	url := baseURL + "/api/plugins/file-manager/delete"
	body, _ := json.Marshal(deleteFileRequest{pluginName, filename})
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("delete file failed: %s", string(respBody))
	}
	return nil
}

// GetFilePath gets the file path via the HTTP API.
func (c *FileManagerClient) GetFilePath(pluginName, filename string) (string, error) {
	url := baseURL + "/api/plugins/file-manager/path"
	body, _ := json.Marshal(getPathRequest{pluginName, filename})
	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("get path failed: %s", string(respBody))
	}
	var res getPathResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}
	return res.Path, nil
}
