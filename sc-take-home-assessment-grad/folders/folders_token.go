package folders

import (
	"encoding/base64"
	"fmt"
	"strconv"
)

func generateNextToken(offset int) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d", offset)))
}

func decodeToken(token string) (int, error) {
	data, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return 0, err
	}
	offset, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}
	return offset, nil
}
