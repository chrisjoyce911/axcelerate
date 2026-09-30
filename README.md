# axcelerate [![Go Report Card](https://goreportcard.com/badge/github.com/chrisjoyce911/axcelerate)](https://goreportcard.com/report/github.com/chrisjoyce911/axcelerate) [![Go Reference](https://pkg.go.dev/badge/github.com/chrisjoyce911/axcelerate.svg)](https://pkg.go.dev/github.com/chrisjoyce911/axcelerate) [![Build Status](https://travis-ci.org/multiplay/go-battleye.svg?branch=master)](https://travis-ci.org/chrisjoyce911/axcelerate) [![License](https://img.shields.io/badge/license-unlicense-blue.svg)](https://github.com/chrisjoyce911/axcelerate/blob/master/LICENSE)

Provides a simple interface to theRESTFul API for interfacing with aXcelerate.


## Roadmap

This library is being initially developed for use with the [aXcelerate RESTFul Service API](https://admin.axcelerate.com.au/apidocs/), so API methods will likely be implemented in the order that they are needed by any project that accesses this service.

## Undocumented endpoints

Some endpoints the library wraps are live but missing from aXcelerate's published docs:

- `GET /domains` — `Accounting.Domains()`: every domain id and name (135 on
  staging and production, 29 Sep 2026). aXcelerate stores any `DOMAINID` on an
  invoice line without checking it exists, so this list is the only way to
  validate one before writing it.

## Contributing

I would like to cover the entire aXcelerate RESTFul Service API and contributions are of course always welcome. See CONTRIBUTING.md for details.
```

## Secrets (SOPS)

`.env` is gitignored and never committed. The encrypted copy `.env.enc` **is**
committed and is the portable source of truth — clone the repo on any machine and
decrypt, rather than copying `.env` around by hand.

```sh
# after cloning or pulling
sops -d --input-type dotenv --output-type dotenv .env.enc > .env

# after changing .env — re-encrypt and commit the .enc file
sops -e --input-type dotenv --output-type dotenv .env > .env.enc
```

Requires the personal age key at `~/.config/sops/age/keys.txt` (0600). On macOS
sops looks in `~/Library/Application Support/sops/age/keys.txt`, so that path must
symlink to it. Recipient: `age1yngetl7fdm7m0dlycfn3mfrgcvnj6ezeyjan2dppla78ndf7lsts3l5wve`.

Note: sops's dotenv parser drops blank lines, so a decrypted `.env` may differ
cosmetically from the original. Keys and values are preserved exactly.
