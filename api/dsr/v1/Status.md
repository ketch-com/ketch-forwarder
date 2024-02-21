# Status

The `status` of a Request progresses through the following state/activity diagram:

![](https://lucid.app/publicSegments/view/61b83862-bbc8-41b1-bdae-d92b7bf87af6/image.png)

Note that multiple Status events can be received with each status. For example, a callback may be called daily with
a status of `in_progress` until finally the final status is sent with status of `completed`.

If a `Response` does not include return a status of `completed`, `cancelled` or `denied`, then a subsequent `StatusEvent`
event MUST be sent to the callbacks with one of those statuses. Otherwise, the request will be marked as failed after
some time.

## Status code

| name          | description                     |
|---------------|---------------------------------|
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
| denied    | *no_match*                         | there is no match for the [Data Subject](README.md#Subject)                                                        |
| denied    | *insufficient_identification*      | the [Data Subject](README.md#Subject) was not sufficiently identified                                              |
| denied    | *insufficient_verification*        | the [Data Subject](README.md#Subject) has insufficiently completed identity verification                           |
| denied    | *claim_not_covered*                | the claim is not covered                                                                                           |
| denied    | *outside_jurisdiction*             | the [Data Subject](README.md#Subject) is outside the covered jurisdiction                                          |
| denied    | *too_many_requests*                | the [Data Subject](README.md#Subject) has made too many requests                                                   |
| denied    | *suspected_fraud*                  | the request is suspected fraud                                                                                     |
| denied    | *invalid_credentials*              | the request failed due to invalid credentials                                                                      |
| denied    | *insufficient_permission*          | the request failed due to insufficient permission                                                                  |
| denied    | *internal_app_error*               | the request failed due to an internal app error                                                                    |
| denied    | *sla_expiry*                       | the request failed due to SLA expiry                                                                               |
