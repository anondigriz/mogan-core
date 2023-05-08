package errors

import (
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

const (
	ReadingXMLFail   = "ReadingXMLFail"
	XMLUnmarshalFail = "XMLUnmarshalFail"
)

func NewReadingXMLFailErr(e error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    ReadingXMLFail,
		Message: errMsgs.ReadingXMLFail,
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewXMLUnmarshalFailErr(e error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    XMLUnmarshalFail,
		Message: errMsgs.XMLUnmarshalFail,
		Err:     e,
		Dt:      map[string]string{},
	}
}
