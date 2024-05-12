# BE Assignment

`Be Assignment` is Backend services which manages user’s accounts and transactions

## Prerequisite

To run this program, you will need

### App Dependencies

```$xslt
- Golang 1.18+
- Go mod enabled
- docker
```

## How to Run

- Verify and download dependencies `make build`
- Run the docker to serve database `make docker`
- Run the application `make run`


### API List

In this services you can login and register new user.
Each Registered user will be given 2 accounts(credit/debit) and 1k debit balance as a **Bonus**

Each account can withdraw or send their balance to another account. 
User can also see all their account and check their account transaction history.

You can check the swagger info for API Documentation by visiting
http://localhost:9000/api/swagger/index.html



--------------


# Take home assignment


## Description:
Build 2 Backend services which manages user’s accounts and transactions (send/withdraw). 

In Account Manager service, we have:
- User: Login with Id/Password
- Payment Account: One user can have multiple accounts like credit, debit, loan...
- Payment History: Records of transactions

In Payment Manager service, we have:
- Transaction: Include basic information like amount, timestamp, toAddress, status...
- We have a core transaction process function, that will be executed by `/send` or `/withdraw` API:

```js
function processTransaction(transaction) {
    return new Promise((resolve, reject) => {
        console.log('Transaction processing started for:', transaction);

        // Simulate long running process
        setTimeout(() => {
            // After 30 seconds, we assume the transaction is processed successfully
            console.log('transaction processed for:', transaction);
            resolve(transaction);
        }, 30000); // 30 seconds
    });
}

// Example usage
let transaction = { amount: 100, currency: 'USD' }; // Sample transaction input
processTransaction(transaction)
    .then((processedTransaction) => {
        console.log('transaction processing completed for:', processedTransaction);
    })
    .catch((error) => {
        console.error('transaction processing failed:', error);
    });
```

Features:
- Users need to register/log in and then be able to call APIs.
- APIs for 2 operations send/withdraw. Account statements will be updated after the transaction is successful.
- APIs to retrieve all accounts and transactions per account of the user.
- Write Swagger docs for implemented APIs (Optional)

### Tech-stack:
- Recommend using authentication 3rd party: Supertokens, Supabase...
- `NodeJs/Golang` for API server (`Fastify/Gin` framework is the best choices)
- `PostgreSQL/MongoDB` for Database. Recommend using `Prisma` for ORM.
- `Docker` for containerization. Recommend using `docker-compose` for running containers.
 
## Target:
- Good document/README to describe your implementation.
- Make sure app functionality works as expected. Run and test it well.
- Containerized and run the app using Docker.
- Using `docker-compose` or any automation script to run the app with single command is a plus.
