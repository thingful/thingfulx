# thingfulx Change Log

## 0.2.0.alpha 2016-09-11

* Add structured Metadata in the form of Tags.
* Compose Fetcher interface of two smaller interfaces: ThingFetcher and
  URLSFetcher
* Remove json struct tags as we don't use this currently

## 0.1.7 2016-09-09

* Switch tests to use thingful/httpmock.

## 0.1.6 2016-08-18

* Add `RawData` attribute to Things.
* Switch tests to using stretchr/testify.

## 0.1.5 2016-08-17

* Remove vendored libraries - switch to using glide in unvendored dependencies
  mode.

## 0.1.4 2016-08-16

* Add context.Context parameter to Fetch allowing request specific data to be
  passed into the main Fetch method.

## 0.1.3 2016-06-30

* Fix bug where URl parameter wrongly marked as being of the Client type.

## 0.1.2 2016-06-01

* Add standalone Client type to provide HTTP requesting functionality.
* Remove unused Response type.

## 0.1.1 2016-04-04

* Remove dataurl attribute from Provider.
* Add specific error type for not found errors.
* Add time provider constructors.

## 0.1.0 2016-02-03

* Initial release version.
