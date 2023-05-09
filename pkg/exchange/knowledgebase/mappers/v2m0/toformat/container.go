package toformat

import (
	uuidGen "github.com/google/uuid"
	"go.uber.org/zap"

	kbEnt "github.com/anondigriz/mogan-core/pkg/entities/containers/knowledgebase"
	errMsgs "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/errors/messages"
	"github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats"
	formatV2M0 "github.com/anondigriz/mogan-core/pkg/exchange/knowledgebase/formats/v2m0"
)

func (tf *ToFormat) processContainer(model *formatV2M0.Model) error {
	tf.prepareModel(model)
	tf.warmUpWorkspace()

	err := tf.processPatterns(processPatternsArgs{
		patterns: tf.cont.Patterns,
		model:    model,
	})
	if err != nil {
		tf.lg.Error(errMsgs.MappingPatternFail, zap.Error(err))
		return err
	}

	err = tf.processGroups(processGroupsArgs{
		childGroups: tf.cont.Groups,
		parentClass: &model.Class,
	})

	if err != nil {
		tf.lg.Error(errMsgs.MappingGroupsFail, zap.Error(err))
		return err
	}

	return tf.processRootClass(model)
}

func (tf *ToFormat) prepareModel(model *formatV2M0.Model) {
	model.BaseInfo = formatV2M0.BaseInfo{
		ID:          tf.cont.KnowledgeBase.BaseInfo.ID,
		ShortName:   tf.cont.KnowledgeBase.BaseInfo.ShortName,
		Description: tf.cont.KnowledgeBase.BaseInfo.Description,
	}
	model.FormatXMLVersion = string(formats.VersionV2M0)
	model.Relations.Relations = []formatV2M0.Relation{}

	model.Class = formatV2M0.Class{
		BaseInfo: formatV2M0.BaseInfo{
			ID:          uuidGen.NewString(),
			ShortName:   tf.cont.KnowledgeBase.ShortName,
			Description: tf.cont.KnowledgeBase.Description,
		},
	}
}

func (tf *ToFormat) warmUpWorkspace() {
	unprocessedParameters := []string{}
	for k := range tf.cont.Parameters {
		unprocessedParameters = append(unprocessedParameters, k)
	}
	tf.ws.SaveUnprocessedParameters(unprocessedParameters)
}

func (tf *ToFormat) processRootClass(model *formatV2M0.Model) error {
	rootGroup := kbEnt.Group{}

	rootGroup.Parameters = append(rootGroup.Parameters, tf.ws.GetUnprocessedParameters()...)
	for k := range tf.cont.Rules {
		rootGroup.Rules = append(rootGroup.Rules, k)
	}

	err := tf.processRules(processRulesArgs{
		parentGroup: rootGroup,
		parentClass: &model.Class,
	})
	if err != nil {
		tf.lg.Error(errMsgs.MappingRulesFail, zap.Error(err))
		return err
	}

	err = tf.processParameters(processParametersArgs{
		parentGroup: rootGroup,
		parentClass: &model.Class,
	})
	if err != nil {
		tf.lg.Error(errMsgs.MappingParametersFail, zap.Error(err))
		return err
	}

	return nil
}
