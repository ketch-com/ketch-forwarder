# Error

An Error MUST be returned with the appropriate HTTP status code with an `Error` JSON object. The `Content-Type` must
be `application/json`.

```http request
HTTP/1.1 404 Not Found
Content-Type: application/json
Content-Length: 238

{
  "apiVersion": "v1",
  "kind": "Error",
  "metadata": {
    "uid": "22880925-aac5-42f9-a653-cb6921d361ff",
    "tenant": "axonic"
  },
  "error": {
    "code": 404,
    "status": "not_found",
    "message": "Not found"
  }
}
```

## Fields

| name              | required? | description                                                |
|-------------------|-----------|------------------------------------------------------------|
| *apiVersion*      | yes       | API version. Must be `dsr/v1`                              |
| *kind*            | yes       | Message kind. Must be `Error`                              |
| *metadata*        | yes       | [Metadata](Metadata.md) object                             |
| *error.code*      | yes       | The HTTP status code                                       |
| *error.status*    | yes       | A string [code](#Error status code) representing the error |
| *error.message*   | yes       | A user-friendly error message (e.g., `"Not found"`)        |

## Error status code

| status        | description                 |
|---------------|-----------------------------|
| conflict      | action cannot be performed  |
| internal      | internal error              |
| unavailable   | service is unavailable      |
| invalid       | validation failed           |
| not_found     | entity does not exist       |
| timeout       | operation timed out         |
| canceled      | operation canceled          |
| forbidden     | operation is not authorized |
| configuration | configuration error         |
| unimplemented | unimplemented error         |
