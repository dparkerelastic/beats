// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package panw

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "panw", asset.ModuleFieldsPri, AssetPanw); err != nil {
		panic(err)
	}
}

// AssetPanw returns asset data.
// This is the base64 encoded zlib format compressed contents of module/panw.
func AssetPanw() string {
	return "eJyszUEKwjAUhOF9TjF0nwu8nRfQgid4mlGCaRKSV0pvL0URceGqs/wX33g8uAqq5sUBFi1RMIyal8EBgf3aYrVYsmA8HP3pjKmEOdEBjYnaKbjQ1AG3yBS6OADwyDrx426ztVJwb2Wu7/JX3/YtvtTfz336MwAA//80X0nz"
}
