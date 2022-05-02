package main

import (
	"context"
	"fmt"
	"time"

	sqladmin "google.golang.org/api/sqladmin/v1beta4"
)

func CloneSQLInstance(project, instance string) (*sqladmin.Operation, error) {
	ctx := context.Background()

	sqladminService, err := sqladmin.NewService(ctx)
	if err != nil {
		return nil, err
	}
	/* clc := &sqladmin.CloneContext{
		DestinationInstanceName: fmt.Sprintf("%s-clone", instance),
	} */
	rb := &sqladmin.InstancesCloneRequest{
		CloneContext: &sqladmin.CloneContext{
			DestinationInstanceName: fmt.Sprintf("%s-clone-%s", instance, time.Now().Format("01-02-2006")),
		},
	}

	resp, err := sqladminService.Instances.Clone(project, instance, rb).Context(ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ListSQLInstances(projectId string) ([]*sqladmin.DatabaseInstance, error) {
	ctx := context.Background()

	sqladminService, err := sqladmin.NewService(ctx)
	if err != nil {
		return nil, err
	}

	// List instances for the project ID.
	instances, err := sqladminService.Instances.List(projectId).Do()
	if err != nil {
		return nil, err
	}
	return instances.Items, nil
}
