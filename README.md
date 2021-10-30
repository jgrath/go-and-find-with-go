# go-lang Demonstrator Application

go-and-find-with-go is a prototype application demonstrating good use of 
go-lang language features and how to properly structure a testable web application.
It covers many language features and aspects around TDD found in a commercial environment.

#### READ THIS before you do anything!
This application is only for learning and demonstration purposes and shall not be used
for anything else and comes with absolutely no warranty or support of any kind whatsoever.

### Get a property entry

#### Request

  `GET /system-settings`
   
    curl -i -H 'Accept: application/json' http://localhost:8080/system-settings

#### Response
    HTTP/1.1 200 OK
    Date: Mon, 12 Apr 2021 14:55:13 GMT
    Content-Length: 280
    Content-Type: text/plain; charset=utf-8

    [{"property-name":"prop1","property-value":""...,"GroupDescription":"Group level Description"}]

* ***Error Response:***
    * **Code:** 500 <br />
      **Content:** `{ error : "Error reading from the database" }`

----
### Get properties for a group

#### Request
`GET /system-settings/group-code/{group-code}
OR
 GET /system-settings/group-name/{group-name}`

    curl -i -H 'Accept: application/json' http://localhost:8080/system-settings/group-code/main_group

#### Response
    HTTP/1.1 200 OK
    Date: Mon, 12 Apr 2021 15:11:32 GMT
    Content-Length: 280
    Content-Type: text/plain; charset=utf-8

    [{"property-name":"prop1","property-value":""...,"GroupDescription":"Group level Description"}]

* ***Error Response:***
    * **Code:** 500 <br />
      **Content:** `{ error : "Error reading from the database" }`

----
### Add property data as json

#### Request
    `POST /system-settings/`
    curl -i -H 'Accept: application/json' http://localhost:8080/system-settings/group-code/main_group

#### Response
    HTTP/1.1 200 OK
    Date: Mon, 12 Apr 2021 15:11:32 GMT
    Content-Length: 280
    Content-Type: text/plain; charset=utf-8

    [{"property-name":"prop1","property-value":""...,"GroupDescription":"Group level Description"}]

* ***Error Response:***
    * **Code:** 422 UNPROCESSABLE ENTITY  <br />
      **Content:** `{"validation failure"}`
OR
    * **Code:** 500 <br />
      **Content:** `{ error : "Error adding data to the database" }`

----
## Installation

GO needs to be installed. Here are the actual steps for running the application.

- rm go.mod (if already present)
- go mod init github.com/{you-username}/go-and-find-with-go
- go mod vendor
- go get
- cd test
- go test
- cd ..
- go run main.go (starts the http service)

## Mongo Schema
Two tables need to be created.

CREATE TABLE "SETTINGS_GROUP"
(
id integer,
group_code character varying COLLATE pg_catalog."default",
group_name character varying COLLATE pg_catalog."default",
description character varying COLLATE pg_catalog."default"
)

CREATE TABLE "SYSTEM_SETTINGS"
(
property_name character varying(100),
property_value character varying(100),
default_value character varying(100),
description character varying(100),
data_type character varying(100),
enabled int ,
active_from_date time,
group_code character varying(10)
)

## Contributing
Pull requests are welcome.
For major changes, please open an issue first to discuss your change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)