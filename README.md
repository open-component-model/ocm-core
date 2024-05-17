[![REUSE status](https://api.reuse.software/badge/github.com/open-component-model/ocm-core)](https://api.reuse.software/info/github.com/open-component-model/ocm-core)

# ocm-core

## About this project

This project provides the core Go library containing an API for interacting with the
[Open Component Model (OCM)](https://github.com/open-component-model/ocm-spec) elements and mechanisms.

The library currently supports the following [repository mappings](https://github.com/open-component-model/ocm/tree/main/docs/ocm/interoperability.md):

- **OCI**: Use the repository prefix path of an OCI repository to implement an OCM
  repository.
- **CTF (Common Transport Format)**: Use a file-based binding to represent any set of
  component versions as filesystem content (directory, tar, tgz).
- **Component Archive**: Compose the content of a component version on the
  filesystem.

For the usage of the library to access OCM repositories, handle configuration and credentials see the [examples section](https://github.com/open-component-model/ocm/tree/main/examples/lib/README.md).

Additionally, OCM provides a generic solution for how to:

- Sign component versions in any supported OCM repository implementation.
- Verify signatures based on public keys or verified certificates.
- Transport component versions, per reference or as values to any of the
  repository implementations.

## Support, Feedback, Contributing

This project is open to feature requests/suggestions, bug reports etc. via [GitHub issues](https://github.com/open-component-model/ocm-core/issues). Contribution and feedback are encouraged and always welcome. For more information about how to contribute, the project structure, as well as additional contribution information, see our [Contribution Guidelines](CONTRIBUTING.md).

## Code of Conduct

We as members, contributors, and leaders pledge to make participation in our community a harassment-free experience for everyone. By participating in this project, you agree to abide by its [Code of Conduct](CODE_OF_CONDUCT.md) at all times.

## Licensing

Copyright 2024 SAP SE or an SAP affiliate company and ocm-core contributors. Please see our [LICENSE](LICENSE) for copyright and license information. Detailed information including third-party components and their licensing/copyright information is available [via the REUSE tool](https://api.reuse.software/info/github.com/open-component-model/ocm-core).
