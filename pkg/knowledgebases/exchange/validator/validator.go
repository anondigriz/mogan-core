package validator

import (
	"bufio"
	"embed"
	"io"
	"regexp"

	"github.com/anondigriz/mogan-core/pkg/knowledgebases/exchange/validator/errors"
	"github.com/lestrrat-go/libxml2"
	"github.com/lestrrat-go/libxml2/xsd"
	"go.uber.org/zap"
)

const defaultMaxCheckLines = 15

//go:embed schemes/*.xsd
var embedSchemes embed.FS

type Validator struct {
	lg             *zap.Logger
	xsdFileBufV2M0 []byte
	maxCheckLines  int
}

func New(lg *zap.Logger) (*Validator, error) {
	xsdFileV2M0, err := embedSchemes.Open("schemes/v2m0.xsd")
	if err != nil {
		lg.Error("open XSD file was failed", zap.Error(err))
		return nil, errors.NewOpeningXSDFailErr(err)
	}
	defer xsdFileV2M0.Close()

	xsdFileBufV2M0, err := io.ReadAll(xsdFileV2M0)
	if err != nil {
		return nil, errors.NewReadingXSDFailErr(err)
	}
	return &Validator{lg: lg, xsdFileBufV2M0: xsdFileBufV2M0, maxCheckLines: defaultMaxCheckLines}, nil
}

// SetMaxCheckLines - set maximum number of lines for checks
func (v *Validator) SetMaxCheckLines(number int) {
	v.maxCheckLines = number
}

func (v *Validator) DetectVersion(scanner *bufio.Scanner) (string, error) {
	// TODO - если сюда прилетает бинарник, то будет плохо.
	// Также сюда прилетают файлы в одной строчке (сжатые XML).
	// Так как сюда прилетают файлы < 10 МБ, то предлагается пока забить.
	// Также не получается читать N первых байтов, так как аттрибут 'description'
	// может иметь любую длину, а аттрибут 'formatXmlVersion' может находиться необязательно
	// в начале.
	scanner.Split(bufio.ScanLines)

	var version string

	for i := 0; i < v.maxCheckLines; i++ {
		scanner.Scan()
		text := scanner.Text()

		re := regexp.MustCompile(`formatXmlVersion="\d\.\d"`)
		matched := re.FindString(text)
		if matched != "" {
			reVer := regexp.MustCompile(`\d\.\d`)
			version = reVer.FindString(matched)
			break
		}
	}

	if version == "" {
		return "", errors.NewFormatXMLVersionNotFoundErr()
	}

	return version, nil
}

func (v *Validator) ValidateV2M0(xmlFile io.Reader) error {
	s, err := xsd.Parse(v.xsdFileBufV2M0)
	if err != nil {
		v.lg.Error("failed to parse XSD file", zap.Error(err))
		return errors.NewParsingXSDFailErr(err)
	}
	defer s.Free()

	buf, err := io.ReadAll(xmlFile)
	if err != nil {
		v.lg.Error("failed to read XML file", zap.Error(err))
		return errors.NewReadingXMLFailErr(err)
	}
	d, err := libxml2.Parse(buf)
	if err != nil {
		v.lg.Error("failed to parse XML file", zap.Error(err))
		return errors.NewParsingXMLFailErr(err)
	}
	defer d.Free()

	if err = s.Validate(d); err != nil {
		v.lg.Error("validation failed", zap.Error(err))
		return errors.NewValidationXMLFailErr(err, err.(xsd.SchemaValidationError).Errors())
	}
	return nil
}
