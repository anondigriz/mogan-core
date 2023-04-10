package errors

import "strconv"

const (
	ValidationXMLFail        = "ValidationXMLFail"
	OpeningXSDFail           = "OpeningXSDFail"
	ReadingXMLFail           = "ReadingXMLFail"
	ReadingXSDFail           = "ReadingXSDFail"
	ParsingXMLFail           = "ParsingXMLFail"
	ParsingXSDFail           = "ParsingXSDFail"
	FormatXMLVersionNotFound = "FormatXMLVersionNotFound"
)

type ValidatorErr struct {
	Stat    string
	Message string
	Err     error
	Dt      map[string]string
}

func (er ValidatorErr) Status() string {
	return er.Stat
}

func (er ValidatorErr) Error() string {
	return er.Message
}

func (er ValidatorErr) Data() map[string]string {
	return er.Dt
}

func (er ValidatorErr) Unwrap() error {
	return er.Err
}

func (er ValidatorErr) IsValidatorErr() bool {
	return true
}

func NewValidationXMLFailErr(baseError error, errors []error) error {
	e := ValidatorErr{
		Stat:    ValidationXMLFail,
		Message: "xml document validation error",
		Err:     baseError,
		Dt:      map[string]string{},
	}

	for i, err := range errors {
		e.Dt[strconv.Itoa(i)] = err.Error()
	}

	return e
}

func NewOpeningXSDFailErr(e error) error {
	return ValidatorErr{
		Stat:    OpeningXSDFail,
		Message: "error opening XSD file",
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewReadingXMLFailErr(e error) error {
	return ValidatorErr{
		Stat:    ReadingXMLFail,
		Message: "error reading XML file",
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewReadingXSDFailErr(e error) error {
	return ValidatorErr{
		Stat:    ReadingXSDFail,
		Message: "error reading XSD file",
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewParsingXMLFailErr(e error) error {
	return ValidatorErr{
		Stat:    ParsingXMLFail,
		Message: "error parsing XML file",
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewParsingXSDFailErr(e error) error {
	return ValidatorErr{
		Stat:    ParsingXSDFail,
		Message: "error parsing XSD file",
		Err:     e,
		Dt:      map[string]string{},
	}
}

func NewFormatXMLVersionNotFoundErr() error {
	return ValidatorErr{
		Stat:    FormatXMLVersionNotFound,
		Message: "xml exchange document file version could not be detected",
		Err:     nil,
		Dt:      map[string]string{},
	}
}
