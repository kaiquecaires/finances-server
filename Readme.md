# Finances App

Introducing my financial application that puts you in control of your finances! With this app, you can effortlessly track all your expenses and set up personalized recurring bills. Never miss a payment again with handy bill payment reminders at your fingertips. Take charge of your financial future by simulating investments and receiving comprehensive reports to help you plan ahead.

# Getting started

1. Create a .env file
2. Copy the content from .env.example
3. Run the migrations (migrations up)
3. Run the command:
```bash
  docker-compose up -d
```

# Migrations
To run the migrations you'll need the go migrate package:
https://github.com/golang-migrate/migrate#cli-usage

## Create
```bash
  migrate create -ext sql -dir ./migrations -seq your_migration_name
```

## Run migrations down
```bash
migrate -path ./migrations/ -database "postgresql://admin:admin@localhost:5432/finances?sslmode=disable" -verbose down
```

## Run migrations up
```bash
migrate -path ./migrations/ -database "postgresql://admin:admin@localhost:5432/finances?sslmode=disable" -verbose up
```

## Features
- [x] Signup
- [x] Sign in
- [x] Register income
- [x] Register outcome
- [x] List transactions
- [x] Get total amount
- [x] Delete transactions
- [ ] Register recurring bills

## Entities
### User
- Unique identifier
- Name
- Social name
- Birthday
- Email

### User Passwords
- Unique Identifier
- Account Unique identifier
- Password

### Bill Categories
- Unique identifier
- Description

### Transactions
- Unique identifier
- Unique user identifier
- Unique bill category
- Type (income or outcome)
- Value
- Date
