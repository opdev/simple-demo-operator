# Simple Demo Operator

The simple demo operator is a working operator project with images
published to Quay.io.

- Container Image: quay.io/opdev/simple-demo-operator
- Bundle Image(s): quay.io/opdev/simple-demo-operator-bundle
- Catalog Image: quay.io/opdev/simple-demo-operator-catalog

## Release Workflow
Change `VERSION` value in Makefile

```bash
# Build and push the operator image with the new version
make docker-build
make docker-push

# Build and push the bundle image with the new version
make bundle
make bundle-build
make bundle push

# Build and push the catalog (from the old catalog) with the new bundle
CATALOG_BASE_IMG=quay.io/opdev/simple-demo-operator-catalog:v<PREVIOUS_VERSION_TAG> make catalog-build
make catalog-push

# Tag latest from the new latest tag
docker tag quay.io/opdev/simple-demo-operator:<NEW_VERSION> quay.io/opdev/simple-demo-operator:latest
docker push quay.io/opdev/simple-demo-operator:latest

docker tag quay.io/opdev/simple-demo-operator-bundle:v<NEW_VERSION> quay.io/opdev/simple-demo-operator-bundle:latest
docker push quay.io/opdev/simple-demo-operator-bundle:latest

docker tag quay.io/opdev/simple-demo-operator-catalog:v<NEW_VERSION> quay.io/opdev/simple-demo-operator-catalog:latest
docker push quay.io/opdev/simple-demo-operator-catalog:latest
```

Update github.com/opdev/simple-demo-pipeline to include the new bundle. Remember
to adjust the bundle.Dockerfile so that relative paths are indeed still relative
when moving the bundle over.
