#!/bin/bash
set -e

# Unless explicitly stated otherwise all files in this repository are licensed
# under the Apache License Version 2.0.
# This product includes software developed at Datadog (https://www.datadoghq.com/).
# Copyright 2016-present Datadog, Inc.

/opt/datadog-agent/bin/datadog-cluster-agent secret-helper read --with-provider-prefixes
