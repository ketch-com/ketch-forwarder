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

### SubscriptionControlSetting

The `SubscriptionControlSetting` object describes the settings of a `SubscriptionControl`.

```json
{
  "status": "granted"
}
```

#### Fields

| name             | required? | description                                             |
|------------------|-----------|---------------------------------------------------------|
| *status*         | yes       | [SubscriptionStatus](#SubscriptionStatus) of the topic. |

### SubscriptionTopicSetting

The `SubscriptionTopicSetting` object describes the settings of a `SubscriptionTopic`.

```json
{
  "contactMethods": ["email", "sms"],
  "status": "granted"
}
```

#### Fields

| name             | required? | description                                             |
|------------------|-----------|---------------------------------------------------------|
| *contactMethods* | no        | Array of granted contact methods                        |
| *status*         | yes       | [SubscriptionStatus](#SubscriptionStatus) of the topic. |

### SubscriptionStatus

The `SubscriptionStatus` enum identifies the status of a `SubscriptionTopic` or `SubscriptionControl`.

#### Values

| value     | description                                           |
|-----------|-------------------------------------------------------|
| `allowed` | The subscription topic/control is allowed by the user |
| `denied`  | The subscription topic/control is denied by the user  |
