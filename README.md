# casekit

Agentic-first case conversion service. Convert between camelCase, PascalCase, snake_case, SCREAMING_SNAKE_CASE, kebab-case, SCREAMING-KEBAB-CASE, dot.case, Title Case, Sentence case, flatcase, Train-Case, alternating case, and inverse case. Detect case, convert to all formats at once. Plain text API, agent-driven, single Go binary.

## Quick Start

```bash
# Build
make build

# Run (defaults to :7101)
./casekit

# Use it
curl "http://localhost:7101/convert?text=hello_world&to=camel"
# camel=helloWorld

curl "http://localhost:7101/detect?text=helloWorld"
# detected=camel name=camelCase

curl "http://localhost:7101/all?text=helloWorld"
# camel=helloWorld
# pascal=HelloWorld
# snake=hello_world
# screaming_snake=HELLO_WORLD
# kebab=hello-world
# ...

curl "http://localhost:7101/snake?text=helloWorld"
# snake=hello_world
```

## Principles

- **The agent IS the interface** — No UI, no SDK. The API is the product.
- **Plain text by default** — One labeled, grepable line per record. JSON on demand via `Accept: application/json` or `?format=json`.
- **Instructive errors** — Every 4xx includes a hint telling the agent what to do next.
- **Self-documenting** — `GET /help` returns a one-page operating manual.
- **Single static binary** — Go, CGO_ENABLED=0, zero external dependencies.
- **Zero config defaults** — Runs out of the box. Config: defaults < env < flags.
- **MCP connector** — Speaks Model Context Protocol at `POST /mcp`.

## API Reference

### Core Endpoints

| Method | Path | Params | Description |
|--------|------|--------|-------------|
| GET | `/convert` | `text`, `to` | Convert text to a specific case |
| GET | `/detect` | `text` | Detect the case type of input text |
| GET | `/all` | `text` | Convert to all supported case formats |

### Individual Case Endpoints

| Method | Path | Params | Output Format |
|--------|------|--------|---------------|
| GET | `/camel` | `text` | camelCase |
| GET | `/pascal` | `text` | PascalCase |
| GET | `/snake` | `text` | snake_case |
| GET | `/screaming-snake` | `text` | SCREAMING_SNAKE_CASE |
| GET | `/kebab` | `text` | kebab-case |
| GET | `/screaming-kebab` | `text` | SCREAMING-KEBAB-CASE |
| GET | `/dot` | `text` | dot.case |
| GET | `/title` | `text` | Title Case |
| GET | `/sentence` | `text` | Sentence case |
| GET | `/flat` | `text` | flatcase |
| GET | `/train` | `text` | Train-Case |
| GET | `/alternating` | `text` | aLtErNaTiNg |
| GET | `/inverse` | `text` | swap case |

### Supported Case Formats

| Case | Example | Description |
|------|---------|-------------|
| camel | `helloWorld` | First word lowercase, rest capitalized |
| pascal | `HelloWorld` | All words capitalized |
| snake | `hello_world` | Lowercase, underscore separator |
| screaming_snake | `HELLO_WORLD` | Uppercase, underscore separator |
| kebab | `hello-world` | Lowercase, hyphen separator |
| screaming_kebab | `HELLO-WORLD` | Uppercase, hyphen separator |
| dot | `hello.world` | Lowercase, dot separator |
| title | `Hello World` | All words capitalized, space separator |
| sentence | `Hello world` | First word capitalized, space separator |
| flat | `helloworld` | All lowercase, no separator |
| train | `Hello-World` | All words capitalized, hyphen separator |
| alternating | `hElLo` | Alternating lower/upper per character |
| inverse | `hELLOwORLD` | Swap case of each character |

### MCP

`POST /mcp` — JSON-RPC 2.0 endpoint with 16 tools for chat client integrations.

## Configuration

| Source | Key | Default | Description |
|--------|-----|---------|-------------|
| Flag | `-addr` | `:7101` | Listen address |
| Flag | `-secret` | (random) | Auth token signing secret |
| Env | `CASEKIT_ADDR` | `:7101` | Listen address |
| Env | `CASEKIT_SECRET` | (random) | Auth token signing secret |

## Build

```bash
make build    # CGO_ENABLED=0 go build -trimpath ./cmd/casekit
make test     # go test -race ./...
make vet      # go vet ./...
```

No database needed. Pure stateless computation. Single Go binary.

## License

MIT
