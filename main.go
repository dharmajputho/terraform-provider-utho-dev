package main

import (
	"context"
	"log"

	"github.com/dharmajputho/terraform-provider-utho/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	err := providerserver.Serve(
		context.Background(),
		provider.New(),
		providerserver.ServeOpts{
			Address: "registry.terraform.io/utho-cloud/utho",
		},
	)
	if err != nil {
		log.Fatal(err)
	}
}
