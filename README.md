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
          "id": "{{ urlParam "id" }}",
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

## Template parameters

- urlParam(paramName string): 
  - Ex.: For URL `/accounts/:id` use `{{ urlParam "id" }}` to print the `:id` parameter. 
- uuid: `{{ uuid }}`
- now: `{{ now }}`
- username: `{{ username }}`
- name: `{{ name }}`
- email: `{{ email }}`
- phone: `{{ phone }}`
- int: `{{ int }}`
- digit: `{{ digit }}`
- digitN(n int): `{{ digitN 4 }}`
- letter: `{{ letter }}`
- letterN(n int): `{{ letterN 10 }}`
- word: `{{ word }}`
- phrase: `{{ phrase }}`
- loremIpsum(n int): `{{ loremIpsum 10 }}`