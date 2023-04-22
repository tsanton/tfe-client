package response

import "time"

type ProviderVersion struct {
	Data providerVersionData `json:"data"`
}

type providerVersionData struct {
	Id            string                           `json:"id"`
	Type          string                           `json:"type"`
	Attributes    providerVersionDataAttributes    `json:"attributes"`
	Relationships providerVersionDataRelationships `json:"relationships"`
	Links         providerVersionDataLinks         `json:"links"`
}

type providerVersionDataAttributes struct {
	Version            string                                   `json:"version"`
	CreatedAt          time.Time                                `json:"created-at"`
	UpdatedAt          time.Time                                `json:"updated-at"`
	KeyId              string                                   `json:"key-id"`
	Protocols          []string                                 `json:"protocols"`
	Permissions        providerVersionDataAttributesPermissions `json:"permissions"`
	ShasumsUploaded    bool                                     `json:"shasums-uploaded"`
	ShasumsSigUploaded bool                                     `json:"shasums-sig-uploaded"`
}

type providerVersionDataAttributesPermissions struct {
	CanDelete      bool `json:"can-delete"`
	CanUploadAsset bool `json:"can-upload-asset"`
}

type providerVersionDataRelationships struct {
	RegistryProvider providerVersionDataRelationshipsResistryProvider `json:"registry-provider"`
	Platforms        providerVersionDataRelationshipsPlatforms        `json:"platforms"`
}

type providerVersionDataRelationshipsResistryProvider struct {
	Data providerVersionDataRelationshipsResistryProviderData `json:"data"`
}

type providerVersionDataRelationshipsResistryProviderData struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type providerVersionDataRelationshipsPlatforms struct {
	Data  []providerVersionDataRelationshipsPlatformsData `json:"data"`
	Links providerVersionDataRelationshipsPlatformsLinks  `json:"links"`
}

type providerVersionDataRelationshipsPlatformsData struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type providerVersionDataRelationshipsPlatformsLinks struct {
	Related string `json:"related"`
}

// The shasums-uploaded and shasums-sig-uploaded properties will be false if those files have not been uploaded to Archivist.
// In this case, instead of including links to shasums-download and shasums-sig-download, the response will include upload links
type providerVersionDataLinks struct {
	ShasumsUploadUrl    string `json:"shasums-upload"`
	ShasumsSigUploadUrl string `json:"shasums-sig-upload"`

	ShasumsDownloadUrl    string `json:"shasums-download"`
	ShasumsSigDownloadUrl string `json:"shasums-sig-download"`
}
