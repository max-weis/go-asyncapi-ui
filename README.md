# AsyncAPI UI for Go

This projects demonstrates, how you can serve a AsyncAPI document and the AsyncAPI UI.

## Setup

Run: `go run main.go` to start the server.

## Config

You can change the path to the static file and UI via env vars:

- `ASYNCAPI_URL` defines the static AsyncAPI document url
- `ASYNCAPI_UI_URL` defines the AsyncAPI UI url

## UI

- Navigate to `http://localhost:8000/static/asyncapi` to view the AsyncAPI document
- Navigate to `http://localhost:8000/asyncapi-ui` to view the AsyncAPI UI
