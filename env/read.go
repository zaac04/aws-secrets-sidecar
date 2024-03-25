package env

import (
	"encoding/json"
	"fmt"
	"os"
	"secrets_sidecar/structs"
)

func ReadEnv() (SecretsParams []structs.SecretParams, VolumeMount string) {
	var SecretParams []structs.SecretParams
	params := os.Getenv("secret_params")

	err := json.Unmarshal([]byte(params), &SecretParams)
	if err != nil {
		fmt.Println(err)
	}

	vol_mnt := os.Getenv("vol_mnt")

	fmt.Printf("Params Received : %+v, Volume : %+v ", SecretParams, vol_mnt)

	return SecretParams, vol_mnt
}
