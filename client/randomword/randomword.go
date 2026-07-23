package randomword

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var ApiUrl = "https://random-word-api.herokuapp.com/word"

var httpClient = &http.Client{Timeout: 5 * time.Second}

// GetRandomWords returns `number` random words. difficulty (1-5) filters by
// commonality; note the API only honors diff when requesting 5 or fewer words.
func GetRandomWords(ctx context.Context, number int, difficulty int) ([]string, error) {
	slog.InfoContext(ctx, "randomword.GetRandomWords", slog.Int("number", number), slog.Int("difficulty", difficulty))

	q := url.Values{}
	q.Set("number", strconv.Itoa(number))
	q.Set("diff", strconv.Itoa(difficulty))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ApiUrl+"?"+q.Encode(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		slog.WarnContext(ctx, "randomword.GetRandomWords: request failed", slog.String("error", err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.WarnContext(ctx, "randomword.GetRandomWords: unexpected status", slog.Int("status", resp.StatusCode))
		return nil, fmt.Errorf("random word api returned status %d", resp.StatusCode)
	}

	var body []string
	if err = json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, err
	}

	if len(body) == 0 {
		slog.WarnContext(ctx, "randomword.GetRandomWords: empty response")
		return nil, fmt.Errorf("random word api returned no words")
	}

	return body, nil
}
