package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func DeleteFileFromTemp(fileName string) error {
	tempDir := "../temp" // Specify the relative path to your temp directory

	// Check if the file exists in the temp directory
	filePath := filepath.Join(tempDir, fileName)
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			// File does not exist
			return fmt.Errorf("file %s does not exist in %s", fileName, tempDir)
		}
		// Some other error occurred when trying to access the file
		return fmt.Errorf("error accessing file %s: %w", fileName, err)
	}

	// Delete the file
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("error deleting file %s: %v", fileName, err)
	}

	fmt.Printf("Successfully deleted %s from %s\n", fileName, tempDir)
	return nil
}
