# go-timeservice
A simple API to return the current time.


## Simple usage
This API has a single endpoint: `/time`.  Visiting that endpoint will return the
current time in UnixTime as [defined in the Go time package](https://golang.org/pkg/time/#pkg-constants) in UTC.

## Query Parameters
The endpoint accepts two arguments passed as query params, the first is
`timezone`.  The second is `format`.

### timezone parameter
The `timezone` parameter should be passed as the timezone
is found in the [tz database](https://en.wikipedia.org/wiki/Tz_database).  For example `/time?timezone=America/New_York`.
If not present, the API will return time in UTC.

### format parameter
The `format` parameter should either be `unix` or `timestamp`.  For example:
`/time?format=unix` or `/time?format=timestamp`.  Passing `unix` is the default
value and will return the date in UnixTime as defined in the Go time package.
Passing timestamp will return the time as the current unix timestap.
