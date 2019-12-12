package client

func (c *clientInfo) cmdSend() {
	mlog.Println("cmdSend")
	for {
		om := <-c.outChan
		data, err := cmdMarsal(om)
		if err != nil {
			mlog.Println(err)
			continue
		}

		n, err := c.quicStream.Write(data)
		if err != nil {
			mlog.Println(err)
		}
		mlog.Println("send ", n)
	}
}
