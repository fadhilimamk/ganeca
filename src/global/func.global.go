package global

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// ObjectToJSON convert object to json
func ObjectToJSON(p interface{}) string {
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

// GetCurrentTime return current time in unix format
func GetCurrentTime() int64 {
	t := time.Now()
	return t.Unix()
}
