package fromformat

import (
	"strings"

	"go.uber.org/zap"

	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

func (ff FromFormat) extractDictionaryFromAttribute(base string) (map[string]string, error) {
	params := map[string]string{}
	if base == "" {
		return params, nil
	}

	pairs := strings.Split(base, ";")
	for _, v := range pairs {
		keyValue := strings.Split(v, ":")
		if len(keyValue) != 2 {
			err := errors.NewBrokenDictionaryInXMLAttributeErr(base, v)
			ff.lg.Error(errMsgs.BrokenDictionaryInXMLAttribute, zap.Error(err))
			return nil, err
		}
		params[keyValue[0]] = keyValue[1]
	}
	return params, nil
}
