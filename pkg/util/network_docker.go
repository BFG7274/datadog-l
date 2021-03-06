// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build docker
// +build docker

package util

import (
	"context"
	"fmt"

	"github.com/DataDog/datadog-agent/pkg/util/cache"
	"github.com/DataDog/datadog-agent/pkg/util/docker"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

// GetAgentNetworkMode retrieves from Docker the network mode of the Agent container
func GetAgentNetworkMode(ctx context.Context) (string, error) {
	cacheNetworkModeKey := cache.BuildAgentKey("networkMode")
	if cacheNetworkMode, found := cache.Cache.Get(cacheNetworkModeKey); found {
		return cacheNetworkMode.(string), nil
	}

	log.Debugf("GetAgentNetworkMode trying Docker")
	networkMode, err := docker.GetAgentContainerNetworkMode(ctx)
	cache.Cache.Set(cacheNetworkModeKey, networkMode, cache.NoExpiration)
	if err != nil {
		return networkMode, fmt.Errorf("could not detect agent network mode: %v", err)
	}
	log.Debugf("GetAgentNetworkMode: using network mode from Docker: %s", networkMode)
	return networkMode, nil
}
