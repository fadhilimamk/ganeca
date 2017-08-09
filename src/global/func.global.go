package global

import (
	"encoding/json"
	"fmt"
	"os"
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
