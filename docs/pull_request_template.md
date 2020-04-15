## Heads up

- You must not edit any YAML files in the `docs/cr` and `docs/crd` folder. These are generated.
  - `docs/cr`: Check the unit tests for the CRD to find how the example is defined.
  - `docs/crd`: These are generated based on a YAML string in the according Go package.

## Checklist

- [ ] Consider SIG UX feedback.
- [ ] Update changelog in CHANGELOG.md.

## After mering

To publish CRD documentation changes, please

1. Tag a new release here in [apiextensions](https://github.com/giantswarm/apiextensions/releases)
2. In the `docs` repository, set the apiextensions version [here](https://github.com/giantswarm/docs/blob/11cb2cd5091ea123305086232377a2ffe313e36d/Makefile#L55) to the new version
3. Tag a new release for the `docs` repository
4. Set the version of the `docs-app` in the frontend cluster to that new version
