package errors

import (
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

const (
	ReadingXMLFail   = "ReadingXMLFail"
	WritingXMLFail   = "ReadingXMLFail"
	XMLUnmarshalFail = "XMLUnmarshalFail"
	XMLMarshalFail   = "XMLMarshalFail"
)

func NewReadingXMLFailErr(e error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    ReadingXMLFail,
		Message: errMsgs.ReadingXMLFail,
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewWritingXMLFailErr(e error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    WritingXMLFail,
		Message: errMsgs.WritingXMLFail,
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

func NewXMLMarshalFailErr(e error) error {
	return ExchangeKnowledgeBaseErr{
		Stat:    XMLMarshalFail,
		Message: errMsgs.XMLMarshalFail,
		Err:     e,
		Dt:      map[string]string{},
	}
}
