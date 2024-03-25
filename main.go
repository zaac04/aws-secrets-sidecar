package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"secrets_sidecar/env"
	"secrets_sidecar/initilizers"
	"secrets_sidecar/secrets"
	"secrets_sidecar/utilities"
	"secrets_sidecar/writer"
)

func main() {
	initilizers.Init()
	
	params, vol := env.ReadEnv()

	for _, param := range params {

		var secretsvalue map[string]string
		value, err := secrets.ReadSecret(param.Name, param.Region)

		if err != nil {
			fmt.Println(err)
			continue
		}

		err = json.Unmarshal([]byte(value), &secretsvalue)

		if err != nil {
			fmt.Println(err)
			continue
		}

		for key, value := range secretsvalue {
			if param.Field == "" || key == param.Field {
				err := writer.CreateDir(vol, param.Name)
				if err != nil {
					fmt.Println(err)
					break
				}

				val, err := base64.StdEncoding.DecodeString(value)
				if err != nil {
					fmt.Println(err)
				} else {
					value = string(val)
				}

				writer.WriteFile(utilities.GetWriteLocation(key, param.Name, vol), value)
			}
		}
	}
	fmt.Println("Done!")
}
