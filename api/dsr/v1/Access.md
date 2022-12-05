# Access

An Access Request is initiated when a [Data Subject](README.md#Subject) invokes a right that allows Access/Portability of personal data.

![](https://lucid.app/publicSegments/view/a3a82c6d-1057-435b-966f-125ab982b59f/image.png)

## Access Request

```http request
POST /endpoint HTTP/1.1
Host: www.example.com
Content-Type: application/json
Accept: application/json
Authorization: Bearer $auth

{
  "apiVersion": "dsr/v1",
  "kind": "AccessRequest",
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
      "description": "Access my data"
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

| name                         | required? | description                                                                                                         |
|------------------------------|-----------|---------------------------------------------------------------------------------------------------------------------|
| *apiVersion*                 | yes       | API version. Must be `dsr/v1`                                                                                       |
| *kind*                       | yes       | Message kind. Must be `AccessRequest`                                                                               |
| *metadata*                   | yes       | [Metadata](../../runtime/v1/Metadata.md) object                                                                     |
| *request.controller*         | no        | Code of the Ketch controller tenant. Only supplied if the ultimate controller is different to the `metadata.tenant` |
| *request.property*           | yes       | Code of the digital property defined in Ketch                                                                       |
| *request.environment*        | yes       | Code environment defined in Ketch                                                                                   |
| *request.regulation*         | yes       | Code of the regulation defined in Ketch                                                                             |
| *request.jurisdiction*       | yes       | Code of the jurisdiction defined in Ketch                                                                           |
| *request.identities*         | yes       | Array of [Identities](README.md#Identity)                                                                           |
| *request.callbacks*          | no        | Array of [Callbacks](README.md#Callback)                                                                            |
| *request.subject*            | yes       | The [Data Subject](README.md#Subject)                                                                               |
| *request.claims*             | no        | Map containing additional claims that have been added via identity verification or other augmentation methods       |
| *request.submittedTimestamp* | yes       | UNIX timestamp in seconds                                                                                           |
| *request.dueTimestamp*       | yes       | UNIX timestamp in seconds                                                                                           |

## Access Response / Status Event

A successful response MUST include the `200 OK` response status code and a `AccessResponse` JSON object.

When the status of Access Request has changed, a `AccessStatusEvent` event JSON object should be sent to all the callbacks
specified in the request.

The `event.results` and `events.documents` are merged with any cached results from previous events. New documents are
added and existing documents are updated.

Once the status is set to `completed`, `cancelled` or `denied`, then no further events will be accepted.

The `Content-Type` MUST be `application/json`.

### Response

```http request
HTTP/1.1 200 OK
Content-Type: application/json

{
  "apiVersion": "dsr/v1",
  "kind": "AccessResponse",
  "metadata": {
    "uid": "22880925-aac5-42f9-a653-cb6921d361ff",
    "tenant": "axonic"
  },
  "response": {
    "status": "in_progress",
    "expectedCompletionTimestamp": 123,
    "results": [
      {
        "url": "https://example.com/results",
        "headers": {
          "Authorization": "Bearer $auth"
        }
      },
      {
        "data": "eyJrZXkiOiAidmFsdWUifQ==",
        "headers": {
          "Content-Type": "application/json"
        }
      }
    ],
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
      "contextVariable3": 1.0,
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
  "kind": "AccessStatusEvent",
  "metadata": {
    "uid": "22880925-aac5-42f9-a653-cb6921d361ff",
    "tenant": "axonic"
  },
  "event": {
    "status": "in_progress",
    "expectedCompletionTimestamp": 123,
    "results": [
      {
        "url": "https://example.com/results",
        "headers": {
          "Authorization": "Bearer $auth"
        }
      },
      {
        "data": "eyJrZXkiOiAidmFsdWUifQ==",
        "headers": {
          "Content-Type": "application/json"
        }
      }
    ],
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
| *reason*                      | no        | The [reason](Status.md#Reason) for the status of the Data Subject Request. Default is `other`.    |
| *expectedCompletionTimestamp* | no        | The UNIX timestamp at which the Data Subject Request is expected to be completed                  |
| *requestID*                   | no        | The request ID known to the destination system                                                    |
| *results*                     | no        | Array of [Documents](README.md#Document) that can be used to download the contents requested      |
| *documents*                   | no        | Array of [Documents](README.md#Document) that can be used to download the contents requested      |
| *claims*                      | no        | Map containing additions or changes to claims.                                                    |
| *redirectUrl*                 | no        | if the [Data Subject](README.md#Subject) should be redirected to a URL (perhaps for confirmation) |
| *subject*                     | no        | Map containing additions or changes to subject values [Data Subject](README.md#Subject).          |
| *identities*                  | no        | Array of [Identities](README.md#Identity) to add to the request                                   |

