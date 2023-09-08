package parser

const (
	defaultAssignmentKeyword                     = "is"
	defaultComparisonKeyword                     = "than"
	defaultCreditsKeyword                        = "Credits"
	defaultQueryTransnumKeyword                  = "how much"
	defaultQueryPriceKeyword                     = "how many " + defaultCreditsKeyword
	defaultQueryPriceComparisonPrefixKeyword     = "Does"
	defaultQueryPriceComparisonMoreKeyword       = "has more " + defaultCreditsKeyword
	defaultQueryPriceComparisonLessKeyword       = "has less " + defaultCreditsKeyword
	defaultQueryTransnumComparisonPrefixKeyword  = "Is"
	defaultQueryTransnumComparisonLargerKeyword  = "larger"
	defaultQueryTransnumComparisonSmallerKeyword = "smaller"
	defaultQuestionMarkKeyword                   = "?"
	defaultUnexpectedInput                       = "I have no idea what you are talking about"
)

func space(s string) string {
	return " " + s + " "
}
