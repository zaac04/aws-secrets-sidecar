package initilizers

import (
	"fmt"

	"github.com/joho/godotenv"
)

func Init() {

	fmt.Println("Started!")
	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
	}
}
