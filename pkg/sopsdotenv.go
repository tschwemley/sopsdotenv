package sopsdotenv

import(
	"github.com/getsops/sops/v3/decrypt"
	"github.com/joho/godotenv"
)

type Env map[string]string

func LoadFile(path string) (Env, error) {
	envBytes, err := decrypt.File(path, "dotenv")
	if err != nil {
		panic(err)
	}
	return godotenv.UnmarshalBytes(envBytes)
}

