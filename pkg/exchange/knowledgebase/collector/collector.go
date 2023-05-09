package collector

import (
	"io"

	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/collector/v2m0"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/collector/v3m0"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/detector"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats"
)

type ParseXMLArgs struct {
	Version formats.Version
	Cont    *kbEnt.Container
	XMLFile io.Writer
}

type Collector struct {
	lg       *zap.Logger
	detector *detector.Detector
}

func New(lg *zap.Logger) *Collector {
	d := detector.New(lg)

	p := &Collector{
		lg:       lg,
		detector: d,
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
		tf := v2m0.New(c.lg)
		b, err = tf.CollectXML(args.Cont)
	default:
		tf := v3m0.New(c.lg)
		b, err = tf.CollectXML(args.Cont)
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
