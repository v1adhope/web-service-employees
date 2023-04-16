# Mongo and clean go

Variables are configured in compose and main.

```
http://localhost:8081/
```
Mongo express

# CRUD examle of operations

```
http://localhost:8090/v1/add/

{
    "name": "Vladislav",
    "surname": "Gardner",
    "phone": "1248901184",
    "companyID": 1,
    "passport": {
        "type": "D",
        "number": "48357419414"
    },
    "deportament": {
        "name": "404",
        "phone": "2390582350235"
    }
}

```
POST

```
http://localhost:8090/v1/delete/

{
    "id": "<ID>"
}
```
DELETE

```
http://localhost:8090/v1/bycompany/?id=1
http://localhost:8090/v1/bydeportament/?name=404
```
GET

```
http://localhost:8090/v1/update/

{
    "id":"643a6c8d8e97bc97d1eadcf9",
    "name": "Noname",
    "surname": "Nosurname",
    "phone": "unknown",
    "companyID": 2,
    "passport": {
        "type": "P",
        "number": "43612684524"
    },
    "deportament": {
        "name": "101",
        "phone": "unknown"
    }
}

```
PATCH
