package etc

import "os"

func ReadProps() map[string]interface{} {
	return map[string]interface{}{
		"mode": os.Getenv("MODE"),
	}
}
