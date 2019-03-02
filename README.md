# search-conversion

## Installation

To run the service locally, first either get it via `go get`

```bash
go get -u github.com/afritzler/search-conversion
```

or clone the repository.

```bash
git clone https://github.com/afritzler/search-conversion.git
cd search-conversion
```

## Usage

Build the app via `make` and start the service

```bash
./search-conversion
```

The service will be exposed via http://localhost:8080. The default port 8080 can be changed
by setting the `PORT` environment variable.

```bash
export PORT=5000
./search-conversion
```

Adjust the `examples/request.json` to your needs and run a `curl` against either `localhost`
or your corresponding service URL.

```bash
curl -vX POST https://YOUR_SERVICE_URL/search -d @examples/request.json \
--header "Content-Type: application/json"
```