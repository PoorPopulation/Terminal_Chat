package core

type ppProto struct {
	code   []byte
	length int
	ppType byte
	token  string
	UUID   string
	body   string
}

func (o *ppProto) writeCode(code []byte) *ppProto {
	o.code = code
	return o
}

func (o *ppProto) writeLength(length int) *ppProto {
	o.length = length
	return o
}
func (o *ppProto) writePpType(ppType byte) *ppProto {
	o.ppType = ppType
	return o
}
func (o *ppProto) writeToken(token string) *ppProto {
	o.token = token
	return o
}

func (o *ppProto) writeUUID(UUID string) *ppProto {
	o.UUID = UUID
	return o
}

func (o *ppProto) writeBody(body string) *ppProto {
	o.body = body
	return o
}
