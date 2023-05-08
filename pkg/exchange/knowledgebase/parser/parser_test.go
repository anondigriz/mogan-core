package parser_test

import (
	"context"
	"os"
	"testing"

	uuidGen "github.com/google/uuid"

	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/parser"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestV2M0(t *testing.T) {
	t.Run("check import from file", func(t *testing.T) {
		f, err := os.Open("./testdata/model.xml")
		assert.NoError(t, err, "fail to open the xml file")
		defer f.Close()

		lg, err := zap.NewDevelopment()
		assert.NoError(t, err, "fail init logger")
		p := parser.New(lg)

		iArgs := parser.ParseXMLArgs{
			KnowledgeBaseUUID: uuidGen.NewString(),
			XMLFile:           f,
			FileName:          f.Name(),
		}
		kbu, err := p.Parse(context.Background(), iArgs)
		assert.NoError(t, err, "no error was expected when parsing xml file")
		// TODO test is incomplete and currently only needed for visual verification
		_ = kbu
	})
}
