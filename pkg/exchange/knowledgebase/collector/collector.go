package collector

import (
	"fmt"
	"io"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/collector/v2m0"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/detector"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats"
)

type ParseXMLArgs struct {
	Version       formats.Version
	KnowledgeBase kbEnt.Container
	XMLFile       io.Writer
}

type Collector struct {
	lg       *zap.Logger
	detector *detector.Detector
	v2m0     *v2m0.V2M0
	//v3m0     *v3m0.V3M0
}

func New(lg *zap.Logger) *Collector {
	d := detector.New(lg)

	p := &Collector{
		lg:       lg,
		detector: d,
		v2m0:     v2m0.New(lg),
		//		v3m0:     v3m0.New(lg),
	}
	return p
}

func (c Collector) Collect(args ParseXMLArgs) error {
	switch args.Version {
	case formats.VersionV2M0:
		break
	case formats.VersionV3M0:
		break
	default:
		err := errors.NewUnsupportedFormatXMLVersionErr(string(args.Version))
		c.lg.Error(errMsgs.UnsupportedFormatXMLVersion, zap.Error(err))
		return err
	}

	return c.collectToFormat(args)
}

func (c Collector) collectToFormat(args ParseXMLArgs) error {
	var err error
	var b []byte

	switch args.Version {
	case formats.VersionV2M0:
		b, err = c.v2m0.CollectXML(args.KnowledgeBase)
	default:
		// b, err = c.v3m0.CollectXML(args.KnowledgeBase)
		return fmt.Errorf("not implemented")
	}

	if err != nil {
		if err != nil {
			c.lg.Error(errMsgs.CollectingXMLFail, zap.Error(err))
			return err
		}
	}

	_, err = args.XMLFile.Write(b)
	if err != nil {
		c.lg.Error(errMsgs.WritingXMLFail, zap.Error(err))
		return errors.NewWritingXMLFailErr(err)
	}

	return nil
}
