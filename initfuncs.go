package main

import (
	"context"
	"log"

	iam "cloud.google.com/go/iam/admin/apiv1"
	"github.com/pkg/errors"
	iampb "google.golang.org/genproto/googleapis/iam/v1"
)

func testPermisions(parent string) {
	ctx := context.Background()
	iamCl, err := iam.NewIamClient(ctx)
	scopes := []string{
		"cloudsql.instances.clone",
		"cloudasset.assets.listResource",
	}
	p, err := iamCl.TestIamPermissions(ctx, &iampb.TestIamPermissionsRequest{
		Resource:    parent,
		Permissions: scopes,
	})
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to test iam permissions"))
	}
	log.Println(p.GetPermissions())
}
