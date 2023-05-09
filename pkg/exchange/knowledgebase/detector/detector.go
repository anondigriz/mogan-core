package detector

import (
	"bufio"
	"regexp"

	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	"go.uber.org/zap"
)

var DefaultMaxCheckLines int = 15

type Detector struct {
	lg            *zap.Logger
	maxCheckLines int
}

func New(lg *zap.Logger) *Detector {
	d := &Detector{
		lg:            lg,
		maxCheckLines: DefaultMaxCheckLines,
	}

	return d
}

// SetMaxCheckLines - set maximum number of lines for checks
func (v *Detector) SetMaxCheckLines(number int) {
	v.maxCheckLines = number
}

func (v *Detector) DetectVersion(scanner *bufio.Scanner) (string, error) {
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
