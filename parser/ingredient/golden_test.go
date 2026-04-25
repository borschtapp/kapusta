package ingredient

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/borschtapp/kapusta/model"
	"github.com/stretchr/testify/assert"
)

func TestParseIngredientsTestdataSnapshot(t *testing.T) {
	file, err := os.Open("testdata/ingredients.txt")
	if err != nil {
		t.Fatalf("failed to open test data file: %v", err)
	}
	defer func() { _ = file.Close() }()

	var results []model.Ingredient
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ing, err := ParseIngredient(line, Options{Lang: "en"})
		assert.NoError(t, err)
		results = append(results, ing)
	}
	if err := scanner.Err(); err != nil {
		t.Fatalf("scanner error: %v", err)
	}

	goldenFile := "testdata/ingredients.golden.json"

	if os.Getenv("UPDATE_GOLDEN") == "1" {
		data, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			t.Fatalf("failed to marshal results: %v", err)
		}
		err = os.WriteFile(goldenFile, data, 0644)
		if err != nil {
			t.Fatalf("failed to write golden file: %v", err)
		}
		t.Log("Updated golden file")
		return
	}

	goldenData, err := os.ReadFile(goldenFile)
	if err != nil {
		t.Fatalf("failed to read golden file (run with UPDATE_GOLDEN=1 to create): %v", err)
	}

	var expected []model.Ingredient
	if err := json.Unmarshal(goldenData, &expected); err != nil {
		t.Fatalf("failed to unmarshal golden file: %v", err)
	}

	assert.Equal(t, len(expected), len(results), "number of results changed")
	for i := range expected {
		assert.Equal(t, expected[i], results[i], "mismatch at line %d", i+1)
	}
}

// go test -bench=. -benchmem -count=6 ./... > benchmarks.txt
// go test -bench=. -benchmem -count=6 ./... > new.txt
// benchstat benchmarks.txt new.txt
func BenchmarkParseIngredientsTestdata(b *testing.B) {
	data, err := os.ReadFile("testdata/ingredients.txt")
	if err != nil {
		b.Fatalf("failed to open test data file: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, line := range lines {
			_, _ = ParseIngredient(line, Options{Lang: "en"})
		}
	}
}
