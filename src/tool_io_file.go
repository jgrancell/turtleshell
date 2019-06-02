package main

import (
    "os"
)

func create_file(path string) bool {
    // Detect if the path already exists
    _, err := os.Stat(path)

    if os.IsNotExist(err) {
        file, err := os.Create(path)
        if err != nil {
            return false
        }
        defer file.Close()
    }

    return true
}
