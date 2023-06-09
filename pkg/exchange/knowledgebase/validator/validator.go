package validator

import (
	"embed"
	"io"

	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xsd"
	"go.uber.org/zap"

	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
)

//go:embed schemes/*.xsd
var embedSchemes embed.FS

type Validator struct {
	lg             *zap.Logger
	xsdFileBufV2M0 []byte
}

func New(lg *zap.Logger) (*Validator, error) {
	xsdFileV2M0, err := embedSchemes.Open("schemes/v2m0.xsd")
	if err != nil {
		lg.Error(errMsgs.OpeningXSDFail, zap.Error(err))
		return nil, errors.NewOpeningXSDFailErr(err)
	}
	defer xsdFileV2M0.Close()

	xsdFileBufV2M0, err := io.ReadAll(xsdFileV2M0)
	if err != nil {
		return nil, errors.NewReadingXSDFailErr(err)
	}
	v := &Validator{
		lg:             lg,
		xsdFileBufV2M0: xsdFileBufV2M0,
	}
	return v, nil
}

func (v *Validator) ValidateV2M0(xmlFile io.Reader) error {
	s, err := xsd.Parse(v.xsdFileBufV2M0)
	if err != nil {
		v.lg.Error(errMsgs.ParsingXSDFail, zap.Error(err))
		return errors.NewParsingXSDFailErr(err)
	}
	defer s.Free()

	buf, err := io.ReadAll(xmlFile)
	if err != nil {
		v.lg.Error(errMsgs.ReadingXMLFail, zap.Error(err))
		return errors.NewReadingXMLFailErr(err)
	}
	d, err := libxml2.Parse(buf)
	if err != nil {
		v.lg.Error(errMsgs.ParsingXMLFail, zap.Error(err))
		return errors.NewReadingXMLFailErr(err)
	}
	defer d.Free()

	if err = s.Validate(d); err != nil {
		v.lg.Error(errMsgs.ValidationXMLFail, zap.Error(err))
		return errors.NewValidationXMLFailErr(err, err.(xsd.SchemaValidationError).Errors())
	}
	return nil
}
