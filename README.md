# gRPC+Istio Demo in Go

This repo contains a basic gRPC service mesh written in Go, to be deployed in Istio-enabled Kubernetes cluster. The demo includes the following features:  

- gRPC service written in Go
- HTTP/gRPC transcoding using
  - Google Endpoints
  - Istio Envoy filter
- Istio routing using virtual services
- Traffic visability through Kiali (optional)

## Service mesh structure

The service mesh is composed of a server hosting the demo service, which calls a mock backend service. Multiple instances of the backend is running and virtual service routes the call to either one, with a distribution probability assigned in the YAML file.
