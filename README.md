# Errors Service

### STAGE API ENDPOINT IS ``8wa5ijraug.execute-api.eu-west-1.amazonaws.com``
### PROD API ENDPOINT IS ````


### Timeout url

* url ``https://{API ENDPOINT}/Prod/timeout``

POST request

Headers:

* Content-Type : application/json

### InvalidToken url

* url ``https://{API ENDPOINT}/Prod/invalidtoken``

POST request

Headers:

* Content-Type : application/json

Possible errorCodes:

* InvalidAccessTokenClientError


### Non ok url

* url ``https://{API ENDPOINT}/Prod/nonok``

POST request

Headers:

* Content-Type : application/json

Returns 503

### TooOldAppVersionClientError url

* url ``https://{API ENDPOINT}/Prod/old_version``

POST request

Headers:

* Content-Type : application/json

Possible errorCodes:

* TooOldAppVersionClientError


old_version