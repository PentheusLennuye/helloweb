# [TEMPLATE]

Description of project

## Dev Environment

The developer using this repo on Microsoft Visual Studio will have a docker
image with the requirements built:

- mkdocs
- mkdocs-mermaid, for diagramming
- python
- go

### Documentation

This repo documentation URL is [URL]

This repository will conform to the documentation standards defined by the
former Eitri team. The standards are written in the
[Eitri Developer Guide for Documentation](https://pages.github.tools.sap/naat/developer-docs/1.1.1/documentation/).

To build the docs for pages no matter the version:

```bash
mkdocs build
mkdocs gh-deploy
git add .
git commit -m'<Message here>'
git push
```

## Additional notes

```sh
docker run -v .:/workspace gcr.io/kaniko-project/executor:latest \
  --dockerfile /workspace/Dockerfile \
  --destination "docker.cummings-online.local/services/helloweb:0.2.0" /
  --context dir:///workspace/
```
