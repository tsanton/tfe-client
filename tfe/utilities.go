package tfe

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func MakeRequest[T any, U any](ctx context.Context, s *TerraformEnterpriseClient, method string, expectedResponseCode int, path string, body T) (*U, error) {
	b := []byte{}
	switch any(body).(type) {
	case nil:
		break
	default:
		var err error
		b, err = json.Marshal(body)
		// fmt.Print(string(b))
		if err != nil {
			return nil, fmt.Errorf("error serializing request body: %w", err)
		}
	}

	buf := bytes.NewBuffer(b)

	u, err := url.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("error parsing path: %w", err)
	}
	reqUrl := s.host.ResolveReference(u)
	req, err := http.NewRequestWithContext(ctx, method, reqUrl.String(), buf)
	if err != nil {
		return nil, err
	}
	if s.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))
	}
	req.Header.Set("Content-Type", "application/vnd.api+json")
	req.Header.Set("Accept", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != expectedResponseCode {
		// responseData, _ := io.ReadAll(resp.Body)
		// fmt.Print(string(responseData))
		return nil, fmt.Errorf("request returned non 200 response: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	var definitionResp U
	if resp.ContentLength == 0 {
		return nil, nil
	}

	err = json.NewDecoder(resp.Body).Decode(&definitionResp)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response: %w", err)
	}

	return &definitionResp, nil
}

func Do[T any](ctx context.Context, s *TerraformEnterpriseClient, expectedResponseCode int, req *http.Request) (*T, error) {
	if s.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.token))
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != expectedResponseCode {
		// responseData, _ := io.ReadAll(resp.Body)
		// fmt.Print(string(responseData))
		return nil, fmt.Errorf("request returned non 200 response: %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	var definitionResp T
	if resp.ContentLength == 0 {
		return nil, nil
	}

	err = json.NewDecoder(resp.Body).Decode(&definitionResp)
	if err != nil {
		return nil, fmt.Errorf("unable to decode response: %w", err)
	}

	return &definitionResp, nil
}
