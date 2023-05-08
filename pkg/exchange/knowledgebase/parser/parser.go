// TODO обработка ошибок
package parser

import (
	"bufio"
	"context"
	"io"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/detector"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/parser/v2m0"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/parser/v3m0"
)

const (
	versionV2M0 = "2.0"
	versionV3M0 = "3.0"
)

type ParseXMLArgs struct {
	KnowledgeBaseUUID string
	XMLFile           XMLFile
	FileName          string
}

type XMLFile interface {
	io.Reader
	io.Seeker
	io.Closer
}

type Parser struct {
	lg       *zap.Logger
	detector *detector.Detector
	v2m0     *v2m0.V2M0
	v3m0     *v3m0.V3M0
}

func New(lg *zap.Logger) *Parser {
	d := detector.New(lg)

	p := &Parser{
		lg:       lg,
		detector: d,
		v2m0:     v2m0.New(lg),
		v3m0:     v3m0.New(lg),
	}
	return p
}

func (p Parser) Parse(ctx context.Context, args ParseXMLArgs) (kbEnt.Container, error) {
	scanner := bufio.NewScanner(args.XMLFile)
	ver, err := p.detector.DetectVersion(scanner)
	if err != nil {
		p.lg.Error("xml exchange document file version could not be detected", zap.Error(err))
		return kbEnt.Container{}, err
	}

	switch ver {
	case versionV2M0:
		break
	case versionV3M0:
		break
	default:
		err := errors.NewUnsupportedFormatXMLVersionErr(ver)
		p.lg.Error(errMsgs.UnsupportedFormatXMLVersion, zap.Error(err))
		return kbEnt.Container{}, err
	}

	err = p.seekFileToBegin(args)
	if err != nil {
		p.lg.Error("fail to seek file to the begin", zap.Error(err))
		return kbEnt.Container{}, err
	}

	return p.parseVersion(args, ver)
}

func (p Parser) seekFileToBegin(args ParseXMLArgs) error {
	_, err := args.XMLFile.Seek(0, 0)
	if err != nil {
		if err != nil {
			p.lg.Error("fail to reset the XML file reading stream to the beginning", zap.Error(err))
			return errors.NewReadingXMLFailErr(err)
		}
	}
	return nil
}

func (p Parser) parseVersion(args ParseXMLArgs, ver string) (kbEnt.Container, error) {
	content, err := io.ReadAll(args.XMLFile)
	if err != nil {
		if err != nil {
			p.lg.Error("fail to read the XML file from stream", zap.Error(err))
			return kbEnt.Container{}, errors.NewReadingXMLFailErr(err)
		}
	}
	var cont kbEnt.Container

	switch ver {
	case versionV2M0:
		cont, err = p.v2m0.ParseXML(args.KnowledgeBaseUUID, content)
	default:
		cont, err = p.v3m0.ParseXML(args.KnowledgeBaseUUID, content)
	}

	if err != nil {
		if err != nil {
			p.lg.Error("fail to parse the XML file", zap.Error(err))
			return kbEnt.Container{}, errors.NewParsingXMLFailErr("fail to parse the XML file", err)
		}
	}
	return cont, nil
}
