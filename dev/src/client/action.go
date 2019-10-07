package client

func userConnect() {

	// outDataChannel <- data
}

func userDisconnect() {
	data, err := packConnectData()
	if err != nil {
		mlog.Println(err)
		return
	}
	outDataChannel <- data
}

func subscribeTopic() {

	// outDataChannel <- data
}

// func userLeaveGroup() {
// 	data, err := packLeaveGroupData()
// 	if err != nil {
// 		mlog.Println(err)
// 		return
// 	}
// 	outDataChannel <- data
// }

// func userHoldMic() {
// 	data, err := packHoldMicData()
// 	if err != nil {
// 		mlog.Println(err)
// 		return
// 	}
// 	outDataChannel <- data
// }

// func userReleaseMic() {
// 	data, err := packReleaseMicData()
// 	if err != nil {
// 		mlog.Println(err)
// 		return
// 	}
// 	outDataChannel <- data
// }
