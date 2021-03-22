package entity

const (
	UnknownRelationType = iota
	QQ
	Tel
	Email
	Address
)

var RelationTypeMap = map[int]string{
	UnknownRelationType: "unknown",
	QQ:                  "qq",
	Tel:                 "tel",
	Email:               "email",
	Address:             "address",
}

func RelationTypeToString(tp int) string {
	if res, ok := RelationTypeMap[tp]; ok {
		return res
	}
	return RelationTypeMap[UnknownRelationType]
}
