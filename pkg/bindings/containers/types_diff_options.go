package containers

import (
	"net/url"

	"github.com/containers/podman/v2/pkg/bindings/util"
)

/*
This file is generated automatically by go generate.  Do not edit.
*/

// Changed
func (o *DiffOptions) Changed(fieldName string) bool {
	return util.Changed(o, fieldName)
}

// ToParams
func (o *DiffOptions) ToParams() (url.Values, error) {
	return util.ToParams(o)
}
