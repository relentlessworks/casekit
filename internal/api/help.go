package api

const helpText = `casekit — agentic-first case conversion service

Convert text between camelCase, PascalCase, snake_case, SCREAMING_SNAKE_CASE,
kebab-case, SCREAMING-KEBAB-CASE, dot.case, Title Case, Sentence case, flatcase,
Train-Case, alternating case, and inverse case. Detect case, convert to all
formats at once. No database needed.

== ENDPOINTS ==

GET  /convert?text=<string>&to=<case>
    Convert text to a specific case format.
    Cases: camel, pascal, snake, screaming_snake, kebab, screaming_kebab,
           dot, title, sentence, flat, train, alternating, inverse
    Example: GET /convert?text=hello_world&to=camel
    Response: camelCase=helloWorld

GET  /detect?text=<string>
    Detect the case type of the input text.
    Example: GET /detect?text=helloWorld
    Response: detected=camel name=camelCase

GET  /all?text=<string>
    Convert text to all supported case formats at once.
    Example: GET /all?text=helloWorld
    Response: one line per case format

GET  /camel?text=<string>
    Convert to camelCase.
    Example: GET /camel?text=hello_world → helloWorld

GET  /pascal?text=<string>
    Convert to PascalCase.
    Example: GET /pascal?text=hello_world → HelloWorld

GET  /snake?text=<string>
    Convert to snake_case.
    Example: GET /snake?text=helloWorld → hello_world

GET  /screaming-snake?text=<string>
    Convert to SCREAMING_SNAKE_CASE.
    Example: GET /screaming-snake?text=helloWorld → HELLO_WORLD

GET  /kebab?text=<string>
    Convert to kebab-case.
    Example: GET /kebab?text=helloWorld → hello-world

GET  /screaming-kebab?text=<string>
    Convert to SCREAMING-KEBAB-CASE.
    Example: GET /screaming-kebab?text=helloWorld → HELLO-WORLD

GET  /dot?text=<string>
    Convert to dot.case.
    Example: GET /dot?text=helloWorld → hello.world

GET  /title?text=<string>
    Convert to Title Case.
    Example: GET /title?text=hello_world → Hello World

GET  /sentence?text=<string>
    Convert to Sentence case.
    Example: GET /sentence?text=hello_world → Hello world

GET  /flat?text=<string>
    Convert to flatcase (all lowercase, no separators).
    Example: GET /flat?text=hello_world → helloworld

GET  /train?text=<string>
    Convert to Train-Case.
    Example: GET /train?text=hello_world → Hello-World

GET  /alternating?text=<string>
    Convert to alternating case (hElLo).
    Example: GET /alternating?text=hello → hElLo

GET  /inverse?text=<string>
    Swap the case of each character (Hello → hELLO).
    Example: GET /inverse?text=HelloWorld → hELLOwORLD

POST /mcp
    MCP (Model Context Protocol) JSON-RPC 2.0 endpoint.
    Methods: initialize, tools/list, tools/call
    Tools: convert, detect, all, camel, pascal, snake, screaming_snake,
           kebab, screaming_kebab, dot, title, sentence, flat, train,
           alternating, inverse

== CASE FORMATS ==

camel           helloWorld        first word lowercase, rest capitalized
pascal          HelloWorld        all words capitalized
snake           hello_world       lowercase, underscore separator
screaming_snake HELLO_WORLD       uppercase, underscore separator
kebab           hello-world       lowercase, hyphen separator
screaming_kebab HELLO-WORLD       uppercase, hyphen separator
dot             hello.world       lowercase, dot separator
title           Hello World       all words capitalized, space separator
sentence        Hello world       first word capitalized, space separator
flat            helloworld        all lowercase, no separator
train           Hello-World       all words capitalized, hyphen separator
alternating     hElLo             alternating lower/upper per character
inverse         hELLOwORLD        swap case of each character

== RESPONSES ==

Plain text by default: one labeled line per result.
    camelCase=helloWorld

For /all, one line per case format:
    camel=helloWorld
    pascal=HelloWorld
    snake=hello_world
    ...

JSON on demand: send Accept: application/json or add ?format=json
    {"camel":"helloWorld","pascal":"HelloWorld","snake":"hello_world",...}

Errors are instructive:
    error: missing text parameter | hint: provide text as ?text=helloWorld

== CONFIG ==

Flags: -addr=:7101 -secret=<token>
Env:   CASEKIT_ADDR=:7101 CASEKIT_SECRET=<token>

No database needed. Pure stateless computation. Single Go binary.
`
