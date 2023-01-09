# Delete

A Delete request is initiated when a [Data Subject](README.md#Subject) selects a right that allows for deleting of personal data.

To forward a Data Subject Request, Ketch sends a message using the POST method to the configured endpoint. The format of
the message and expected responses depend on the type of right invoked by the [Data Subject](README.md#Subject).

![](https://lucid.app/publicSegments/view/acf2f881-fd25-4e98-98c1-3d803f81ed89/image.png)

## Delete Request

```http request
POST /endpoint HTTP/1.1
Host: www.example.com
Content-Type: application/json
Accept: application/json
Authorization: Bearer $auth

{
  "apiVersion": "dsr/v1",
  "kind": "DeleteRequest",
  "metadata": {
    "uid": "22880925-aac5-42f9-a653-cb6921d361ff",
    "tenant": "axonic"
  },
  "request": {
    "controller": "axonic",
    "property": "axonic.io",
    "environment": "production",
    "regulation": "gdpr",
    "jurisdiction": "eugdpr",
    "identities": [
      {
        "identitySpace": "account_id",
        "identityFormat": "raw",
        "identityValue": "123"
      }
    ],
    "callbacks": [
      {
        "url": "https://dsr.ketch.com/callback",
        "headers": {
          "Authorization": "Bearer $auth"
        }
      }
    ],
    "subject": {
      "email": "test@subject.com",
      "firstName": "Test",
      "lastName": "Subject",
      "addressLine1": "123 Main St",
      "addressLine2": "",
      "city": "Anytown",
      "stateRegionCode": "MA",
      "postalCode": "10123",
      "countryCode": "US",
      "description": "Delete my data"
    },
    "claims": {
      "account_id": "123"
    },
    "submittedTimestamp": 123,
    "dueTimestamp": 123
  }
}
```

### Fields

| name                         | required? | description                                                                                                                                                                            |
|------------------------------|-----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| *apiVersion*                 | yes       | API version. Must be `dsr/v1`                                                                                                                                                          |
| *kind*                       | yes       | Must be `DeleteRequest`                                                                                                                                                                |
| *metadata*                   | yes       | [Metadata](../../runtime/v1/Metadata.md) object                                                                                                                                        |
| *request.controller*         | no        | Code of the Ketch controller tenant. Only supplied if the ultimate controller is different to the `metadata.tenant`                                                                    |
| *request.property*           | yes       | Code of the digital property defined in Ketch                                                                                                                                          |
| *request.environment*        | yes       | Code environment defined in Ketch                                                                                                                                                      |
| *request.regulation*         | yes       | Code of the regulation defined in Ketch                                                                                                                                                |
| *request.jurisdiction*       | yes       | Code of the jurisdiction defined in Ketch                                                                                                                                              |
| *request.identities*         | yes       | Array of [Identities](README.md#Identity)                                                                                                                                              |
| *request.callbacks*          | no        | Array of [Callbacks](README.md#Callback)                                                                                                                                               |
| *request.subject*            | yes       | The [Data Subject](README.md#Subject)                                                                                                                                                  |
| *request.claims*             | no        | Map containing additional non-identity claims that have been added via identity verification or other augmentation methods. Identity claims should be included in `request.identities` |
| *request.submittedTimestamp* | yes       | UNIX timestamp in seconds when the request was submitted.                                                                                                                              |
| *request.dueTimestamp*       | yes       | UNIX timestamp in seconds when the request must be completed by.                                                                                                                       |

## Delete Response / Status Event

A successful response MUST include the `200 OK` response status code and a `DeleteResponse` JSON object.

The `event.results` and `events.documents` are merged with any cached results from previous events. New documents are
added and existing documents are updated.

When the status of Delete Request has changed, a `DeleteStatusEvent` event should be sent to all the callbacks specified
in the request.

Once the status is set to `completed`, `cancelled` or `denied`, then no further events will be accepted.

The `Content-Type` MUST be `application/json`.

### Response

```http request
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 238

{
  "apiVersion": "dsr/v1",
  "kind": "DeleteResponse",
  "metadata": {
    "uid": "22880925-aac5-42f9-a653-cb6921d361ff",
    "tenant": "axonic"
  },
  "response": {
    "status": "in_progress",
    "reason": "other",
    "expectedCompletionTimestamp": 123,
    "requestID": "abc123"
    "identites": [
      {
        "identitySpace": "account_id",
        "identityFormat": "raw",
        "identityValue": "123"
      }
    ],
    "subject": {
      "firstName": "Test",
      "lastName": "Subject",
      "addressLine1": "123 Main St",
      "addressLine2": "Apt 123",
      "stateRegionCode": "MA",
      "postalCode": "10123",
      "countryCode": "US"
    },
    "claims": {
      "contextVariable1": "foo",
      "contextVariable2": 1,
      "contextVariable4": true
    }
  }
}
```

### Event

```http request
POST /callback HTTP/1.1
Host: dsr.ketch.com
Content-Type: application/json
Accept: application/json
Authorization: Bearer $auth

{
  "apiVersion": "dsr/v1",
  "kind": "DeleteStatusEvent",
  "metadata": {
    "uid": "22880925-aac5-42f9-a653-cb6921d361ff",
    "tenant": "axonic"
  },
  "event": {
    "status": "in_progress",
    "reason": "other",
    "expectedCompletionTimestamp": 123,
    "requestID": "abc123"
    "identites": [
      {
        "identitySpace": "account_id",
        "identityFormat": "raw",
        "identityValue": "123"
      }
    ],
    "subject": {
      "firstName": "Test",
      "lastName": "Subject",
      "addressLine1": "123 Main St",
      "addressLine2": "Apt 123",
      "stateRegionCode": "MA",
      "postalCode": "10123",
      "countryCode": "US"
    },
    "claims": {
      "contextVariable1": "foo",
      "contextVariable2": 1,
      "contextVariable4": true
    }
  }
}
```

### Fields

| name                          | required? | description                                                                                       |
|-------------------------------|-----------|---------------------------------------------------------------------------------------------------|
| *status*                      | yes       | The [status](Status.md#Status code) of the Data Subject Request                                   |
| *reason*                      | no        | The [reason](Status.md#Reason) for the status of the Data Subject Request                         |
| *expectedCompletionTimestamp* | no        | The UNIX timestamp at which the Data Subject Request is expected to be completed                  |
| *requestID*                   | no        | The request ID known to the destination system                                                    | 
| *results*                     | no        | Array of [Documents](README.md#Document) that can be used to download the contents requested      |
| *documents*                   | no        | Array of [Documents](README.md#Document) that can be used to download the contents requested      |
| *claims*                      | no        | Map containing additions or changes to claims.                                                    |
| *redirectUrl*                 | no        | if the [Data Subject](README.md#Subject) should be redirected to a URL (perhaps for confirmation) |
| *subject*                     | no        | Map containing additions or changes to subject values [Data Subject](README.md#Subject).          |
| *identities*                  | no        | Array of [Identities](README.md#Identity) to add to the request                                   |
