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

## Invoice writes — rules the API docs get wrong

`CreateInvoice`, `AddInvoiceLines`, `ApproveInvoice`, `ApplyPayment` and
`Transactions` enforce these before a request is sent. The order aXcelerate
requires is create → add lines → approve → pay:

```go
a := client.Accounting
inv, _, err := a.CreateInvoice(axcelerate.InvoiceRequest{
	ContactID: 14518907, FirstName: "Chris", LastName: "Joyce",
	Lines: []axcelerate.InvoiceLine{{Description: "Pocket Mask Training", Qty: 1,
		UnitPriceGross: "19.00", TaxPercent: "0.00", DomainID: axcelerate.Int(5566)}},
})
_, err = a.AddInvoiceLines(inv.InvoiceID, axcelerate.InvoiceLine{Description: "Afterpay transaction fee",
	Qty: 1, UnitPriceGross: "2.00", TaxPercent: "0.00", DomainID: axcelerate.Int(5494)})
_, err = a.ApproveInvoice(inv.InvGUID) // a draft cannot take a payment
txn, _, err := a.ApplyPayment(axcelerate.InvoicePayment{
	ContactID: 14518907, InvoiceID: inv.InvoiceID, Amount: "21.00"}) // dollars
```

The full, runnable version is `invoiceFlow` in [`example/`](example/README.md).
Each rule below was established against the live API (most of them in
production):

- **The Bill To surname is `lastname`.** The docs say `surname`; aXcelerate
  ignores it and stores `LASTNAME` null. `firstname` and `lastname` are both
  required ("The Bill To Given Name and Surname fields cannot be empty").
- **`invoiceDate` and `orderDate` are required** ("The invoice order date is
  out of bounds"); the client defaults both to today.
- **`ITEMCODE` must be on every line**, even as `""` — a line without the key
  fails the whole request with "The invoice items are invalid."
- **The invoice PUT's `aItem` replaces every line.** `AddInvoiceLines` reads
  the invoice and re-sends the existing lines with their `ITEMID`; sending
  only the new line deletes the rest. The JSON goes in the form body — a PUT's
  query string does not survive `&`, `#` or `+` in a description.
- **A draft invoice cannot take a payment — and a payment against one still
  records the money**, as an unassigned receipt, while answering HTTP 400.
  Retrying after that 400 pays twice. `ApplyPayment` refuses a draft before
  sending anything; `ApproveInvoice` (by GUID) issues it first.
- **Amounts are dollars**, as decimal strings — not a gateway's cents.
- **`DOMAINID` is not validated** by aXcelerate; check ids against `Domains()`.
- **A draft cannot be voided** — only an issued invoice.
- **`Transaction` decoding:** `ORGANISATION` is a name, `ORGID` a quoted id and
  `UNASSIGNEDAMOUNT` a decimal string. Before these were typed correctly a
  payment that succeeded returned a decode error.

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
