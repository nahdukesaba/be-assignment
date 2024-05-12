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


### Limitation

This services only run the db on docker. I still cannot join both the app and db dockers into 1 container.
Somehow the db always refuse the connection from the app. Still searching for the workaround for this one.

If you feel very adventurous, you can uncomment the app service on `docker-compose.yml` 
and apply your pre-discovered solution.

if you feel not so adventurous but still want to run both app and db in the same container, 
you can uncomment _again_ the section 
`#    network_mode: host`. This will allow you to do so but faced with another limitation 
which is you cannot access the api from your localhost. This time you are refused by the app. 
This is to be expected since `Published ports are discarded when using host network`

##### Note
My Github account got tangled on my IDE so i need to authorized my other Github account so i can commit my changes