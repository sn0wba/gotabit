package types

const (
	// ModuleName is the name of the module
	ModuleName = "sendmsg"

	// StoreKey is the string store representation
	StoreKey string = ModuleName

	// QuerierRoute is the querier route for the module
	QuerierRoute string = ModuleName

	// RouterKey is the msg router key for the module
	RouterKey string = ModuleName
)

var (
	PrefixSendMsg           = []byte{0x01}
	PrefixSendMsgBySender   = []byte{0x02}
	PrefixSendMsgByReceiver = []byte{0x03}
	KeyLastSendMsgId        = []byte{0x04}
)
