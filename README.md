# mockit 

`mockit` is a dead simple yet powerful HTTP API mock CLI.

## How to use

1. Define your mocks in the configuration file:
```yml
endpoints:
  - method: "GET"
    url: "/accounts"
    response:
      code: 200
      headers:
        "Content-Type": "application/json"
      body: >
        {
          "id":"1",
          "name":"Gopher"
        }
  - method: "POST"
    url: "/accounts"
    response:
      code: 201
      headers:
        "Content-Type": "application/json"
      body: >
        {
          "id":"1",
          "name":"Gopher"
        }
```

2. Run CLI:
```bash
mockit --config config.yml --addr :8080
```