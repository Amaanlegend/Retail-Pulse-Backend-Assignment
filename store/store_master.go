package store

import (
    "encoding/json"
    "os"
)

var storeData map[string]string

func LoadStoreMasterData(filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    return json.NewDecoder(file).Decode(&storeData)
}

func IsValidStoreID(storeID string) bool {
    _, exists := storeData[storeID]
    return exists
}
