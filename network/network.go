package network

type NetworkInfo struct {
	NetworkId     byte
	ProtocolMagic uint32
}

func TestNet() *NetworkInfo {
	return &NetworkInfo{
		NetworkId:     0b0000,
		ProtocolMagic: 1097911063,
	}
}

func MainNet() *NetworkInfo {
	return &NetworkInfo{
		NetworkId:     0b0001,
		ProtocolMagic: 764824073,
	}
}

func PreProd() *NetworkInfo {
	return &NetworkInfo{
		NetworkId:     0b0010,
		ProtocolMagic: 1,
	}
}

func Preview() *NetworkInfo {
	return &NetworkInfo{
		NetworkId:     0b0011,
		ProtocolMagic: 2,
	}
}
