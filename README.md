# The Made to Measure API

## Prerequisites

It's important to install Go version 1.4

Probably the easiest way to manage go would be using gvm

`bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)`

Install godep

`go get github.com/tools/godep`
`godep restore`

## Running the app

Runnin the app is as simple as

`go run api.go`

## Testing

We're using gingko and gomega

`go get github.com/onsi/ginkgo/ginkgo`
`go get github.com/onsi/gomega`

Do note that Gomega relies on a Ω which can be written in vim using `<C-k> W*` in insert mode

## RAML

To get RAML to render in the browser we need to install raml2html

`npm i -g raml2html`


