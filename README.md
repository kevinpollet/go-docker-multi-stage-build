# go-docker-multi-stage-build

[![Build Status](https://github.com/kevinpollet/go-docker-multi-stage-build/workflows/build/badge.svg)](https://github.com/kevinpollet/go-docker-multi-stage-build/actions)
[![License](https://img.shields.io/github/license/kevinpollet/go-docker-multi-stage-build)](./LICENSE.md)

This repository shows how to build a Docker image for a [Go](https://go.dev/) app with the [Docker multi-stage feature](https://docs.docker.com/develop/develop-images/multistage-build/) available since version `17.05`.

## Build image

```shell
$ docker build . -t greet
```

## Run container

```shell
$ docker run --rm greet
Hello, World!
```

```shell
$ docker run --rm greet John
Hello, John!
```

## Contributing

Contributions are welcome!

Want to file a bug, request a feature or contribute some code?

- Check out the [Code of Conduct](./CODE_OF_CONDUCT.md)
- Check if there is an existing issue corresponding to your bug or feature
- Open an issue to discuss your bug or new feature

## License

[MIT](./LICENSE.md)
