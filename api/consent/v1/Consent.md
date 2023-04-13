# Consent

A Consent Request is initiated when a Data Subject specifies or changes consent preferences.

## Consent Request

```http request
POST /endpoint HTTP/1.1
Host: www.example.com
Content-Type: application/json
Accept: application/json
Authorization: $auth

{
  "apiVersion": "consent/v1",
  "kind": "ConsentRequest",
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
    "purposes": {
      "advertising": "granted"
      "data_sales": "granted"
      "email_mktg": "denied"
    },
    "legalBasis": {
      "advertising": "consent_optin",
      "data_sales": "consent_optout",
      "email_mktg": "disclosure"
    },
    "vendors": [
      "79",
    ],
    "context": {
      "account_id": "123"
    },
    "collectedAt": 12345984398
  }
}
```

### Fields

| name                   | required? | description                                                                                                         |
|------------------------|-----------|---------------------------------------------------------------------------------------------------------------------|
| *apiVersion*           | yes       | API version. Must be `dsr/v1`                                                                                       |
| *kind*                 | yes       | Message kind. Must be `ConsentRequest`                                                                              |
| *metadata*             | yes       | [Metadata](../../runtime/v1/Metadata.md) object                                                                     |
| *request.controller*   | no        | Code of the Ketch controller tenant. Only supplied if the ultimate controller is different to the `metadata.tenant` |
| *request.property*     | yes       | Code of the digital property defined in Ketch                                                                       |
| *request.environment*  | yes       | Code environment defined in Ketch                                                                                   |
| *request.regulation*   | yes       | Code of the regulation defined in Ketch                                                                             |
| *request.jurisdiction* | yes       | Code of the jurisdiction defined in Ketch                                                                           |
| *request.identities*   | yes       | Array of [Identities](../../dsr/v1/README.md#Identity)                                                              |
| *request.purposes*     | yes       | Map of [ConsentStatus](#ConsentStatus). The key is the purpose code.                                                |
| *request.legalBasis*   | yes       | Map of [LegalBasis](#LegalBasis) for the purposes.                                                                  |
| *request.context*      | no        | Map containing additional context that have been added via identity verification or other augmentation methods      |
| *request.collectedAt*  | yes       | UNIX Timestamp of when the consent was collected                                                                    |

## Consent Response

A successful response SHOULD return the `204 No Content` response status code.

```http request
HTTP/1.1 204 No Content
```

### ConsentStatus

The `ConsentStatus` enum identifies the status of a `Consent` purpose.

#### Values

| value     | description                           |
|-----------|---------------------------------------|
| `granted` | The purpose is allowed by the user    |
| `denied`  | The purpose is disallowed by the user |

### LegalBasis

The `LegalBasis` enum identifies the legal basis of a `Consent` purpose.

#### Values

| value            | name              |
|------------------|-------------------|
| `consent_optin`  | Consent - Opt In  |
| `consent_optout` | Consent - Opt Out |
| `disclosure`     | Disclosure        |
| `other`          | Other legal basis |
