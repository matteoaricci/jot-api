package dictionary

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

var BaseUrl = "https://api.dictionaryapi.dev/api/v2/entries/en/"

// ErrDefinitionNotFound means the word has no dictionary entry (a 404). Callers
// can retry with a different word; other errors indicate an API failure.
var ErrDefinitionNotFound = errors.New("definition not found")

var httpClient = &http.Client{Timeout: 5 * time.Second}

// Definition is the clean domain type returned to callers.
type Definition struct {
	Word         string
	PartOfSpeech string
	Meaning      string
	Example      string
}

// dictionaryResponse mirrors the API's shape (a JSON array of entries). Private
// to this package — callers only ever see Definition.
type dictionaryResponse struct {
	Word     string `json:"word"`
	Meanings []struct {
		PartOfSpeech string `json:"partOfSpeech"`
		Definitions  []struct {
			Definition string `json:"definition"`
			Example    string `json:"example"`
		} `json:"definitions"`
	} `json:"meanings"`
}

func GetWordDefinition(ctx context.Context, word string) (Definition, error) {
	slog.InfoContext(ctx, "dictionary.GetWordDefinition", slog.String("word", word))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, BaseUrl+word, nil)
	if err != nil {
		return Definition{}, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		slog.WarnContext(ctx, "dictionary.GetWordDefinition: request failed", slog.String("word", word), slog.String("error", err.Error()))
		return Definition{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		slog.InfoContext(ctx, "dictionary.GetWordDefinition: not found", slog.String("word", word))
		return Definition{}, ErrDefinitionNotFound
	}
	if resp.StatusCode != http.StatusOK {
		slog.WarnContext(ctx, "dictionary.GetWordDefinition: unexpected status", slog.String("word", word), slog.Int("status", resp.StatusCode))
		return Definition{}, fmt.Errorf("dictionary api returned status %d for %q", resp.StatusCode, word)
	}

	var body []dictionaryResponse
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return Definition{}, err
	}

	if len(body) == 0 {
		slog.InfoContext(ctx, "dictionary.GetWordDefinition: not found", slog.String("word", word))
		return Definition{}, ErrDefinitionNotFound
	}
	entry := body[0]
	if len(entry.Meanings) == 0 || len(entry.Meanings[0].Definitions) == 0 {
		slog.InfoContext(ctx, "dictionary.GetWordDefinition: not found", slog.String("word", word))
		return Definition{}, ErrDefinitionNotFound
	}

	meaning := entry.Meanings[0]
	def := meaning.Definitions[0]

	return Definition{
		Word:         entry.Word,
		PartOfSpeech: meaning.PartOfSpeech,
		Meaning:      def.Definition,
		Example:      def.Example,
	}, nil
}
