package vrcwss

import (
	"EternityGUI/cmd/vrchat/vrcapi"
	"EternityGUI/utils"
	"github.com/project-vrcat/vrchat-go"
	"github.com/project-vrcat/vrchat-go/events"
)

//InitClient initiates a websocket connection to VRChat and sends data out to a select statement
func InitClient() {
	vrc := vrchat.NewClient()

	vrc.RegisterEvent(events.FriendUpdate, func(params interface{}) {
		p := params.(events.FriendUpdateParams)
		data := "Friend Update: " + p.User.DisplayName
		vrcapi.FriendUpdate <- data
	})
	vrc.RegisterEvent(events.FriendLocation, func(params interface{}) {
		p := params.(events.FriendLocationParams)
		data := "Friend Location: " + p.User.DisplayName + " -> " + p.World.Name
		vrcapi.FriendLocation <- data
	})
	vrc.RegisterEvent(events.Notification, func(params interface{}) {
		p := params.(events.NotificationParams)
		data := "Notification: " + p.SenderUsername + " | " + p.Type
		vrcapi.VRCNotification <- data
	})
	vrc.RegisterEvent(events.FriendOnline, func(params interface{}) {
		p := params.(events.FriendOnlineParams)
		data := "Friend Online: " + p.User.DisplayName
		vrcapi.FriendOnline <- data
	})
	vrc.RegisterEvent(events.FriendOffline, func(params interface{}) {
		p := params.(events.FriendOfflineParams)
		data := "Friend Offline: " + p.UserID
		vrcapi.FriendOffline <- data
	})

	vrc.Pipeline.Connect(utils.ReadConfig().VRChatLogin)
	defer vrc.Pipeline.Close()

	// Select to stop reaching end of function
	select {}
}
