package nlp

import (
	"log"
	"strings"
	"sync"

	"borscht.app/kapusta/model"
)

// Parse is the main parser for a given recipe.
// It looks for the following
// - Contains number
// - Contains mass/volume
// - Contains ingredients
// - Number occurs before ingredients
// - Number occurs before mass/volume
// - Number of ingredients is 1
// - Percent of other words is less than 50%
// - Part of list (contains - or *)
func Parse(p *model.InputData, r *model.Recipe) error {
	if len(p.Text) == 0 {
		return nil
	}

	lines := strings.Split(p.Text, "\n")
	scores := make([]float64, len(lines))
	lineInfos := make([]LineInfo, len(lines))

	i := -1
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}
		i++
		lineInfos[i].LineOriginal = line
		line = SanitizeLine(line)
		lineInfos[i].Line = line
		lineInfos[i].IngredientsInString = GetIngredientsInString(line)
		lineInfos[i].AmountInString = GetNumbersInString(line)
		lineInfos[i].MeasureInString = GetMeasuresInString(line)

		score := 0.0
		// does it contain an ingredients?
		if len(lineInfos[i].IngredientsInString) > 0 {
			score++
		}
		// does it contain an amount?
		if len(lineInfos[i].AmountInString) > 0 {
			score++
		}
		// does it contain a measure (cups, tsps)?
		if len(lineInfos[i].MeasureInString) > 0 {
			score++
		}
		// does the ingredients come after the measure?
		if len(lineInfos[i].IngredientsInString) > 0 && len(lineInfos[i].MeasureInString) > 0 && lineInfos[i].IngredientsInString[0].Position > lineInfos[i].MeasureInString[0].Position {
			score++
		}
		// does the ingredients come after the amount?
		if len(lineInfos[i].IngredientsInString) > 0 && len(lineInfos[i].AmountInString) > 0 && lineInfos[i].IngredientsInString[0].Position > lineInfos[i].AmountInString[0].Position {
			score++
		}
		// does the measure come after the amount?
		if len(lineInfos[i].MeasureInString) > 0 && len(lineInfos[i].AmountInString) > 0 && lineInfos[i].MeasureInString[0].Position > lineInfos[i].AmountInString[0].Position {
			score++
		}
		// is the line really long? (ingredients lines are short)
		if score > 0 && len(lineInfos[i].LineOriginal) > 100 {
			score--
		}
		// does it start with a list indicator (* or -)?
		fields := strings.Fields(line)
		if len(fields) > 0 && (fields[0] == "*" || fields[0] == "-") {
			score++
		}
		// if only one thing is right, its wrong
		if score == 1 {
			score = 0.0
		}
		// log.Printf("'%s' (%d)", line, score)
		scores[i] = score
	}
	scores = scores[:i+1]
	lineInfos = lineInfos[:i+1]

	// debugging purposes
	// lines = make([]string, len(lineInfos))
	// for i, li := range lineInfos {
	// 	lines[i] = li.Line
	// }
	// ioutil.WriteFile("out", []byte(strings.Join(lines, "\n")), 0644)

	// get the most likely location
	start, end := getBestTopHatPositions(scores)

	var wg sync.WaitGroup
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		err := parseDirections(r, lineInfos[end:])
		if err != nil {
			log.Print(err)
		}
	}(&wg)

	ingredientLines := []LineInfo{}
	if start-3 > 0 && end+3 < len(lineInfos) {
		for _, lineInfo := range lineInfos[start-3 : end+3] {
			if len(strings.TrimSpace(lineInfo.Line)) < 3 {
				continue
			}

			// get amount, continue if there is an error
			err := lineInfo.getTotalAmount()
			if err != nil {
				log.Print(err)
				continue
			}

			// get ingredients, continue if its not found
			err = lineInfo.getIngredient()
			if err != nil {
				log.Print(err)
				continue
			}

			ingredientLines = append(ingredientLines, lineInfo)
		}
	}

	// consolidate ingredients
	for _, line := range ingredientLines {
		r.Ingredients = append(r.Ingredients, line.Line)
	}

	wg.Done()
	wg.Wait()

	return nil
}

func parseDirections(r *model.Recipe, lis []LineInfo) (rerr error) {
	log.Print(len(lis))
	scores := make([]float64, len(lis))
	for i, li := range lis {
		if i > 30 {
			break
		}
		if len(strings.TrimSpace(li.Line)) < 3 {
			continue
		}
		score := 0.0
		for _, corpusDirection := range CorpusDirections {
			if strings.Contains(li.Line, corpusDirection) {
				score++
			}
		}
		for _, corpusDirection := range CorpusDirectionsNeg {
			if strings.Contains(li.Line, corpusDirection) {
				score--
			}
		}
		if len(li.Line) < 5 {
			score = 0
		}
		scores[i] = score
	}

	start, end := getBestTopHatPositions(scores)
	log.Printf("direction are from line %d to %d", start, end)
	directionI := 1

	for i := start; i <= end; i++ {
		if len(strings.TrimSpace(lis[i].Line)) == 0 {
			continue
		}

		direction := strings.TrimSpace(lis[i].LineOriginal)
		if string(direction[0]) == "*" {
			direction = strings.TrimSpace(direction[1:])
		}

		if len(strings.Fields(direction)) < 5 {
			continue
		}
		log.Printf("%d) %s", directionI, direction)
		directionI++
		r.Instructions = append(r.Instructions, &model.Instruction{HowToStep: model.HowToStep{Text: direction}})
	}
	return
}
