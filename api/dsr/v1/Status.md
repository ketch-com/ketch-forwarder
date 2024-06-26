# Status

The `status` of a Request progresses through the following state/activity diagram:

![](https://lucid.app/publicSegments/view/6ee0ee1c-cf45-49fd-9a09-eabf6eba8b5e/image.png)

Note that multiple Status events can be received with each status. For example, a callback may be called daily with
a status of `in_progress` until finally the final status is sent with status of `completed`.

If a `Response` does not include return a status of `completed`, then a subsequent `StatusEvent`
event MUST be sent to the callbacks with one of those statuses. Otherwise, the request will be marked as failed after
some time.

## Status code

| name          | description                     |
|---------------|---------------------------------|
| *unknown*     | the status is unknown           |
| *pending*     | the request is pending approval |
| *in_progress* | the request is in progress      |
| *completed*   | the request has been completed  |

The status code is case sensitve.

## Reason

| status    | reason                             | description                                                                                                        |
|-----------|------------------------------------|--------------------------------------------------------------------------------------------------------------------|
| any       | *unknown* (default)                | the reason for the status is unknown/other                                                                         |
| pending   | *need_user_verification*           | the status is pending because the [Data Subject](README.md#Subject) needs to complete identity verification        |
| pending   | *pending*                          | the request is pending                                                                                             |
| completed | *requested*                        | the request has been forwarded for execution                                                                       |
| completed | *no_match*                         | the request was completed, but there is no match for the [Data Subject](README.md#Subject)                         |
| completed | *insufficient_identification*      | the [Data Subject](README.md#Subject) was not sufficiently identified                                              |
| completed | *executed*                         | the request has been executed                                                                                      |
| completed | *executed_direct_subject_delivery* | the request completed and data will be delivered directly to [Data Subject](README.md#Subject) email or app portal |
