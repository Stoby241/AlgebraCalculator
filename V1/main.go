package V1

import "AlgebraCalculator/V1/log"

var initilized bool

func init() {
	if !initilized {
		log.InitLog()
		initTerm()

		initParser()
		initSimplifyer()
		initTermFunctions()
		initSolve()
		initVector()
		initSort()
		initReplace()
		log.GetLog()

		initilized = true
	}
}

type Result struct {
	AnswerStrings map[int]string
	TermStrings   map[int]string
	Log           []string
}

var terms []*term

func Calculate(termStrings ...string) Result {
	terms = nil

	result := Result{
		AnswerStrings: map[int]string{},
		TermStrings:   map[int]string{},
	}

	for i, termString := range termStrings {
		if termString == "" {
			continue
		}

		term, err := parseTerm(termString)
		result.Log = append(result.Log, log.GetLog())
		if handelError(err) {
			result.Log = append(result.Log, log.GetLog())
			continue
		}

		simplifyRoot(term.root)
		if r := recover(); r != nil {
			handelError(newError(errorTypSolving, errorCriticalLevelPartial, "Some Error accured!"))
			result.Log = append(result.Log, log.GetLog())
			continue
		}
		result.Log = append(result.Log, log.GetLog())

		term.root.print()
		result.AnswerStrings[i] = log.GetLog()

		term.print()
		result.TermStrings[i] = log.GetLog()

		terms = append(terms, term)
	}
	return result
}
