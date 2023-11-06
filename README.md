# mockit 

`mockit` is a dead simple yet powerful HTTP API mock CLI.

## Features

- Multiple configuration files.
- Template values for response body.

## How to use

1. Define your mocks in the configuration file:
```yml
endpoints:
  - method: "GET"
    url: "/accounts/:id"
    response:
      code: 200
      headers:
        Content-Type: "application/json"
      body: |
        {
          "id": "{{ params "id" }}",
          "name": "{{ name }}",
          "email": "{{ email }}",
          "username": "{{ username }}",
          "fixed": "abcd",
          "createdAt": "{{ now.Format "2006-01-02T15:04:05Z" }}"
        }
  - method: "POST"
    url: "/accounts"
    response:
      code: 201
      headers:
        Content-Type: "application/json"
      body: |
        {
          "id": "{{ uuid }}",
          "name": "{{ name }}",
          "email": "{{ email }}",
          "username": "{{ username }}",
          "fixed": "abcd",
          "createdAt": "{{ now.Format "2006-01-02T15:04:05Z" }}"
        }
```

2. Run CLI:
```bash
mockit --config config.yml --addr :8080
```