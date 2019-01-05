package types

type Ip []byte

func (this *Ip) isEmpty() bool {
	return len(*this) == 0
}

func (this *Ip) correctIp() bool {
	return len(*this) == 4
}