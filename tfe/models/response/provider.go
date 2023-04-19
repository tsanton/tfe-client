package response

import (
	"time"

	me "github.com/tsanton/tfe-client/tfe/models/enum"
)

type Provider struct {
	Data struct {
		Id         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name         string          `json:"name"`
			Namespace    string          `json:"namespace"`
			RegistryName me.RegistryType `json:"registry-name"`
			Created      time.Time       `json:"created-at"`
			Updated      time.Time       `json:"updated-at"`
			Permissions  struct {
				Name string `json:"type"`
			} `json:"permissions"`
		} `json:"attributes"`
	} `json:"data"`

	Relationships struct {
		Organization struct {
			Data struct {
				Id   string `json:"id"`
				Type string `json:"type"`
			} `json:"data"`
		} `json:"organization"`
		Versions struct {
			Links struct {
				Related string `json:"related"`
			} `json:"links"`
		} `json:"versions"`
	} `json:"relationships"`

	Links struct {
		Self string `json:"self"`
	} `json:"links"`
}
