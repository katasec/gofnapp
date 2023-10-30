# Overview

This repo is an example of how we might create an [Microsoft Azure Function](https://azure.microsoft.com/en-au/products/functions) using [Go language](https://go.dev/). Go lang doesn't have an existing template whilst using [Azure Functions Core Tools](https://learn.microsoft.com/en-us/azure/azure-functions/functions-run-local?tabs=macos%2Cisolated-process%2Cnode-v4%2Cpython-v2%2Chttp-trigger%2Ccontainer-apps&pivots=programming-language-csharp#install-the-azure-functions-core-tools). As such we need to use a "custom handler" to run Go lang applications either locally or in the cloud.


# Prerequisites

- Install the [Azure Functions Core Tools](https://learn.microsoft.com/en-us/azure/azure-functions/functions-run-local?tabs=macos%2Cisolated-process%2Cnode-v4%2Cpython-v2%2Chttp-trigger%2Ccontainer-apps&pivots=programming-language-csharp#install-the-azure-functions-core-tools).
- Install Go
- Clone this repo

# Running the sample App

- Go to the root of this repo
- Check make options:
    ![Running](./images/options.png)

- Run `make run` to start the app:

    ![Running](./images/running.png)

- Click on `http://localhost:7071/api/func1` to browse the app

# App details

The application is enclosed in the [func1](./func1/) folder in `hello.go`. It uses a standard `http.ServeMux` with a single registered route:

```go
	mux := http.NewServeMux()
	route := fmt.Sprintf("/api/%s", getDirName())
	mux.HandleFunc(route, helloHandler)
```

# Host.json

The host.json file tells the Azure Function runtime how to kick off the app:

```
  "customHandler": {
    "enableForwardingHttpRequest": true,
    "description": {
      "defaultExecutablePath": "func1/func1",
      "workingDirectory": "func1",
      "arguments": [

      ]
    }
  }
```
