# is-go-first

[![Go Report Card](https://goreportcard.com/badge/github.com/svetlana-rezvaya/is-go-first)](https://goreportcard.com/report/github.com/svetlana-rezvaya/is-go-first)
[![Build Status](https://app.travis-ci.com/svetlana-rezvaya/is-go-first.svg?branch=master)](https://app.travis-ci.com/svetlana-rezvaya/is-go-first)
[![codecov](https://codecov.io/gh/svetlana-rezvaya/is-go-first/branch/master/graph/badge.svg)](https://codecov.io/gh/svetlana-rezvaya/is-go-first)

The playful project for checking that Go is at the top of GitHub.

## Testing

```
$ go test -race -cover -tags integration
```

### GitHub Authentication

It supports GitHub Basic Authentication via a [personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token).

To enable that, set the following environment variables:

- `GITHUB_USERNAME` &mdash; GitHub username;
- `GITHUB_TOKEN` &mdash; GitHub [personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token).

```
$ GITHUB_USERNAME=username GITHUB_TOKEN=token \
  go test -race -cover -tags integration
```

## License

The MIT License (MIT)

Copyright &copy; 2022 svetlana-rezvaya
