# Build the manager binary
FROM golang:1.22 AS builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY api/ api/
COPY controllers/ controllers/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o manager main.go

FROM registry.access.redhat.com/ubi9/ubi-micro
WORKDIR /
COPY --from=builder /workspace/manager .
COPY  LICENSE /licenses/LICENSE

LABEL \
    name="simple-demo-operator" \
    License="Apache-2.0" \
    summary="This is the controller for the simple-demo-operator" \
    maintainer="opdev" \
    vendor="Partner Engineering" \
    release="1" \
    version="0.0.8" \
    description="A simple Kubernetes operator to be used for testing purposes"

USER 65532:65532

ENTRYPOINT ["/manager"]
