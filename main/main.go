package main

import (
	"github.com/operator-framework/operator-registry/pkg/client"
	"github.com/operator-framework/operator-registry/pkg/lib/bundle"
)

func main() {
	_, _ = client.NewClient("test")
	_ = bundle.RoleKind
}
