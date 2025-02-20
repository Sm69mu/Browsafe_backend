package configs

import "os"

type AgoraConfig struct {
	AppID          string
	AppCertificate string
}

func LoadAgoraConfigs() AgoraConfig {
	return AgoraConfig{
		AppID: os.Getenv("AGORA_APP_ID"),
		AppCertificate: os.Getenv("AGORA_APP_CERTIFICATE"),
	}
}
