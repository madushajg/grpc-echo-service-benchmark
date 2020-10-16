# -------------------------------------------------------------------------------------
#
# Copyright (c) 2020, WSO2 Inc. (http://www.wso2.com). All Rights Reserved.
#
# This software is the property of WSO2 Inc. and its suppliers, if any.
# Dissemination of any information or reproduction of any material contained
# herein in any form is strictly forbidden, unless permitted by WSO2 expressly.
# You may not alter or remove any copyright or other notice from copies of this content.
#
# --------------------------------------------------------------------------------------

ARG BUILDER_WORKSPACE="/workspace"
FROM golang:1.13.5-alpine3.10 AS builder

ARG BUILDER_WORKSPACE
WORKDIR ${BUILDER_WORKSPACE}

ADD cmd cmd
ADD pkg pkg
ADD go.mod go.mod
ADD go.sum go.sum

RUN CGO_ENABLED=0 GOOS=linux go build -o grpcperf -x ./cmd/

FROM ubuntu:18.04

ARG BUILDER_WORKSPACE

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/* \
    && apt-get clean

COPY --from=builder ${BUILDER_WORKSPACE}/grpcperf ./

CMD /bin/sh -c ./grpcperf
