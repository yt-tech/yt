package client

func userLogin() {
	data, err := getLoginData()
	if err != nil {
		mlog.Println(err)
		return
	}
	outDataChannel <- data
}
