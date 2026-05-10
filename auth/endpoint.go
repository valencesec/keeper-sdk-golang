package auth

import (
	"github.com/valencesec/keeper-sdk-golang/proto_auth"
)

type IKeeperEndpoint interface {
	ClientVersion() string
	SetClientVersion(string)
	DeviceName() string
	SetDeviceName(string)
	Locale() string
	SetLocale(string)
	Server() string
	SetServer(string)
	ServerKeyId() int32

	CommunicateKeeper(string, []byte, []byte) ([]byte, error)

	PushServer() string
	ConnectToPushServer(*proto_auth.WssConnectionRequest) (IPushEndpoint, error)
	ConfigurationStorage() IConfigurationStorage
}
