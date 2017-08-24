# API
Base URL:
```
http://{server}:8080/api/v1/
```

## Choice

### Get Choice
```
GET
URL/choice/{path-to-choice-file}
```

>If `{path-to-choice-file}` is empty, a list of available choice file paths will be returned.

Returns JSON contents of choice.

### Set Choice
```
POST
URL/choice/{path-to-choice-file}
BODY: {
    "id": int,
    "name": string
}
```

A new choice file will be created with the JSON contents from the `BODY`.

If choice file already exists, it will be overwritten.

Returns JSON contents of choice.

### Delete Choice
```
DELETE
URL/choice/{path-to-choice-file}
```

Deletes the choice file if it exists.

Returns `{path-to-choice-file}`.