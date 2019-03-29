package main

import (
    "fmt"
    "bufio"
    "os"
    "path/filepath"
)

func LinesInFile(fileName string) []string {
    f, _ := os.Open(fileName)
    // Create new Scanner.
    scanner := bufio.NewScanner(f)
    result := []string{}
    for scanner.Scan() {
        line := scanner.Text()
        result = append(result, line)
    }
    return result
}


func main() {
    var files []string

    root := os.Args[1]
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    if err != nil {
        panic(err)
    }
    total_lines := 0
    for _, file := range files {
        lines := LinesInFile(file)
        total_lines += len(lines)
    }
    fmt.Println(total_lines)
}
