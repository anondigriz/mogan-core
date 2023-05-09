package knowledgebase_test

import (
	"os"
	"testing"

	uuidGen "github.com/google/uuid"

	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/collector"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/parser"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestTransformationFromV2M0ToV2M0(t *testing.T) {
	type args struct {
		path     string
		toFormat formats.Version
	}
	type want struct {
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "test file #1 v2m0 to v2m0",
			args: args{
				path:     "./testdata/v2m0/model1.xml",
				toFormat: formats.VersionV2M0,
			},
		},
		{
			name: "test file #2 v2m0 to v2m0",
			args: args{
				path:     "./testdata/v2m0/model2.xml",
				toFormat: formats.VersionV2M0,
			},
		},
		{
			name: "test file #3 v2m0 to v2m0",
			args: args{
				path:     "./testdata/v2m0/model3.xml",
				toFormat: formats.VersionV2M0,
			},
		},
		{
			name: "test file #1 v2m0 to v3m0",
			args: args{
				path:     "./testdata/v2m0/model1.xml",
				toFormat: formats.VersionV3M0,
			},
		},
		{
			name: "test file #2 v2m0 to v3m0",
			args: args{
				path:     "./testdata/v2m0/model2.xml",
				toFormat: formats.VersionV3M0,
			},
		},
		{
			name: "test file #3 v2m0 to v3m0",
			args: args{
				path:     "./testdata/v2m0/model3.xml",
				toFormat: formats.VersionV3M0,
			},
		},
		{
			name: "test file #1 v3m0 to v2m0",
			args: args{
				path:     "./testdata/v3m0/model1.xml",
				toFormat: formats.VersionV2M0,
			},
		},
		{
			name: "test file #2 v3m0 to v2m0",
			args: args{
				path:     "./testdata/v3m0/model2.xml",
				toFormat: formats.VersionV2M0,
			},
		},
		{
			name: "test file #3 v3m0 to v2m0",
			args: args{
				path:     "./testdata/v3m0/model3.xml",
				toFormat: formats.VersionV2M0,
			},
		},
		{
			name: "test file #1 v3m0 to v3m0",
			args: args{
				path:     "./testdata/v3m0/model1.xml",
				toFormat: formats.VersionV3M0,
			},
		},
		{
			name: "test file #2 v3m0 to v3m0",
			args: args{
				path:     "./testdata/v3m0/model2.xml",
				toFormat: formats.VersionV3M0,
			},
		},
		{
			name: "test file #3 v3m0 to v3m0",
			args: args{
				path:     "./testdata/v3m0/model3.xml",
				toFormat: formats.VersionV3M0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			from, err := os.Open(tt.args.path)
			assert.NoError(t, err, "fail to open the xml file")
			defer from.Close()

			lg, err := zap.NewDevelopment()
			assert.NoError(t, err, "fail init logger")
			p := parser.New(lg)

			cont, err := p.Parse(parser.ParseXMLArgs{
				KnowledgeBaseUUID: uuidGen.NewString(),
				XMLFile:           from,
				FileName:          from.Name(),
			})
			assert.NoError(t, err, "no error was expected when parsing XML file")
			to, err := os.CreateTemp("./testdata", "tmpfile-")
			defer func() {
				if err = to.Close(); err != nil {
					t.Fatalf("Close() error = %v", err)
				}
				if err = os.Remove(to.Name()); err != nil {
					t.Fatalf("Remove() error = %v", err)
				}
			}()
			assert.NoError(t, err, "no error was expected when creating the temp file")

			c := collector.New(lg)
			err = c.Collect(collector.ParseXMLArgs{
				Version: tt.args.toFormat,
				Cont:    &cont,
				XMLFile: to,
				Prefix:  "",
				Indent:  "  ",
			})
			assert.NoError(t, err, "no error was expected when collecting XML file")
		})
	}
}
