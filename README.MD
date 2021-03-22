# Bolt API.
 Go API to create tasks and projects.

 ## Table of Contents
- [Prerequisites](#Prerequisites)
- [Configure your local environment](#configure-your-local-environment)
- [Using the local environment](#using-the-local-environment)


 ## Prerequisites
- [Docker Desktop](https://hub.docker.com/?overlay=onboarding) 
- [Golang 1.15.8](https://golang.org/)

## Configure your local environment
(only for the fist time)
**0. First Steps:**
```
# Download Repository in your golang path
git clone https://github.com/getchipman/bolt-api.git
# Access the folder
cd bolt-api
# Install Golang Dependencies
go install
```
**1. Create containers:**
```
docker-compose up
```
**1.1 Configs for your database GUI**
| Property       | Value                 |
| -------------- | --------------------- |
| Host           | localhost             |
| Port           | 9000                  |
| Database       | postgres              |
| Username       | postgres              |
| Password       | abcd1234              |
|                |                       |
|                |                       |

**2. Execute the application in 'web mode' and wait for 1 minute - this will create database structure.**
```shell
go run main.go web
```

**Server Ports**
| Application    | Address                 |
| -------------- | ----------------------- |
| Internal Web   | http://localhost:9000   |
|                |                         |
---

## Using the local environment
1. Up containers:
```
docker-compose up
```

2. Run Web Service.
```
go run main.go web
```

3. Stop containers:
```
docker-compose stop
```

4. Clear containers:
```
docker-compose down
```
---