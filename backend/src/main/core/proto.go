package core

type PpProto struct {
	Code   []byte
	Length uint32
	PpType byte
	Token  string
	UUID   string
	Body   string
}

func (o *PpProto) WriteCode(Code []byte) *PpProto {
	o.Code = Code
	return o
}

func (o *PpProto) WriteLength(Length uint32) *PpProto {
	o.Length = Length
	return o
}
func (o *PpProto) WritePpType(PpType byte) *PpProto {
	o.PpType = PpType
	return o
}
func (o *PpProto) WriteToken(Token string) *PpProto {
	o.Token = Token
	return o
}

func (o *PpProto) WriteUUID(UUID string) *PpProto {
	o.UUID = UUID
	return o
}

func (o *PpProto) WriteBody(Body string) *PpProto {
	o.Body = Body
	return o
}
