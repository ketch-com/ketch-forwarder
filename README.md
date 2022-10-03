# Ketch Forwarder

## Introduction

This specification defines a web protocol encoding a set of standardized request/response data flows such that Data Subjects
can exercise Personal Data Rights and express Consent choices.

### Terminology

The keywords “MUST”, “MUST NOT”, “REQUIRED”, “SHALL”, “SHALL NOT”, “SHOULD”, “SHOULD NOT”, “RECOMMENDED”, “NOT
RECOMMENDED”, “MAY”, and “OPTIONAL” in this document are to be interpreted as described in BCP 14 [RFC2119] [RFC8174]
when, and only when, they appear in all capitals, as shown here.

* Data Subject is the individual who is exercising their rights or expressing their consents. This Data Subject may or
  may not have a direct business relationship or login credentials with the Covered Business.

## Endpoint

To establish forwarding, an endpoint MUST be configured in the Ketch user interface. The required information includes
the endpoint URL, authorization header key and value to send. The endpoint MUST use the `https` protocol.

For the examples below, the endpoint is configured as follows (note this URL does not really exist):

* URL: `https://example.com/endpoint`
* Header Key: `Authorization`
* Header Value: `Bearer $auth`

For the examples below, a callback could be configured as follows (note this URL does not really exist):

* URL: `https://dsr.ketch.com/callback`
* Header Key: `Authorization`
* Header Value: `Bearer $auth`

For the examples below, a result could be configured as follows (note this URL does not really exist):

* URL: `https://example.com/results`
* Header Key: `Authorization`
* Header Value: `Bearer $auth`

## Delete

