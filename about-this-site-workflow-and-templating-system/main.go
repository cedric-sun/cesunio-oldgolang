package main

import (
    "os"
    "path/filepath"
)

func Data(curDir string) map[string]interface{} {
    fileInfo, err := os.Stat(filepath.Join(curDir,"index.html"))
    if err!=nil {panic(err)}
    return map[string]interface{}{
        "lastEdit": fileInfo.ModTime().Format("2006-01-02 15:04:05 -0700 MST"),
    }
}
