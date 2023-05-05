package knowledgebase

import "github.com/anondigriz/mogan-core/pkg/entities/types"

type Parameter struct {
	BaseInfo
	GroupUUID string
	Type      types.ParameterType
	ExtraData ExtraDataParameter
}

type ExtraDataParameter struct {
	Description  string
	DefaultValue string
}
