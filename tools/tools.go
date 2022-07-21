//go:build tools

package tools

import (
	// Documentation generation
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"

	// Semantic version numbers based on git tags
	_ "github.com/restechnica/semverbot/cmd/sbot"
)
