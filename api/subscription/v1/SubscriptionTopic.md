# SubscriptionTopic

A SubscriptionTopic Request is initiated when a Data Subject subscribes to or unsubscribes from a subscription topic.

## SubscriptionTopic Request

```http request
POST /endpoint HTTP/1.1
Host: www.example.com
Content-Type: application/json
Accept: application/json
Authorization: $auth

{
  "apiVersion": "subscription/v1",
  "kind": "SubscriptionTopicRequest",
  "metadata": {
    "uid": "22880925-aac5-42f9-a653-cb6921d361ff",
    "tenant": "axonic"
  },
  "request": {
    "controller": "axonic",
    "property": "axonic.io",
    "environment": "production",
    "identities": [
      {
        "identitySpace": "account_id",
        "identityFormat": "raw",
        "identityValue": "123"
      }
    ],
    "topics": {
      "competitions": {
        "email": {
          "status": "granted"
        },
        "sms": {
          "status": "denied"
        }
      },
      "marketing": {
        "email": {
          "status": "denied"
        },
        "sms": {
          "status": "denied"
        }
      }
    },
    "submittedTimestamp": 123
  }
}
```

### Fields

| name                         | required? | description                                                                                                                            |
|------------------------------|-----------|----------------------------------------------------------------------------------------------------------------------------------------|
| *apiVersion*                 | yes       | API version. Must be `subscription/v1`                                                                                                 |
| *kind*                       | yes       | Message kind. Must be `SubscriptionTopicRequest`                                                                                       |
| *metadata*                   | yes       | [Metadata](../../runtime/v1/Metadata.md) object                                                                                        |
| *request.controller*         | no        | Code of the Ketch controller tenant. Only supplied if the ultimate controller is different to the `metadata.tenant`                    |
| *request.property*           | yes       | Code of the digital property defined in Ketch                                                                                          |
| *request.environment*        | yes       | Code environment defined in Ketch                                                                                                      |
| *request.identities*         | yes       | Array of [Identities](README.md#Identity)                                                                                              |
| *request.topics*             | yes       | Map of subscription topics codes mapped to a map of channels mapped to [SubscriptionTopicSetting](README.md#SubscriptionTopicSetting). |
| *request.submittedTimestamp* | yes       | UNIX timestamp in seconds                                                                                                              |

## Consent Response

A successful response SHOULD return either `202 Accepted` or `204 No Content` response status code.

```http request
HTTP/1.1 202 Accepted
```

or

```http request
HTTP/1.1 204 No Content
```
