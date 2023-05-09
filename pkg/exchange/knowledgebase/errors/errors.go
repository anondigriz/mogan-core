package errors


type ExchangeKnowledgeBaseErr struct {
	Stat    string
	Message string
	Err     error
	Dt      map[string]string
}

func (er ExchangeKnowledgeBaseErr) Status() string {
	return er.Stat
}

func (er ExchangeKnowledgeBaseErr) Error() string {
	return er.Message
}

func (er ExchangeKnowledgeBaseErr) Data() map[string]string {
	return er.Dt
}

func (er ExchangeKnowledgeBaseErr) Unwrap() error {
	return er.Err
}

func (er ExchangeKnowledgeBaseErr) IsExchangeKnowledgeBaseErr() bool {
	return true
}
