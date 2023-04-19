package response

import "time"

type ProviderVersion struct {
	Data ProviderVersionData `json:"data"`
}

type ProviderVersionData struct {
	Id            string                           `json:"id"`
	Type          string                           `json:"type"`
	Attributes    ProviderVersionDataAttributes    `json:"attributes"`
	Relationships ProviderVersionDataRelationships `json:"relationships"`
	Links         ProviderVersionDataLinks         `json:"links"`
}

type ProviderVersionDataAttributes struct {
	Version            string                                   `json:"version"`
	CreatedAt          time.Time                                `json:"created-at"`
	UpdatedAt          time.Time                                `json:"updated-at"`
	KeyId              string                                   `json:"key-id"`
	Protocols          []string                                 `json:"protocols"`
	Permissions        ProviderVersionDataAttributesPermissions `json:"permissions"`
	ShasumsUploaded    bool                                     `json:"shasums-uploaded"`
	ShasumsSigUploaded bool                                     `json:"shasums-sig-uploaded"`
}

type ProviderVersionDataAttributesPermissions struct {
	CanDelete      bool `json:"can-delete"`
	CanUploadAsset bool `json:"can-upload-asset"`
}

type ProviderVersionDataRelationships struct {
	RegistryProvider ProviderVersionDataRelationshipsResistryProvider `json:"registry-provider"`
	Platforms        ProviderVersionDataRelationshipsPlatforms        `json:"platforms"`
}

type ProviderVersionDataRelationshipsResistryProvider struct {
	Data ProviderVersionDataRelationshipsResistryProviderData `json:"data"`
}

type ProviderVersionDataRelationshipsResistryProviderData struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type ProviderVersionDataRelationshipsPlatforms struct {
	Data  []ProviderVersionDataRelationshipsPlatformsData `json:"data"`
	Links ProviderVersionDataRelationshipsPlatformsLinks  `json:"links"`
}

type ProviderVersionDataRelationshipsPlatformsData struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type ProviderVersionDataRelationshipsPlatformsLinks struct {
	Related string `json:"related"`
}

// The shasums-uploaded and shasums-sig-uploaded properties will be false if those files have not been uploaded to Archivist.
// In this case, instead of including links to shasums-download and shasums-sig-download, the response will include upload links
type ProviderVersionDataLinks struct {
	ShasumsUploadUrl    string `json:"shasums-upload"`
	ShasumsSigUploadUrl string `json:"shasums-sig-upload"`

	ShasumsDownloadUrl    string `json:"shasums-download"`
	ShasumsSigDownloadUrl string `json:"shasums-sig-download"`
}
