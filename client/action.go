package client

func userConnect() {
	data, err := packConnectData()
	if err != nil {
		mlog.Println(err)
		return
	}
	outDataChannel <- data
}

func userDisconnect() {
	data, err := packConnectData()
	if err != nil {
		mlog.Println(err)
		return
	}
	outDataChannel <- data
}

func userJoinGroup() {
	data, err := packJoinGroupData()
	if err != nil {
		mlog.Println(err)
		return
	}
	outDataChannel <- data
}

func userLeaveGroup() {
	data, err := packLeaveGroupData()
	if err != nil {
		mlog.Println(err)
		return
	}
	outDataChannel <- data
}

func userHoldMic() {
	data, err := packHoldMicData()
	if err != nil {
		mlog.Println(err)
		return
	}
	outDataChannel <- data
}

func userReleaseMic() {
	data, err := packReleaseMicData()
	if err != nil {
		mlog.Println(err)
		return
	}
	outDataChannel <- data
}
