package in

import "github.com/concourse/s3-resource"

type InRequest struct {
	Source  s3resource.Source  `json:"source"`
	Version s3resource.Version `json:"version"`
	Params  Params             `json:"params"`
}

type Params struct {
	Unpack bool `json:"unpack"`
}

type InResponse struct {
	Version  s3resource.Version        `json:"version"`
	Metadata []s3resource.MetadataPair `json:"metadata"`
}
