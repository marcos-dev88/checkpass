package entity

func uint32ToPointer(v uint32) *uint32 {
	return &v
}

var successRules []RuleType = []RuleType{
	{
		Rule:  "minSize",
		Value: uint32ToPointer(8),
	},
	{
		Rule:  "minSpecialChars",
		Value: uint32ToPointer(8),
	},
	{
		Rule:  "noRepeted",
		Value: uint32ToPointer(0),
	},
	{
		Rule:  "minDigit",
		Value: uint32ToPointer(8),
	},
	{
		Rule:  "minUppercase",
		Value: uint32ToPointer(8),
	},
	{
		Rule:  "minLowercase",
		Value: uint32ToPointer(8),
	},
}

var invalidRule []RuleType = []RuleType{
	{
		Rule:  "minSize",
		Value: uint32ToPointer(8),
	},
	{
		Rule:  "minnSpecialChars",
		Value: uint32ToPointer(8),
	},
	{
		Rule:  "noRepeted",
		Value: uint32ToPointer(0),
	},
	{
		Rule:  "minDigit",
		Value: uint32ToPointer(8),
	},
}
