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

A Delete request is initiated when a Data Subject selects a right that allows for deleting of personal data.

To forward a Data Subject Request, Ketch sends a message using the POST method to the configured endpoint. The format of 
the message and expected responses depend on the type of right invoked by the Data Subject.

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

* *apiVersion* - must be `dsr/v1`
* *kind* - must be `DeleteRequest`
* *metadata.uid* - will be a unique UUIDv4, and uniquely identifies the request
* *metadata.tenant* - will be the Ketch tenant code where the request originated
* *request.controller* - code of the Ketch controller tenant. Only supplied if the ultimate controller is different to the `metadata.tenant`
* *request.property* - code of the digital property defined in Ketch
* *request.environment* - code environment defined in Ketch
* *request.regulation* - code of the regulation defined in Ketch
* *request.jurisdiction* - code of the jurisdiction defined in Ketch
* *request.identities* - array of [Identities](#Identity)
* *request.callbacks* - array of [Callbacks](#Callback)
* *request.subject* - the [Data Subject](#Subject)
* *request.claims* - map containing additional non-identity claims that have been added via identity verification or other
  augmentation methods. Identity claims should be included in `request.identities`.
* *request.submittedTimestamp* - UNIX timestamp in seconds
* *request.dueTimestamp* - UNIX timestamp in seconds

#### Callback

* *url* - URL of the callback endpoint
* *headers* - map of headers to send to the callback endpoint

#### Identity

* *identitySpace* - identity space code setup in Ketch
* *identityFormat* - format of the identity value (`raw`, `md5`, `sha1`)
* *identityValue* - value of the identity

#### Subject

* *email* - email address provided by the Data Subject
* *firstName* - first name provided by the Data Subject
* *lastName* - last name provided by the Data Subject
* *addressLine1* - address line 1 provided by the Data Subject
* *addressLine2* - address line 2 provided by the Data Subject
* *city* - city provided by the Data Subject
* *stateRegionCode* - state/region code (e.g., CA) provided by the Data Subject
* *postalCode* - zip/postal code provided by the Data Subject
* *countryCode* - two-character ISO country code (e.g., US) provided by the Data Subject
* *description* - free-text description provided by the Data Subject

### Successful Delete Response

Headers:
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
    "status": "pending",
    "reason": "need_user_verification",
    "expectedCompletionTimestamp": 123,
    "redirectUrl": "https://verifyidentity.com/123"
  }
}
```

#### Fields

* *apiVersion* - must be `dsr/v1`
* *kind* - must be `DeleteResponse`
* *metadata.uid* - will be a unique UUIDv4, and uniquely identifies the request
* *metadata.tenant* - will be the Ketch tenant code where the request originated
* *response.status* - the [status](#Status) of the Data Subject Request
* *response.reason* - the [reason](#Reason) for the status of the Data Subject Request
* *response.expectedCompletionTimestamp* - the UNIX time stamp at which the Data Subject Request is expected to be completed
* *response.redirectUrl* - if the Data Subject should be redirected to a URL (perhaps for confirmation), this field specifies the URL to redirect the Data Subject to

### Delete Error Response

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

* *apiVersion* - must be `dsr/v1`
* *kind* - must be `Error`
* *metadata.uid* - will be a unique UUIDv4, and uniquely identifies the request
* *metadata.tenant* - will be the Ketch tenant code where the request originated
* *error.code* - the HTTP status code
* *error.status* - a string code representing the error
* *error.message* - a user-friendly error message (e.g., `"Not found"`)

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

* *apiVersion* - must be `dsr/v1`
* *kind* - must be `DeleteStatusEvent`
* *metadata.uid* - will be a unique UUIDv4, and uniquely identifies the request
* *metadata.tenant* - will be the Ketch tenant code where the request originated
* *response.status* - the [status](#Status) of the Data Subject Request
* *response.reason* - the [reason](#Reason) for the status of the Data Subject Request
* *response.expectedCompletionTimestamp* - the UNIX time stamp at which the Data Subject Request is expected to be completed

## Access

An Access Request is initiated when a Data Subject invokes a right that allows Access/Portability of personal data.

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

* *apiVersion* - must be `dsr/v1`
* *kind* - must be `AccessRequest`
* *metadata.uid* - will be a unique UUIDv4, and uniquely identifies the request
* *metadata.tenant* - will be the Ketch tenant code where the request originated
* *request.controller* - code of the Ketch controller tenant. Only supplied if the ultimate controller is different to the `metadata.tenant`
* *request.property* - code of the digital property defined in Ketch
* *request.environment* - code environment defined in Ketch
* *request.regulation* - code of the regulation defined in Ketch
* *request.jurisdiction* - code of the jurisdiction defined in Ketch
* *request.identities* - array of [Identities](#Identity)
* *request.callbacks* - array of [Callbacks](#Callback)
* *request.subject* - the [Data Subject](#Subject)
* *request.claims* - map containing additional claims that have been added via identity verification or other augmentation methods
* *request.submittedTimestamp* - UNIX timestamp in seconds
* *request.dueTimestamp* - UNIX timestamp in seconds

#### Callback

* *url* - URL of the callback endpoint
* *headers* - map of headers to send to the callback endpoint

#### Identity

* *identitySpace* - identity space code setup in Ketch
* *identityFormat* - format of the identity value (`raw`, `md5`, `sha1`)
* *identityValue* - value of the identity

#### Subject

* *email* - email address provided by the Data Subject
* *firstName* - first name provided by the Data Subject
* *lastName* - last name provided by the Data Subject
* *addressLine1* - address line 1 provided by the Data Subject
* *addressLine2* - address line 2 provided by the Data Subject
* *city* - city provided by the Data Subject
* *stateRegionCode* - state/region code (e.g., CA) provided by the Data Subject
* *postalCode* - zip/postal code provided by the Data Subject
* *countryCode* - two-character ISO country code (e.g., US) provided by the Data Subject
* *description* - free-text description provided by the Data Subject

### Successful Access Response

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
    "status": "pending",
    "reason": "need_user_verification",
    "expectedCompletionTimestamp": 123,
    "redirectUrl": "https://verifyidentity.com/123",
    "results": []
  }
}
```

#### Fields

* *apiVersion* - must be `dsr/v1`
* *kind* - must be `AccessResponse`
* *metadata.uid* - will be a unique UUIDv4, and uniquely identifies the request
* *metadata.tenant* - will be the Ketch tenant code where the request originated
* *response.status* - the [status](#Status) of the Data Subject Request
* *response.reason* - the [reason](#Reason) for the status of the Data Subject Request
* *response.expectedCompletionTimestamp* - the UNIX time stamp at which the Data Subject Request is expected to be completed
* *response.redirectUrl* - if the Data Subject should be redirected to a URL (perhaps for confirmation)
* *response.results* - array of [Callbacks](#Callback) that can be used to download the contents requested

### Access Error Response

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

* *apiVersion* - must be `dsr/v1`
* *kind* - must be `Error`
* *metadata.uid* - will be a unique UUIDv4, and uniquely identifies the request
* *metadata.tenant* - will be the Ketch tenant code where the request originated
* *error.code* - the HTTP status code
* *error.status* - a string code representing the error
* *error.message* - a user-friendly error message (e.g., `"Not found"`)

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

* *apiVersion* - must be `dsr/v1`
* *kind* - must be `AccessStatusEvent`
* *metadata.uid* - will be a unique UUIDv4, and uniquely identifies the request
* *metadata.tenant* - will be the Ketch tenant code where the request originated
* *response.status* - the [status](#Status) of the Data Subject Request
* *response.reason* - the [reason](#Reason) for the status of the Data Subject Request
* *response.expectedCompletionTimestamp* - the UNIX time stamp at which the Data Subject Request is expected to be completed
* *response.results* - array of [Callbacks](#Callback) that can be used to download the contents requested

## Restrict Processing

A Restrict Processing Request is initiated when a Data Subject invokes a right that allows restriction of processing
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

* *apiVersion* - must be `dsr/v1`
* *kind* - must be `RestrictProcessingRequest`
* *metadata.uid* - will be a unique UUIDv4, and uniquely identifies the request
* *metadata.tenant* - will be the Ketch tenant code where the request originated
* *request.controller* - code of the Ketch controller tenant. Only supplied if the ultimate controller is different to the `metadata.tenant`
* *request.property* - code of the digital property defined in Ketch
* *request.environment* - code environment defined in Ketch
* *request.regulation* - code of the regulation defined in Ketch
* *request.jurisdiction* - code of the jurisdiction defined in Ketch
* *request.purposes* - list of purpose codes defined in Ketch
* *request.identities* - array of [Identities](#Identity)
* *request.callbacks* - array of [Callbacks](#Callback)
* *request.subject* - the [Data Subject](#Subject)
* *request.claims* - map containing additional claims that have been added via identity verification or other augmentation methods
* *request.submittedTimestamp* - UNIX timestamp in seconds
* *request.dueTimestamp* - UNIX timestamp in seconds

#### Callback

* *url* - URL of the callback endpoint
* *headers* - map of headers to send to the callback endpoint

#### Identity

* *identitySpace* - identity space code setup in Ketch
* *identityFormat* - format of the identity value (`raw`, `md5`, `sha1`)
* *identityValue* - value of the identity

#### Subject

* *email* - email address provided by the Data Subject
* *firstName* - first name provided by the Data Subject
* *lastName* - last name provided by the Data Subject
* *addressLine1* - address line 1 provided by the Data Subject
* *addressLine2* - address line 2 provided by the Data Subject
* *city* - city provided by the Data Subject
* *stateRegionCode* - state/region code (e.g., CA) provided by the Data Subject
* *postalCode* - zip/postal code provided by the Data Subject
* *countryCode* - two-character ISO country code (e.g., US) provided by the Data Subject
* *description* - free-text description provided by the Data Subject

### Successful Restrict Processing Response

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
    "reason": "need_user_verification",
    "expectedCompletionTimestamp": 123,
    "redirectUrl": "https://verifyidentity.com/123",
    "results": []
  }
}
```

#### Fields

* *apiVersion* - must be `dsr/v1`
* *kind* - must be `RestrictProcessingResponse`
* *metadata.uid* - will be a unique UUIDv4, and uniquely identifies the request
* *metadata.tenant* - will be the Ketch tenant code where the request originated
* *response.status* - the [status](#Status) of the Data Subject Request
* *response.reason* - the [reason](#Reason) for the status of the Data Subject Request
* *response.expectedCompletionTimestamp* - the UNIX time stamp at which the Data Subject Request is expected to be completed
* *response.redirectUrl* - if the Data Subject should be redirected to a URL (perhaps for confirmation)
* *response.results* - array of [Callbacks](#Callback) that can be used to download the contents requested

### Restrict Processing Error Response

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

* *apiVersion* - must be `dsr/v1`
* *kind* - must be `Error`
* *metadata.uid* - will be a unique UUIDv4, and uniquely identifies the request
* *metadata.tenant* - will be the Ketch tenant code where the request originated
* *error.code* - the HTTP status code
* *error.status* - a string code representing the error
* *error.message* - a user-friendly error message (e.g., `"Not found"`)

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

* *apiVersion* - must be `dsr/v1`
* *kind* - must be `RestrictProcessingStatusEvent`
* *metadata.uid* - will be a unique UUIDv4, and uniquely identifies the request
* *metadata.tenant* - will be the Ketch tenant code where the request originated
* *response.status* - the [status](#Status) of the Data Subject Request
* *response.reason* - the [reason](#Reason) for the status of the Data Subject Request
* *response.expectedCompletionTimestamp* - the UNIX time stamp at which the Data Subject Request is expected to be completed
* *response.results* - array of [Callbacks](#Callback) that can be used to download the contents requested

## Status

The `status` of a Request progresses through the following state/activity diagram:

![](https://lucid.app/publicSegments/view/61b83862-bbc8-41b1-bdae-d92b7bf87af6/image.png)

### Status code

* *unknown* - the status is unknown
* *pending* - the request is pending approval
* *in_progress* - the request is in progress
* *completed* - the request has been completed
* *cancelled* - the request has been cancelled
* *denied* - the request has been denied

### Reason

* *unknown* - the reason for the status is unknown
* *need_user_verification* - the status is pending because the Data Subject needs to complete identity verification
* *suspected_fraud* - the status is denied because the request is suspected fraud
* *insufficient_verification* - the status is denied because the Data Subject has insufficiently completed identity verification
* *no_match* - the status is cancelled/denied because there is no match for the Data Subject
* *claim_not_covered* - the status is cancelled/denied because the claim is not covered
* *outside_jurisdiction* - the status is cancelled/denied because the Data Subject is outside the covered jurisdiction
* *too_many_requests* - the status is cancelled/denied because the Data Subject has made too many requests
* *other* - the reason for the status some other reason 
