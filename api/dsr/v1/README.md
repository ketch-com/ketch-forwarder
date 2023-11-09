# Data Subject Requests

## Request Types

* [Access](Access.md)
* [Correction](Correction.md)
* [Delete](Delete.md)
* [Restrict Processing](RestrictProcessing.md)

## Response Status

Each response can have a [status](Status.md) that communicates to the Ketch platform how to handle the activity. The
status is case-sensitive.

## Common objects

### Callback

The Callback object defines a URL and associated headers to be used for communicating information.
The protocol for communicating with the Callback is defined for every object that uses the Callback.

```json
{
  "url": "https://dsr.ketch.com/callback",
  "headers": {
    "Authorization": "$auth"
  }
}
```

#### Fields

| name      | required? | description                                     |
| --------- | --------- | ----------------------------------------------- |
| *url*     | yes       | URL of the callback endpoint                    |
| *headers* | no        | map of headers to send to the callback endpoint |

### Document

The Document object defines a way of providing document data to the Ketch platform.

#### As a Callback

The Document object can look like a [Callback](#Callback) object which allows Ketch to download
the document using a simple HTTP GET.

```json
{
  "url": "https://dsr.controller.com/get/my/document",
  "headers": {
    "Authorization": "$auth"
  }
}
```

See [Callback](#Callback) for more details of this Document subtype.

#### As an embedded object

The Document object can provide the data directly in the value.

```json
{
  "data": "standard-base64-encoded-data",
  "headers": {
    "Content-Type": "application/json"
  }
}
```

Here, the document is returned directly in the response/event using base64-encoded data. The
`Content-Type` header is required so Ketch knows the [mime type](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types/Common_types)
of data received. The only supported content types are `application/json` and `application/pdf`. File size is
limited to 3.5Mb.

#### JSON Data

JSON document data is combined across all callback requests for a given DSR request using merge patches as defined 
in [RFC7396](https://www.rfc-editor.org/rfc/rfc7396). The total response data across all callbacks is
limited to 1MB.

#### Fields

| name                   | required? | description                  |
| ---------------------- | --------- | ---------------------------- |
| *data*                 | yes       | Standard base64-encoded data |
| *headers*              | yes       | map of headers for the data  |
| *headers.Content-Type* | yes       | mime type of the data        |

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
| ---------------- | --------- | ---------------------------------------------------------------------- |
| *identitySpace*  | yes       | Identity space code setup in Ketch                                     |
| *identityFormat* | no        | Format of the identity value (`raw`, `md5`, `sha1`). Default is `raw`. |
| *identityValue*  | yes       | Value of the identity                                                  |

### Subject

The Subject object describes the Data Subject making the request. The fields type, email, city, and description are 
read only. Any empty subject values will be ignored when patching the subject. Additional properties may exist
in this object depending on the subject type.

```json
{
  "type": "employee",
  "email": "test@subject.com",
  "firstName": "Test",
  "lastName": "Subject",
  "addressLine1": "123 Main St",
  "addressLine2": "",
  "city": "Anytown",
  "stateRegionCode": "MA",
  "postalCode": "10123",
  "countryCode": "US",
  "description": "Restrict my data",
  "formData": {
    "customFormField1": "value1",
    "customFormField2": "value1"
  }
}
```

#### Fields

| name              | required? | description                                                                     |
| ----------------- | --------- | ------------------------------------------------------------------------------- |
| *type*            | no        | Type of the data subject (customer, employee, etc.)                             |
| *email*           | yes       | Email address provided by the Data Subject                                      |
| *firstName*       | yes       | First name provided by the Data Subject                                         |
| *lastName*        | yes       | Last name provided by the Data Subject                                          |
| *addressLine1*    | no        | Address line 1 provided by the Data Subject                                     |
| *addressLine2*    | no        | Address line 2 provided by the Data Subject                                     |
| *city*            | no        | City provided by the Data Subject                                               |
| *stateRegionCode* | no        | State/region code (e.g., CA) provided by the Data Subject                       |
| *postalCode*      | no        | Zip/postal code provided by the Data Subject                                    |
| *countryCode*     | no        | Two-character ISO country code (e.g., US) provided by the Data Subject          |
| *description*     | no        | Free-text description provided by the Data Subject                              |
| *formData*        | no        | Map containing additional form data that have been added via custom Form Fields |

## Response Augmentation

All Response and Event objects support augmenting the Data Subject Request, whether adding context variables
to the workflow, updating the data subject, mutating identities, adding results/internal documents or adding
messages to the request conversation.

### Add context variables

To add or update context variables, return the `context` property. Context is a map from the context variable 
code (string) to the value (string, integer or boolean).

### Update data subject

Map containing additions or changes to subject values [Data Subject](README.md#Subject). To add or update data subject
values, set the key and value.

### Mutate identities

To add identities, return additional [Identities](README.md#Identity) in the `identities` property.

### Add results

To add results that should be made available to the Data Subject, return result documents in the `results` property.

### Add documents

To add documents that should be available internally to a Ketch operator, but not the Data Subject, return result documents
in the `documents` property.

### Add messages to conversation

To add messages to the Data Subject Request conversation, return an array of JSON Messages in the `messages` property.
