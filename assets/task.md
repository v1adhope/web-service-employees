# Task

Web-service of employees. Language: Golang.

The service must be able to:
- Add employees, the response should be Id of the added employee.
- Delete employees by Id.
- Output the list of employees for a specified company. All available fields.
- Print the list of employees for a specified department of a company. All available fields.
- Change an employee by his Id. Only those fields must be changed, that are specified in the request.

```
type Employee struct {
	Id        string
	Name      string
	Surname   string
	Phone     string
	CompanyId string
	Passport  struct {
		Type   string
		Number string
	}
	Deportment struct {
		Name  string
		Phone string
	}
}

```
Model

All methods must be implemented as HTTP requests in JSON format.

Database: any.

Status: Complete
