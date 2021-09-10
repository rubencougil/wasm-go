# Go ❤️ Web Assembly (simple example)

![go-build](https://github.com/rubencougil/wasm-go/workflows/go-build/badge.svg)

This is a simple example of a [WASM](https://webassembly.org/) file generated using Go. You can change the following
slider to modify a gray scale for this gopher. This image manipulation is not done using JS, is actually performed by
a Go program running into your browser.

Find more in this Medium blog post [here](https://medium.com/@rcougil/web-assembly-go-d01bbfc004cc)

## 🕹 Try it

[Try it here](https://ruben.cougil.com/wasm-go/html/)

## How to run it (for local development)

Execute:

`make start`

To compile Go program into a `.wasm` file and serve it through `nginx` listening in port `:8080`

Then go to http://localhost:8080.

![wasm+go](https://user-images.githubusercontent.com/1073799/84015118-b02c8280-a97b-11ea-9715-970337f888bc.png)
