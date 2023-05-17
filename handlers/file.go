package handlers

type FileList []File

type File struct {
	Name string `json:"name"`
	// Type string `json:"type"`
	Path  string `json:"path,omitempty"`
	IsDir bool   `json:"isdir"`
}