A Delete request is initiated when a [Data Subject](#Subject) selects a right that allows for deleting of personal data.

To forward a Data Subject Request, Ketch sends a message using the POST method to the configured endpoint. The format of 
the message and expected responses depend on the type of right invoked by the [Data Subject](#Subject).

![](https://lucid.app/publicSegments/view/acf2f881-fd25-4e98-98c1-3d803f81ed89/image.png)

### Delete Request

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

#### Fields

| name                         | required? | description                                                                                                                                                                            |
|------------------------------|-----------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| *apiVersion*                 | yes       | API version. Must be `dsr/v1`                                                                                                                                                          |
| *kind*                       | yes       | Must be `DeleteRequest`                                                                                                                                                                |
| *metadata*                   | yes       | [Metadata](#Metadata) object                                                                                                                                                           |
| *request.controller*         | no        | Code of the Ketch controller tenant. Only supplied if the ultimate controller is different to the `metadata.tenant`                                                                    |
| *request.property*           | yes       | Code of the digital property defined in Ketch                                                                                                                                          |
| *request.environment*        | yes       | Code environment defined in Ketch                                                                                                                                                      |
| *request.regulation*         | yes       | Code of the regulation defined in Ketch                                                                                                                                                |
| *request.jurisdiction*       | yes       | Code of the jurisdiction defined in Ketch                                                                                                                                              |
| *request.identities*         | yes       | Array of [Identities](#Identity)                                                                                                                                                       |
| *request.callbacks*          | no        | Array of [Callbacks](#Callback)                                                                                                                                                        |
| *request.subject*            | yes       | The [Data Subject](#Subject)                                                                                                                                                           |
| *request.claims*             | no        | Map containing additional non-identity claims that have been added via identity verification or other augmentation methods. Identity claims should be included in `request.identities` |
| *request.submittedTimestamp* | yes       | UNIX timestamp in seconds when the request was submitted.                                                                                                                              |
| *request.dueTimestamp*       | yes       | UNIX timestamp in seconds when the request must be completed by.                                                                                                                       |

### Successful Delete Response

A successful response MUST include the `200 OK` response status code and a `DeleteResponse` JSON object. The `Content-Type`
MUST be `application/json`.

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
    "requestID": "abc123",
    "claims": {}
  }
}
```

#### Fields

| name                                   | required? | description                                                                      |
|----------------------------------------|-----------|----------------------------------------------------------------------------------|
| *apiVersion*                           | yes       | API version. Must be `dsr/v1`                                                    |
| *kind*                                 | yes       | Message kind. Must be `DeleteResponse`                                           |
| *metadata*                             | yes       | [Metadata](#Metadata) object                                                     |
| *response.status*                      | yes       | The [status](#Status) of the Data Subject Request                                |
| *response.reason*                      | no        | The [reason](#Reason) for the status of the Data Subject Request                 |
| *response.expectedCompletionTimestamp* | no        | The UNIX timestamp at which the Data Subject Request is expected to be completed |
| *response.requestID*                   | no        | The request ID known to the destination system                                   | 
| *response.claims*                      | no        | The response claims to provide as augmentation of the original request           |

<!-- | *response.redirectUrl* | no | If the [Data Subject](#Subject) should be redirected to a URL (perhaps for confirmation), this field specifies the URL to redirect the [Data Subject](#Subject) to | --> 

### Delete Error Response

If an error occurs, a standard [Error](#Error) object is returned.

### Delete Status Event

When the status of Delete Request has changed, an event should be sent to all the callbacks specified. The following is
an example of a `DeleteStatusEvent`.

```http request
POST /callback HTTP/1.1
Host: dsr.ketch.com
Content-Type: application/json
Accept: application/json

{
  "apiVersion": "dsr/v1",
  "kind": "DeleteStatusEvent",
  "metadata": {
    "uid": "22880925-aac5-42f9-a653-cb6921d361ff",
    "tenant": "axonic"
  },
  "event": {
    "status": "completed",
    "expectedCompletionTimestamp": 123
  }
}
```

#### Fields

| name                                   | required? | description                                                                       |
|----------------------------------------|-----------|-----------------------------------------------------------------------------------|
| *apiVersion*                           | yes       | API version. Must be `dsr/v1`                                                     |
| *kind*                                 | yes       | Message kind. Must be `DeleteStatusEvent`                                         |
| *metadata*                             | yes       | [Metadata](#Metadata) object                                                      |
| *response.status*                      | yes       | The [status](#Status) of the Data Subject Request                                 |
| *response.reason*                      | no        | The [reason](#Reason) for the status of the Data Subject Request                  |
| *response.expectedCompletionTimestamp* | no        | The UNIX time stamp at which the Data Subject Request is expected to be completed |
| *response.requestID*                   | no        | The request ID known to the destination system                                    |
| *response.claims*                      | no        | The response claims to provide as augmentation of the original request            |

<!-- | *response.redirectUrl* | no | If the [Data Subject](#Subject) should be redirected to a URL (perhaps for confirmation), this field specifies the URL to redirect the [Data Subject](#Subject) to | --> 

## Access

An Access Request is initiated when a [Data Subject](#Subject) invokes a right that allows Access/Portability of personal data.

![](https://lucid.app/publicSegments/view/a3a82c6d-1057-435b-966f-125ab982b59f/image.png)

### Access Request

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

#### Fields

| name                         | required? | description                                                                                                         |
|------------------------------|-----------|---------------------------------------------------------------------------------------------------------------------|
| *apiVersion*                 | yes       | API version. Must be `dsr/v1`                                                                                       |
| *kind*                       | yes       | Message kind. Must be `AccessRequest`                                                                               |
| *metadata*                   | yes       | [Metadata](#Metadata) object                                                                                        |
| *request.controller*         | no        | Code of the Ketch controller tenant. Only supplied if the ultimate controller is different to the `metadata.tenant` |
| *request.property*           | yes       | Code of the digital property defined in Ketch                                                                       |
| *request.environment*        | yes       | Code environment defined in Ketch                                                                                   |
| *request.regulation*         | yes       | Code of the regulation defined in Ketch                                                                             |
| *request.jurisdiction*       | yes       | Code of the jurisdiction defined in Ketch                                                                           |
| *request.identities*         | yes       | Array of [Identities](#Identity)                                                                                    |
| *request.callbacks*          | no        | Array of [Callbacks](#Callback)                                                                                     |
| *request.subject*            | yes       | The [Data Subject](#Subject)                                                                                        |
| *request.claims*             | no        | Map containing additional claims that have been added via identity verification or other augmentation methods       |
| *request.submittedTimestamp* | yes       | UNIX timestamp in seconds                                                                                           |
| *request.dueTimestamp*       | yes       | UNIX timestamp in seconds                                                                                           |

### Successful Access Response

A successful response MUST include the `200 OK` response status code and a `AccessResponse` JSON object. The `Content-Type`
MUST be `application/json`.

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
    "results": []
  }
}
```

#### Fields

| name                                   | required? | description                                                                           |
|----------------------------------------|-----------|---------------------------------------------------------------------------------------|
| *apiVersion*                           | yes       | API version. Must be `dsr/v1`                                                         |
| *kind*                                 | yes       | Message kind. Must be `AccessResponse`                                                |
| *metadata*                             | yes       | [Metadata](#Metadata) object                                                          |
| *response.status*                      | yes       | The [status](#Status) of the Data Subject Request                                     |
| *response.reason*                      | no        | The [reason](#Reason) for the status of the Data Subject Request. Default is `other`. |
| *response.expectedCompletionTimestamp* | no        | The UNIX time stamp at which the Data Subject Request is expected to be completed     |
| *response.results*                     | no        | Array of [Callbacks](#Callback) that can be used to download the contents requested   |
| *response.requestID*                   | no        | The request ID known to the destination system                                        |
| *response.claims*                      | no        | The response claims to provide as augmentation of the original request                |

<!-- | *response.redirectUrl* | no        | if the [Data Subject](#Subject) should be redirected to a URL (perhaps for confirmation)         | -->

### Access Error Response

If an error occurs, a standard [Error](#Error) object is returned.

### Access Status Event

When the status of Access Request has changed, an event should be sent to all the callbacks specified. The `event.results`
are merged with any cached results from previous events. New URLs are added and existing URLs are updated.
The following is an example of a `AccessStatusEvent`. Once the status is set to `completed`, `cancelled` or `denied`,
then no further events will be accepted.

```http request
POST /callback HTTP/1.1
Host: dsr.ketch.com
Content-Type: application/json
Accept: application/json

{
  "apiVersion": "dsr/v1",
  "kind": "AccessStatusEvent",
  "metadata": {
    "uid": "22880925-aac5-42f9-a653-cb6921d361ff",
    "tenant": "axonic"
  },
  "event": {
    "status": "completed",
    "expectedCompletionTimestamp": 123,
    "results": [
      {
        "url": "https://example.com/results",
        "headers": {
          "Authorization": "Bearer $auth"
        }
      }
    ]
  }
}
```

#### Fields

| name                                   | required? | description                                                                         |
|----------------------------------------|-----------|-------------------------------------------------------------------------------------|
| *apiVersion*                           | yes       | API version. Must be `dsr/v1`                                                       |
| *kind*                                 | yes       | Message kind. Must be `AccessStatusEvent`                                           |
| *metadata*                             | yes       | [Metadata](#Metadata) object                                                        |
| *response.status*                      | yes       | The [status](#Status) of the Data Subject Request                                   |
| *response.reason*                      | no        | The [reason](#Reason) for the status of the Data Subject Request                    |
| *response.expectedCompletionTimestamp* | no        | The UNIX time stamp at which the Data Subject Request is expected to be completed   |
| *response.results*                     | no        | Array of [Callbacks](#Callback) that can be used to download the contents requested |
| *response.requestID*                   | no        | The request ID known to the destination system                                      |
| *response.claims*                      | no        | The response claims to provide as augmentation of the original request              |

## Restrict Processing

A Restrict Processing Request is initiated when a [Data Subject](#Subject) invokes a right that allows restriction of processing
of personal data.

### Restrict Processing Request

```http request
POST /endpoint HTTP/1.1
Host: www.example.com
Content-Type: application/json
Accept: application/json
Authorization: Bearer $auth

{
  "apiVersion": "dsr/v1",
  "kind": "RestrictProcessingRequest",
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
    "purposes": [
      "advertising",
      "retargeting",
      "analytics"
    ],
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
      "description": "Restrict my data"
    },
    "claims": {
      "account_id": "123"
    },
    "submittedTimestamp": 123,
    "dueTimestamp": 123
  }
}
```

#### Fields

| name                         | required? | description                                                                                                         |
|------------------------------|-----------|---------------------------------------------------------------------------------------------------------------------|
| *apiVersion*                 | yes       | API version. Must be `dsr/v1`                                                                                       |
| *kind*                       | yes       | Message kind. Must be `RestrictProcessingRequest`                                                                   |
| *metadata*                   | yes       | [Metadata](#Metadata) object                                                                                        |
| *request.controller*         | no        | Code of the Ketch controller tenant. Only supplied if the ultimate controller is different to the `metadata.tenant` |
| *request.property*           | yes       | Code of the digital property defined in Ketch                                                                       |
| *request.environment*        | yes       | Code environment defined in Ketch                                                                                   |
| *request.regulation*         | yes       | Code of the regulation defined in Ketch                                                                             |
| *request.jurisdiction*       | yes       | Code of the jurisdiction defined in Ketch                                                                           |
| *request.purposes*           | yes       | List of purpose codes defined in Ketch                                                                              |
| *request.identities*         | yes       | Array of [Identities](#Identity)                                                                                    |
| *request.callbacks*          | no        | Array of [Callbacks](#Callback)                                                                                     |
| *request.subject*            | yes       | The [Data Subject](#Subject)                                                                                        |
| *request.claims*             | no        | Map containing additional claims that have been added via identity verification or other augmentation methods       |
| *request.submittedTimestamp* | yes       | UNIX timestamp in seconds                                                                                           |
| *request.dueTimestamp*       | yes       | UNIX timestamp in seconds                                                                                           |

### Successful Restrict Processing Response

A successful response MUST include the `200 OK` response status code and a `RestrictProcessingResponse` JSON object. The `Content-Type`
MUST be `application/json`.

```http request
HTTP/1.1 200 OK
Content-Type: application/json

{
  "apiVersion": "dsr/v1",
  "kind": "RestrictProcessingResponse",
  "metadata": {
    "uid": "22880925-aac5-42f9-a653-cb6921d361ff",
    "tenant": "axonic"
  },
  "response": {
    "status": "pending",
    "expectedCompletionTimestamp": 123
  }
}
```

#### Fields

| name                                   | required? | description                                                                        |
|----------------------------------------|-----------|------------------------------------------------------------------------------------|
| *apiVersion*                           | yes       | API version. Must be `dsr/v1`                                                      |
| *kind*                                 | yes       | Message kind. Must be `RestrictProcessingResponse`                                 |
| *metadata*                             | yes       | [Metadata](#Metadata) object                                                       |
| *response.status*                      | yes       | The [status](#Status) of the Data Subject Request                                  |
| *response.reason*                      | no        | The [reason](#Reason) for the status of the Data Subject Request                   |
| *response.expectedCompletionTimestamp* | no        | The UNIX time stamp at which the Data Subject Request is expected to be completed  |
| *response.requestID*                   | no        | The request ID known to the destination system                                     |
| *response.claims*                      | no        | The response claims to provide as augmentation of the original request             |

<!-- | *response.redirectUrl*                 | no                                                                                 | If the [Data Subject](#Subject) should be redirected to a URL (perhaps for confirmation)      | -->

### Restrict Processing Error Response

If an error occurs, a standard [Error](#Error) object is returned.

### Restrict Processing Status Event

When the status of Restrict Processing Request has changed, an event should be sent to all the callbacks specified. The `event.results`
are merged with any cached results from previous events. New URLs are added and existing URLs are updated.
The following is an example of a `RestrictProcessingStatusEvent`. Once the status is set to `completed`, `cancelled` or `denied`,
then no further events will be accepted.

```http request
POST /callback HTTP/1.1
Host: dsr.ketch.com
Content-Type: application/json
Accept: application/json

{
  "apiVersion": "dsr/v1",
  "kind": "RestrictProcessingStatusEvent",
  "metadata": {
    "uid": "22880925-aac5-42f9-a653-cb6921d361ff",
    "tenant": "axonic"
  },
  "event": {
    "status": "completed",
    "expectedCompletionTimestamp": 123
  }
}
```

#### Fields

| name                                   | required? | description                                                                       |
|----------------------------------------|-----------|-----------------------------------------------------------------------------------|
| *apiVersion*                           | yes       | API version. Must be `dsr/v1`                                                     |
| *kind*                                 | yes       | Message kind. Must be `RestrictProcessingStatusEvent`                             |
| *metadata*                             | yes       | [Metadata](#Metadata) object                                                      |
| *response.status*                      | yes       | The [status](#Status) of the Data Subject Request                                 |
| *response.reason*                      | no        | The [reason](#Reason) for the status of the Data Subject Request                  |
| *response.expectedCompletionTimestamp* | no        | The UNIX time stamp at which the Data Subject Request is expected to be completed |
| *response.requestID*                   | no        | The request ID known to the destination system                                    |
| *response.claims*                      | no        | The response claims to provide as augmentation of the original request            |

## Correction

A Correction Request is initiated when a [Data Subject](#Subject) invokes a right that allows correction of personal data. Note that
only the basic details are transported since collecting specific correction details are currently beyond the scope.

### Correction Request

```http request
POST /endpoint HTTP/1.1
Host: www.example.com
Content-Type: application/json
Accept: application/json
Authorization: Bearer $auth

{
  "apiVersion": "dsr/v1",
  "kind": "CorrectionRequest",
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
      "description": "Correct my name to Test Object"
    },
    "claims": {
      "account_id": "123"
    },
    "submittedTimestamp": 123,
    "dueTimestamp": 123
  }
}
```

#### Fields

| name                         | required? | description                                                                                                         |
|------------------------------|-----------|---------------------------------------------------------------------------------------------------------------------|
| *apiVersion*                 | yes       | API version. Must be `dsr/v1`                                                                                       |
| *kind*                       | yes       | Message kind. Must be `CorrectionRequest`                                                                           |
| *metadata*                   | yes       | [Metadata](#Metadata) object                                                                                        |
| *request.controller*         | no        | Code of the Ketch controller tenant. Only supplied if the ultimate controller is different to the `metadata.tenant` |
| *request.property*           | yes       | Code of the digital property defined in Ketch                                                                       |
| *request.environment*        | yes       | Code environment defined in Ketch                                                                                   |
| *request.regulation*         | yes       | Code of the regulation defined in Ketch                                                                             |
| *request.jurisdiction*       | yes       | Code of the jurisdiction defined in Ketch                                                                           |
| *request.identities*         | yes       | Array of [Identities](#Identity)                                                                                    |
| *request.callbacks*          | no        | Array of [Callbacks](#Callback)                                                                                     |
| *request.subject*            | yes       | The [Data Subject](#Subject)                                                                                        |
| *request.claims*             | no        | Map containing additional claims that have been added via identity verification or other augmentation methods       |
| *request.submittedTimestamp* | yes       | UNIX timestamp in seconds                                                                                           |
| *request.dueTimestamp*       | yes       | UNIX timestamp in seconds                                                                                           |

### Successful Correction Response

A successful response MUST include the `200 OK` response status code and a `CorrectionResponse` JSON object. The `Content-Type`
MUST be `application/json`.

```http request
HTTP/1.1 200 OK
Content-Type: application/json

{
  "apiVersion": "dsr/v1",
  "kind": "CorrectionResponse",
  "metadata": {
    "uid": "22880925-aac5-42f9-a653-cb6921d361ff",
    "tenant": "axonic"
  },
  "response": {
    "status": "in_progress",
    "expectedCompletionTimestamp": 123
  }
}
```

#### Fields

| name                                   | required? | description                                                                       |
|----------------------------------------|-----------|-----------------------------------------------------------------------------------|
| *apiVersion*                           | yes       | API version. Must be `dsr/v1`                                                     |
| *kind*                                 | yes       | Message kind. Must be `CorrectionResponse`                                        |
| *metadata*                             | yes       | [Metadata](#Metadata) object                                                      |
| *response.status*                      | yes       | The [status](#Status) of the Data Subject Request                                 |
| *response.reason*                      | no        | The [reason](#Reason) for the status of the Data Subject Request                  |
| *response.expectedCompletionTimestamp* | no        | The UNIX time stamp at which the Data Subject Request is expected to be completed |
| *response.requestID*                   | no        | The request ID known to the destination system                                    |
| *response.claims*                      | no        | The response claims to provide as augmentation of the original request            |

<!-- | *response.redirectUrl* | no | If the [Data Subject](#Subject) should be redirected to a URL (perhaps for confirmation)     | -->

### Correction Error Response

If an error occurs, a standard [Error](#Error) object is returned.

### Correction Status Event

When the status of Correction Request has changed, an event should be sent to all the callbacks specified. The following
is an example of a `CorrectionStatusEvent`. Once the status is set to `completed`, `cancelled` or `denied`, then no
further events will be accepted.

```http request
POST /callback HTTP/1.1
Host: dsr.ketch.com
Content-Type: application/json
Accept: application/json

{
  "apiVersion": "dsr/v1",
  "kind": "CorrectionStatusEvent",
  "metadata": {
    "uid": "22880925-aac5-42f9-a653-cb6921d361ff",
    "tenant": "axonic"
  },
  "event": {
    "status": "completed",
    "expectedCompletionTimestamp": 123
  }
}
```

#### Fields

| name                                   | required? | description                                                                       |
|----------------------------------------|-----------|-----------------------------------------------------------------------------------|
| *apiVersion*                           | yes       | API version. Must be `dsr/v1`                                                     |
| *kind*                                 | yes       | Message kind. Must be `CorrectionStatusEvent`                                     |
| *metadata*                             | yes       | [Metadata](#Metadata) object                                                      |
| *response.status*                      | yes       | The [status](#Status) of the Data Subject Request                                 |
| *response.reason*                      | no        | The [reason](#Reason) for the status of the Data Subject Request                  |
| *response.expectedCompletionTimestamp* | no        | The UNIX time stamp at which the Data Subject Request is expected to be completed |
| *response.requestID*                   | no        | The request ID known to the destination system                                    |
| *response.claims*                      | no        | The response claims to provide as augmentation of the original request            |

## Common objects

### Metadata

| name     | required? | description                                                  |
|----------|-----------|--------------------------------------------------------------|
| *uid*    | yes       | Will be a unique UUIDv4, and uniquely identifies the request |
| *tenant* | yes       | Will be the Ketch tenant code where the request originated   |

### Error

An Error MUST be returned with the appropriate HTTP status code with an `Error` JSON object. The `Content-Type` must
be `application/json`.

```http request
HTTP/1.1 404 Not Found
Content-Type: application/json
Content-Length: 238

{
  "apiVersion": "dsr/v1",
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

#### Fields

| name              | required? | description                                                  |
|-------------------|-----------|--------------------------------------------------------------|
| *apiVersion*      | yes       | API version. Must be `dsr/v1`                                |
| *kind*            | yes       | Message kind. Must be `Error`                                |
| *metadata*        | yes       | [Metadata](#Metadata) object                                 |
| *error.code*      | yes       | The HTTP status code                                         |
| *error.status*    | yes       | A string code representing the error                         |
| *error.message*   | yes       | A user-friendly error message (e.g., `"Not found"`)          |

### Claims

The Claims object provides a free-form object with additional information in the request or response that does not
fit in other standard fields.  This can be a complex object.

### Callback

The Callback object defines a URL and associated headers to be used for communicating information.

```json
{
  "url": "https://dsr.ketch.com/callback",
  "headers": {
    "Authorization": "Bearer $auth"
  }
}
```

#### Fields

| name      | required? | description                                     |
|-----------|-----------|-------------------------------------------------|
| *url*     | yes       | URL of the callback endpoint                    |
| *headers* | no        | map of headers to send to the callback endpoint |

### Identity

The Identity object describes the identity of a [Data Subject](#Subject).

```json
{
  "identitySpace": "account_id",
  "identityFormat": "raw",
  "identityValue": "123"
}
```

#### Fields

| name             | required? | description                                                            |
|------------------|-----------|------------------------------------------------------------------------|
| *identitySpace*  | yes       | Identity space code setup in Ketch                                     |
| *identityFormat* | no        | Format of the identity value (`raw`, `md5`, `sha1`). Default is `raw`. |
| *identityValue*  | yes       | Value of the identity                                                  |

### Subject

The Subject object describes the Data Subject making the request.

```json
{
  "email": "test@subject.com",
  "firstName": "Test",
  "lastName": "Subject",
  "addressLine1": "123 Main St",
  "addressLine2": "",
  "city": "Anytown",
  "stateRegionCode": "MA",
  "postalCode": "10123",
  "countryCode": "US",
  "description": "Restrict my data"
}
```

#### Fields

| name              | required? | description                                                            |
|-------------------|-----------|------------------------------------------------------------------------|
| *email*           | yes       | Email address provided by the Data Subject                             |
| *firstName*       | yes       | First name provided by the Data Subject                                |
| *lastName*        | yes       | Last name provided by the Data Subject                                 |
| *addressLine1*    | no        | Address line 1 provided by the Data Subject                            |
| *addressLine2*    | no        | Address line 2 provided by the Data Subject                            |
| *city*            | no        | City provided by the Data Subject                                      |
| *stateRegionCode* | no        | State/region code (e.g., CA) provided by the Data Subject              |
| *postalCode*      | no        | Zip/postal code provided by the Data Subject                           |
| *countryCode*     | no        | Two-character ISO country code (e.g., US) provided by the Data Subject |
| *description*     | no        | Free-text description provided by the Data Subject                     |

### Status

The `status` of a Request progresses through the following state/activity diagram:

![](https://lucid.app/publicSegments/view/61b83862-bbc8-41b1-bdae-d92b7bf87af6/image.png)

Note that multiple Status events can be received with each status. For example, a callback may be called daily with
a status of `in_progress` until finally the final status is sent with status of `completed`.

If a `Response` does not include return a status of `completed`, `cancelled` or `denied`, then a subsequent `StatusEvent`
event MUST be sent to the callbacks with one of those statuses. Otherwise, the request will be marked as failed after
some time.

#### Status code

| name          | description                     |
|---------------|---------------------------------|
| *unknown*     | the status is unknown           |
| *pending*     | the request is pending approval |
| *in_progress* | the request is in progress      |
| *completed*   | the request has been completed  | 
| *cancelled*   | the request has been cancelled  |
| *denied*      |  the request has been denied    |

#### Reason

| status    | reason                      | description                                                                     |
|-----------|-----------------------------|---------------------------------------------------------------------------------|
| any       | *unknown* (default)         | the reason for the status is unknown/other                                      |
| completed | *executed*                  | the request has been executed                                                   |
| completed | *requested*                 | the request has been forwarded for execution                                    |
| denied    | *suspected_fraud*           | the request is suspected fraud                                                  |
| denied    | *insufficient_verification* | the [Data Subject](#Subject) has insufficiently completed identity verification |
| denied    | *no_match*                  | there is no match for the [Data Subject](#Subject)                              |
| denied    | *claim_not_covered*         | the claim is not covered                                                        |
| denied    | *outside_jurisdiction*      | the [Data Subject](#Subject) is outside the covered jurisdiction                |
| denied    | *too_many_requests*         | the [Data Subject](#Subject) has made too many requests                         |

<!-- | *need_user_verification*    | the status is pending because the [Data Subject](#Subject) needs to complete identity verification           | -->
