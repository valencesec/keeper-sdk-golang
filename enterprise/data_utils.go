package enterprise

import (
	"encoding/json"
	"github.com/valencesec/keeper-sdk-golang/api"
	"github.com/valencesec/keeper-sdk-golang/internal/database"
)

func parseEncryptedData(encryptedData string, treeKey []byte) (result *database.EncryptedData, err error) {
	if len(encryptedData) > 0 {
		var data = api.Base64UrlDecode(encryptedData)
		if data, err = api.DecryptAesV1(data, treeKey); err == nil {
			err = json.Unmarshal(data, &result)
		}
	} else {
		err = api.NewKeeperError("Encrypted data is empty")
	}
	return
}
