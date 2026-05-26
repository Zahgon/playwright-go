package playwright

import (
	"errors"
)

const fileSizeLimitInBytes = 50 * 1024 * 1024

var ErrInputFilesSizeExceeded = errors.New("Cannot set buffer larger than 50Mb, please write it to a file and pass its path instead.")

type inputFiles struct {
	Selector        *string             `json:"selector,omitempty"`
	Streams         []*channel          `json:"streams,omitempty"` // writableStream
	LocalPaths      []string            `json:"localPaths,omitempty"`
	Payloads        []map[string]string `json:"payloads,omitempty"`
	LocalDirectory  *string             `json:"localDirectory,omitempty"`
	DirectoryStream *channel            `json:"directoryStream,omitempty"`
}

type fileItem struct {
	LastModifiedMs *int64 `json:"lastModifiedMs,omitempty"`
	Name           string `json:"name"`
}

// convertInputFiles converts files to proper format for Playwright
//
//   - files should be one of: string, []string, InputFile, []InputFile,
//     string: local file path
func convertInputFiles(files any, context *browserContextImpl) (*inputFiles, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// local file path

// remote

func getFileLastModifiedMs(path string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func sizeOfInputFiles(files []InputFile) int { _ = "STUB: not implemented"; return 0 }

func normalizeFilePayloads(files []InputFile) []map[string]string {
	_ = "STUB: not implemented"
	return nil
}

func resolvePathsAndDirectoryForInputFiles(items []string) (localPaths []string, localDirectory *string, e error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// if the path is a directory

func listFiles(dir string) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }
