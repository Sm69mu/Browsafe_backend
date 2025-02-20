package services

import (
	"browsafe_backend/configs"
	"time"

	"github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/RtcTokenBuilder"
)

const RolePublisher = 1

func GenerateAgoraToken(channelName string, uID uint32) (string, error) {
	expireTimeInSeconds :=uint32(3600) //1hr
	currentTimestamp := uint32(time.Now().Unix())
	expireTimeStamps := currentTimestamp+expireTimeInSeconds


	token ,err := rtctokenbuilder.BuildTokenWithUID(
		configs.LoadAgoraConfigs().AppID,
		configs.LoadAgoraConfigs().AppCertificate,
		channelName,
		uID,
		RolePublisher,
		expireTimeStamps,
	)
	if err != nil {
		return "", err
	}
	return token, nil
}