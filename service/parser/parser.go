package parser

import (
	"fmt"
	"strconv"
	"strings"
)

// Parse parsing user's input and returns system's output
func (s *ServiceParser) Parse(in string) string {

	inSplitByIs := strings.Split(in, space(defaultAssignmentKeyword))
	inSplitByThan := strings.Split(in, space(defaultComparisonKeyword))
	isSplitByIs := false
	isSplitByThan := false

	// validation
	if len(inSplitByIs) != 2 && len(inSplitByThan) != 2 {
		return defaultUnexpectedInput
	}
	leftIn := make([]string, 0)
	rightIn := make([]string, 0)

	// Getting left and right side of sentence
	if len(inSplitByIs) == 2 {
		isSplitByIs = true
		leftIn = strings.Split(inSplitByIs[0], " ")
		rightIn = strings.Split(inSplitByIs[1], " ")
	} else if len(inSplitByThan) == 2 {
		isSplitByThan = true
		leftIn = strings.Split(inSplitByThan[0], " ")
		rightIn = strings.Split(inSplitByThan[1], " ")
	}

	// Parse logic
	// TODO: better code writing? Maybe using regeluar expression?
	if isSplitByIs && (len(leftIn) == 1) && (len(rightIn) == 1) && (len(rightIn[0]) == 1) {
		// Add Galactic Unit to dictionary
		err := s.ServiceTransnum.AddGalacticUnit(leftIn[0], rune(rightIn[0][0]))
		if err != nil {
			return defaultUnexpectedInput
		}

		return ""
	} else if isSplitByIs && (len(leftIn) >= 1) && (len(rightIn) == 2) && (rightIn[len(rightIn)-1] == defaultCreditsKeyword) {
		// Translate Galactic unit to decimal
		unit, err := s.ServiceTransnum.GalaticToInt(leftIn[:len(leftIn)-1])
		if err != nil {
			return defaultUnexpectedInput
		}

		// Parse price from string to integer
		price, err := strconv.ParseInt(rightIn[0], 10, 64)
		if err != nil {
			return defaultUnexpectedInput
		}

		// Add price to dictionary
		err = s.ServiceResource.AddResourcePrice(leftIn[len(leftIn)-1], unit, int(price))
		if err != nil {
			return defaultUnexpectedInput
		}

		return ""
	} else if isSplitByIs && (strings.Join(leftIn, " ") == defaultQueryTransnumKeyword) && (rightIn[len(rightIn)-1] == defaultQuestionMarkKeyword) {
		words := rightIn[:len(rightIn)-1]

		// Translate galactic unit to decimal
		result, err := s.ServiceTransnum.GalaticToInt(words)
		if err != nil {
			return defaultUnexpectedInput
		}

		return strings.Join(append(words, []string{defaultAssignmentKeyword, fmt.Sprint(result)}...), " ")
	} else if isSplitByIs && (strings.Join(leftIn, " ") == defaultQueryPriceKeyword) && (len(rightIn) >= 3) && (rightIn[len(rightIn)-1] == defaultQuestionMarkKeyword) {
		rightIn = rightIn[:len(rightIn)-1]

		// Translate galactic unit to decimal
		unit, err := s.ServiceTransnum.GalaticToInt(rightIn[:len(rightIn)-1])
		if err != nil {
			return defaultUnexpectedInput
		}

		// Get and calculate total price of a resource
		totalPrice, err := s.ServiceResource.GetResourcePrice(unit, rightIn[len(rightIn)-1])
		if err != nil {
			return defaultUnexpectedInput
		}

		return strings.Join(append(rightIn, []string{defaultAssignmentKeyword, fmt.Sprint(totalPrice), defaultCreditsKeyword}...), " ")
	} else if isSplitByThan && (len(leftIn) >= 6) && (leftIn[0] == defaultQueryPriceComparisonPrefixKeyword) && (len(rightIn) >= 3) && (rightIn[len(rightIn)-1] == defaultQuestionMarkKeyword) {
		compareIn := leftIn[len(leftIn)-3:]
		leftIn = leftIn[1 : len(leftIn)-3]
		rightIn = rightIn[:len(rightIn)-1]

		// Translate galactic unit to decimal
		lUnit, err := s.ServiceTransnum.GalaticToInt(leftIn[:len(leftIn)-1])
		if err != nil {
			return defaultUnexpectedInput
		}

		// Get and calculate total price of a resource
		lTotalPrice, err := s.ServiceResource.GetResourcePrice(lUnit, leftIn[len(leftIn)-1])
		if err != nil {
			return defaultUnexpectedInput
		}

		// Translate galactic unit to decimal
		rUnit, err := s.ServiceTransnum.GalaticToInt(rightIn[:len(rightIn)-1])
		if err != nil {
			return defaultUnexpectedInput
		}

		// Get and calculate total price of a resource
		rTotalPrice, err := s.ServiceResource.GetResourcePrice(rUnit, rightIn[len(rightIn)-1])
		if err != nil {
			return defaultUnexpectedInput
		}

		if strings.Join(compareIn, " ") == defaultQueryPriceComparisonMoreKeyword {
			if lTotalPrice > rTotalPrice {
				return "yes"
			}
			return "no"
		} else if strings.Join(compareIn, " ") == defaultQueryPriceComparisonLessKeyword {
			if lTotalPrice < rTotalPrice {
				return "yes"
			}
			return "no"
		}

		return defaultUnexpectedInput
	} else if isSplitByThan && (len(leftIn) >= 3) && (leftIn[0] == defaultQueryTransnumComparisonPrefixKeyword) && (len(rightIn) >= 2) && (rightIn[len(rightIn)-1] == defaultQuestionMarkKeyword) {
		compareIn := leftIn[len(leftIn)-1]
		leftIn = leftIn[1 : len(leftIn)-1]
		rightIn = rightIn[:len(rightIn)-1]

		// Translate galactic unit to decimal
		lUnit, err := s.ServiceTransnum.GalaticToInt(leftIn)
		if err != nil {
			return defaultUnexpectedInput
		}

		// Translate galactic unit to decimal
		rUnit, err := s.ServiceTransnum.GalaticToInt(rightIn)
		if err != nil {
			return defaultUnexpectedInput
		}

		if compareIn == defaultQueryTransnumComparisonLargerKeyword {
			if lUnit > rUnit {
				return "yes"
			}
			return "no"
		} else if compareIn == defaultQueryTransnumComparisonSmallerKeyword {
			if lUnit < rUnit {
				return "yes"
			}
			return "no"
		}

		return defaultUnexpectedInput
	}

	return defaultUnexpectedInput
}
