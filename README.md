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

For the examples, the endpoint is configured as follows (note this URL does not really exist):

* URL: `https://example.com/endpoint`
* Header Key: `Authorization`
* Header Value: `Bearer $auth`

For the examples, a callback could be configured as follows (note this URL does not really exist):

* URL: `https://dsr.ketch.com/callback`
* Header Key: `Authorization`
* Header Value: `Bearer $auth`

For the examples, a result could be configured as follows (note this URL does not really exist):

* URL: `https://example.com/results`
* Header Key: `Authorization`
* Header Value: `Bearer $auth`

## Data Subject Requests

* [Access](dsr/v1/Access.md)
* [Correction](dsr/v1/Correction.md)
* [Delete](dsr/v1/Delete.md)
* [Restrict Processing](dsr/v1/RestrictProcessing.md)
* [Augment](dsr/v1/Augment.md)

## Consent

* [Update Consent](consent/v1/UpdateConsent.md)
