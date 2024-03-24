package fileutils

import "os"

func ReadTextFile(filename string) (string, error) {
  content, err := os.ReadFile(filename)

  if err != nil {
    // we could not read filename
    return "", err
  } else {
    // Read file 
    return string(content), nil
  }
}
