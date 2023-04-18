# Mongo and clean go

[task](/assets/task.md)

Variables are configured in compose and main.

```
http://localhost:8081/
```
Mongo express

# CRUD examle of operations

```
http://localhost:8090/v1/employee/

{
    "name": "Vladislav",
    "surname": "Gardner",
    "phone": "1248901184",
    "companyID": "1",
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
http://localhost:8090/v1/employee/

{
    "id": "<ID>"
}
```
DELETE

```
http://localhost:8090/v1/employees-by-company/?id=1
http://localhost:8090/v1/employees-by-deportament/?name=404
```
GET

```
http://localhost:8090/v1/employee/

{
    "id":"643a6c8d8e97bc97d1eadcf9",
    "name": "Noname",
    "surname": "Nosurname",
    "phone": "unknown",
    "companyID": "2",
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
