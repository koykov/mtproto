package mtproto

const (
	layer = 23

	// https://core.telegram.org/schema/mtproto
	//CrcVector                     = 0x1cb5c415
	CrcResPQ       = 0x05162463
	CrcPQInnerData = 0x83c95aec
	//CrcServerDHParamsFail      = 0x79cb045d
	CrcServerDHParamsOk  = 0xd0e8075c
	CrcServerDHInnerData = 0xb5890dba
	CrcClientDHInnerData = 0x6643b654
	CrcDHGenOk           = 0x3bcbf734
	//CrcDHGenRetry              = 0x46dc1fb9
	//CrcDHGenFail               = 0xa69dae02
	CrcRpcResult = 0xf35c6d01
	CrcRpcError  = 0x2144ca19
	//CrcRpcAnswerUnknown        = 0x5e2ad36e
	//CrcRpcAnswerDroppedRunning = 0xcd78e586
	//CrcRpcAnswerDropped        = 0xa43ad8b7
	//CrcFutureSalt              = 0x0949d9dc
	//CrcFutureSalts             = 0xae500895
	CrcPong = 0x347773c5
	//CrcDestroySessionOk        = 0xe22045fc
	//CrcDestroySessionNone      = 0x62d350c9
	CrcNewSessionCreated = 0x9ec20908
	CrcMsgContainer      = 0x73f1f8dc
	//CrcMsgCopy                 = 0xe06046b2
	CrcGzipPacked         = 0x3072cfa1
	CrcMsgsAck            = 0x62d6b459
	CrcBadMsgNotification = 0xa7eff811
	CrcBadServerSalt      = 0xedab447b
	//CrcMsgResendReq            = 0x7d861a08
	//CrcMsgsStateReq            = 0xda69fb52
	//CrcMsgsStateInfo           = 0x04deb57d
	//CrcMsgsAllInfo             = 0x8cc0d131
	//CrcMsgDetailedInfo         = 0x276d3ec6
	//CrcMsgNewDetailedInfo      = 0x809db6df
	CrcReqPq             = 0x60469778
	CrcReqDHParams       = 0xd712e4be
	CrcSetClientDHParams = 0xf5045f1f
	//CrcRpcDropAnswer           = 0x58e4a740
	//CrcGetFutureSalts          = 0xb921bd04
	CrcPing = 0x7abe77ec
	//CrcPingDelayDisconnect     = 0xf3427b8c
	//CrcDestroySession          = 0xe7512126
	//CrcHttpWait                = 0x9299359f
)
