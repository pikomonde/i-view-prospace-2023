package parser_test

import (
	"testing"

	servPars "github.com/pikomonde/i-view-prospace-2023/service/parser"
	"github.com/stretchr/testify/assert"
)

const defaultUnexpectedInput = "I have no idea what you are talking about"

func testCases() [][2]string {
	return [][2]string{
		{"glob is I", ""},
		{"prok is V", ""},
		{"pish is X", ""},
		{"tegj is L", ""},
		{"glob glob Silver is 34 Credits", ""},
		{"glob prok Gold is 57800 Credits", ""},
		{"pish pish Iron is 3910 Credits", ""},
		{"how much is pish tegj glob glob ?", "pish tegj glob glob is 42"},
		{"how many Credits is glob prok Silver ?", "glob prok Silver is 68 Credits"},
		{"how many Credits is glob prok Gold ?", "glob prok Gold is 57800 Credits"},
		{"how many Credits is glob prok Iron ?", "glob prok Iron is 782 Credits"},
		{"how much wood could a woodchuck chuck if a woodchuck could chuck wood ?", defaultUnexpectedInput},
		{"ccc is #", defaultUnexpectedInput},
		{"lolo kkkk Porridge is 123 Credits", defaultUnexpectedInput},
		{"glob Porridge is ddd Credits", defaultUnexpectedInput},
		{"how much is eee fff ggg ?", defaultUnexpectedInput},
		{"Iron is 3910 Credits", defaultUnexpectedInput},
		{"how many Credits is eee fff ggg Porridge ?", defaultUnexpectedInput},
		{"how many Credits is glob Porridge ?", defaultUnexpectedInput},
		{"aaa bbb is M", defaultUnexpectedInput},
		{"Does pish tegj glob glob Iron has more Credits than glob glob Gold ?", "no"},
		{"Does pish tegj glob glob Iron has less Credits than glob glob Gold ?", "yes"},
		{"Does glob glob Gold has less Credits than pish tegj glob glob Iron ?", "no"},
		{"Does glob glob Gold has more Credits than pish tegj glob glob Iron ?", "yes"},
		{"Does glob glob Gold has xxxx Credits than pish tegj glob glob Iron ?", defaultUnexpectedInput},
		{"Does xxxx Gold has more Credits than pish tegj glob glob Iron ?", defaultUnexpectedInput},
		{"Does glob glob Gold has more Credits than xxxx Iron ?", defaultUnexpectedInput},
		{"Does glob glob zzzz has more Credits than pish tegj glob glob Iron ?", defaultUnexpectedInput},
		{"Does glob glob Gold has more Credits than pish tegj glob glob zzzz ?", defaultUnexpectedInput},
		{"Is glob prok larger than pish pish ?", "no"},
		{"Is glob prok smaller than pish pish ?", "yes"},
		{"Is tegj glob glob smaller than glob prok ?", "no"},
		{"Is tegj glob glob larger than glob prok ?", "yes"},
		{"Is tegj glob glob xxxx than glob prok ?", defaultUnexpectedInput},
		{"Is xxxx larger than glob prok ?", defaultUnexpectedInput},
		{"Is tegj glob glob larger than xxxx ?", defaultUnexpectedInput},
	}
}

func TestParse(t *testing.T) {
	sPars := servPars.New()
	for _, v := range testCases() {
		assert.Equal(t, v[1], sPars.Parse(v[0]))
	}
}
