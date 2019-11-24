package biz

import (
    "bufio"
    "encoding/json"
    "fmt"
    "io"
    "os"
)


func ReadJSONFile(filename string) []map[string]string {

    file, err := os.Open(filename)

    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return nil
    }
    defer file.Close()

    retMapArr := make([]map[string]string, 0)

    reader := bufio.NewReader(file)
    for {
        data, _, errRead := reader.ReadLine()
        if errRead == io.EOF {
            break
        }

        retMapArr = append(retMapArr, readJSONData(data))
    }

    return retMapArr
}


func readJSONData(jsonData []byte) map[string]string {

    var ret map[string]string

    err := json.Unmarshal(jsonData, &ret)

    if err != nil {
        ret = nil
    }

    return ret
}
