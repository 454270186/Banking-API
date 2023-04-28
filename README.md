## Banking backend API

**Based on**
- Hexagonal Architecture
- Postgresql
- Gin

练习gin的玩具

### API

- GET ```/customer```: Get all customer

- GET ```/customer/:id```: Get customer by id

- POST ```/customer/:id/account```: Create a new account for customer with given id
    - POST Request Body<br>
    ```json
    {
        "account_type": "xxx",
        "amount": xxx
    }
    ```
    - Response
    ```json
    {
        "account_id": xxx
    }
    ```

- POST ```/customer/:id/account/:account_id```: Make a transaction for given account. 
<br>
Supporting transaction type: 1. withdrawal 2. deposit


### Validation

**Using jwt to authenticate users**

- Enable users to only operate and query their own account information