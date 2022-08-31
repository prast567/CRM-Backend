# CRM-Backend
# Customer Relation Management System

This is the Customer Relationship Management System's backend application. 


### Run

cd into project directory(i.e from the /crm directory)

```shell
go run .\cmd\api\main.go   
```

### Testing

cd into project directory and

```shell

go test ./... -v -coverprofile=coverage.out 
```

### Build

cd into project directory and

```shell

go build -o ./bin/ ./cmd/...
```



# API Documentation



## Server:

[localhost:3000](http://localhost:3000/)

## Endpoints example

- GET  http://localhost:3000/customers 
- GET  http://localhost:3000/customers/1
- POST http://localhost:3000/customers
	   json data :- {
                      "name": "Testing",
                      "email": "test@yahoo.com",
                      "phone": "+14155552671",
                      "role": "test Customer",
                      "contacted": true
                     }
- PUT  http://localhost:3000/customers/3
	    json data :- {
                      "name": "Testing_update",
                      "email": "test_update@yahoo.com",
                      "phone": "+14155552671",
                      "role": "Update Customer",
                      "contacted": true
                     }
- DELETE http://localhost:3000/customers/3


## Customer

### `Get` `/Customers`

Responses with a array of all customer in system

#### Responses

##### Body

##### `200`

###### Schema:

- **status:** `string` 

- **message:** `string` *[optional]*

- **data:**  `array`**of:** `object`
  
   Customer contains all the information of a customer
  
  - **id:** `integer`
  
  - **name:** `string` (5-127 character)
  
  - **role:** `string`(1-15 characters)
  
  - **email:** `string` (valid email)
  
  - **name:** `string` (e164 format)
  
  - **contacted:** `boolean`

###### Example:

```json
{
  "status": "success",
  "data": [
    {
      "id": 1,
      "name": "Mr. Robot",
      "role": "Customer",
      "email": "robot@gmail.com",
      "phone": "+8801814567342",
      "contacted": true
    },
    {
      "id": 2,
      "name": "Mr. John",
      "role": "Investor",
      "email": "john@yahoo.com",
      "phone": "+14155552671",
      "contacted": false
    },
    {
      "id": 3,
      "name": "James Mark",
      "role": "Investor",
      "email": "mark@facebook.com",
      "phone": "+14155552671",
      "contacted": true
    }
  ]
}
```

### `POST` `/Customers`

Creates a new customer in the system

#### Request

##### Body:

###### Schema

- **name:** `string` (5-127 character)

- **email:** `string` (valid email)

- **role:** `string`(1-15 characters)

- **contacted:** `boolean` *[Optional]*

###### Example

```json
{
  "name": "Mr. Robot",
  "email": "robot@yahoo.com",
  "phone": "+14155552671",
  "role": "Customer",
  "contacted": true
}
```

#### Responses

##### `200` Success

###### Schema:

- **status:** `string`

- **message:** `string` *[optional]*

###### Example:

```json
{
    "status": "success"
}
```

##### `404`  Invalid Input Body

###### Schema:

- **status:** `string`

- **message:** `string` *[optional]*

###### Example:

```json
{
    "status": "failed",
    "message": "failed to get input body. error:EOF"
}
```

##### `422` Input Field Validation Failed

###### Schema:

- **status:** `string`

- **message:** `string` *[optional]*

- **data:** `map[string]string`
  
  Input validation failed messages

###### Example:

```json
{
    "status": "failed",
    "message": "failed to validate customer input.",
    "data": {
        "email": "must be a valid email address",
        "role": "cannot be blank"
    }
}
```

### `Get` `/Customers/{id}`

Get a single customer from the system with the given id in url parameter

#### Request

##### URL Parameters

- **id:** `string`
  
  The id of the customer that you want to view

#### Responses

##### `200` Success

###### Schema:

- **status:** `string`

- **message:** `string` *[optional]*

- **data:** `object`
  
  Customer contains all the information of a customer
  
  - **id:** `integer`
  
  - **name:** `string` (5-127 character)
  
  - **role:** `string`(1-15 characters)
  
  - **email:** `string` (valid email)
  
  - **name:** `string` (e164 format)
  
  - **contacted:** `boolean`

###### Example:

```json
{
  "status": "success",
  "data": {
    "id": 3,
    "name": "James Mark",
    "role": "Investor",
    "email": "mark@facebook.com",
    "phone": "+14155552671",
    "contacted": true
  }
}
```

##### `404` Not Found

###### Schema:

- **status:** `string`

- **message:** `string` *[optional]*

###### Example:

```json
{
    "status": "failed",
    "message": "user not found"
}
```

### `PUT` `/Customers/{id}`

Updates an exisiting customer in the system

#### Request

##### URL Parameters

- **id:** `string`
  
  The id of the customer that you want to view

##### Body:

###### Schema

- **name:** `string` (5-127 character)

- **email:** `string` (valid email)

- **role:** `string`(1-15 characters)

- **contacted:** `boolean` *[Optional]*

###### Example

```json
{
  "name": "Mr. Robot",
  "email": "robot@yahoo.com",
  "phone": "+14155552671",
  "role": "Customer",
  "contacted": true
}
```

#### Responses

##### `200` Success

###### Schema:

- **status:** `string`

- **message:** `string` *[optional]*

###### Example:

```json
{
    "status": "success"
}
```

##### `404` Invalid Input Body

###### Schema:

- **status:** `string`

- **message:** `string` *[optional]*

###### Example:

```json
{
    "status": "failed",
    "message": "failed to get input body. error:EOF"
}
```

##### `422` Input Field Validation Failed

###### Schema:

- **status:** `string`

- **message:** `string` *[optional]*

- **data:** `map[string]string`
  
  Input validation failed messages

###### Example:

```json
{
    "status": "failed",
    "message": "failed to validate customer input.",
    "data": {
        "email": "must be a valid email address",
        "role": "cannot be blank"
    }
}
```

### `DELETE` `/Customers/{id}`

Deletes an exisiting customer from the system

#### Request

##### URL Parameters

- **id:** `string`
  
  The id of the customer that you want to view

#### Responses

##### `200` Success

###### Schema:

- **status:** `string`

- **message:** `string` *[optional]*

###### Example:

```json
{
    "status": "success"
}
```

##### `404` Not Found

###### Schema:

- **status:** `string`

- **message:** `string` *[optional]*

###### Example:

```json
{
    "status": "failed",
    "message": "user not found"
}
```
