package messages

const (
	ImportKnowledgeBaseFromXMLFail = "error when importing knowledge base from the XML file"

	FormatXMLVersionNotFound       = "xml exchange document file version could not be detected"
	UnsupportedFormatXMLVersion    = "xml exchange document file version is not supported"
	UnexpectedJobExecution         = "unexpected job execution error occurred"
	RelationNotFoundForRule        = "relation was not found for the rule"

	UnknownPatternTypeInXML        = "unknown pattern type found in the XML file"
	UnknownParameterTypeInXML      = "unknown parameter type found in the XML file"
	UnknownScriptLanguageTypeInXML = "unknown script language type found in the XML file"
	
	BrokenDictionaryInXMLAttribute = "dictionary in the xml attribute is broken"

	ParsingRelationsFromXMLFail          = "error when parsing relations from the XML file"
	ParsingParametersFromXMLFail         = "error when parsing parameters from the XML file"
	ParsingRelationParametersFromXMLFail = "error when parsing the relation's parameters from the XML file"
	ParsingRuleParametersFromXMLFail     = "error when parsing the rule's parameters from the XML file"
	ParsingRulesFromXMLFail              = "error when parsing rules from the XML file"
	ParsingConstraintsFromXMLFail        = "error when parsing constraints from the XML file"
	ParsingClassFromXMLFail              = "error when parsing class from the XML file"
	ParsingGroupFromXMLFail              = "error when parsing group from the XML file"
	ParsingPatternFromXMLFail            = "error when parsing pattern from the XML file"
	ParsingParameterFromXMLFail          = "error when parsing parameter from the XML file"
	ParsingClassesFromXMLFail            = "error when parsing classes from the XML file"

	ExtractPatternTypeFail         = "pattern type could not be determined"
	ExtractScriptLanguageTypeInXML = "script language type could not be determined"
	ExtractParameterTypeFail       = "parameter type could not be determined"
)
