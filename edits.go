package goopenai

import (
	"context"
	"encoding/json"
)

const EDITS_URL = "https://api.openai.com/v1/edits"

type CreateEditsRequest struct {
	Model       string  `json:"model,omitempty"`
	Input       string  `json:"input,omitempty"`
	Instruction string  `json:"instruction,omitempty"`
	N           int     `json:"n,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
	TopP        float64 `json:"top_p,omitempty"`
}

type CreateEditsResponse struct {
	Object  string  `json:"object,omitempty"`
	Created int     `json:"created,omitempty"`
	Choices Choices `json:"choices,omitempty"`
	Usage   Usage   `json:"usage,omitempty"`
	Error   Error   `json:"error,omitempty"`
}

func (c *Client) CreateEditsRaw(ctx context.Context, r CreateEditsRequest) ([]byte, error) {
	return c.Post(ctx, EDITS_URL, r)
}

func (c *Client) CreateEdits(ctx context.Context, r CreateEditsRequest) (response CreateEditsResponse, err error) {
	raw, err := c.CreateEditsRaw(ctx, r)
	if err != nil {
		return response, err
	}

	err = json.Unmarshal(raw, &response)
	return response, err
}
