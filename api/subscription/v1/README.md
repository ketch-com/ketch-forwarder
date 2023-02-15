# Subscription Requests

## Request Types

* [SubscriptionTopic](SubscriptionTopic.md)
* [SubscriptionControl](SubscriptionControl.md)

## Common Objects

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

### SubscriptionStatus

The SubscriptionStatus enum identifies the status of a SubscriptionTopic or SubscriptionControl.

#### Values

| value     | description                                           |
|-----------|-------------------------------------------------------|
| `allowed` | The subscription topic/control is allowed by the user |
| `denied`  | The subscription topic/control is denied by the user  |
