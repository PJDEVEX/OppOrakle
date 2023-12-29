# OppOrakle
Embark on a transformative journey with CareerTracker, an empowering application designed to elevate your job application experience. 🌟  

Unleash the potential of your career aspirations by effortlessly tracking your job applications and seamlessly managing the status of each opportunity. 🌐

## 5 Planes of UX
Creating User Centric Design
### :dart: Strategy - Goals and Objective

### 🚧 Scope - Boundries
### :classical_building: Structure - Organizing content and features

#### CRUD tables

##### User CRUD table

##### Jobs CRUD table
|HTTP|URI|CRUD operation|View name|
|:----|:----|:----|:----|
|GET|/jobs|list all jobs|LIST|
|POST|/jobs|crate a job|LIST|
|GET|/jobs/{id}|retrive a job by id|Detail|
|PUT|/jobs/{id}|update a job by id|Detail|
|DELETE|/jobs/{id}|Delete a job by id|Detail|

##### Contacts CRUD table
|HTTP|URI|CRUD operation|View name|
|:----|:----|:----|:----|
|GET|/contacts|list all contacts|LIST|
|POST|/contacts|crate a contact|LIST|
|GET|/contacts/{id}|retrieve a contact by id|Detail|
|PUT|/contacts/{id}|update a contact by id|Detail|
|DELETE|/contacts/{id}|Delete a contact by id|Detail|

##### Company CRUD table
|HTTP|URI|CRUD operation|View name|
|:----|:----|:----|:----|
|GET|/companies|list all companies|LIST|
|POST|/companies|crate a company|LIST|
|GET|/companies/{id}|retrieve a company by id|Detail|
|PUT|/companies/{id}|update a company by id|Detail|
|DELETE|/companies/{id}|Delete a company by id|Detail|




### 🗺️ Skeleton - navigation and flow
### 💻 Surface - look and feel

### Technology/ Tools  🔧 
#### Backend

| # | Dependency                                         | Usage                                                |
|---|----------------------------------------------------------|------------------------------------------------------|
| 1 | [![gin-gonic/gin](https://img.shields.io/badge/gin--gonic%2Fgin-v1.7.4-green)](https://gin-gonic.com/docs/quickstart/) | Web framework for building Go applications with ease. |
| 2 | [![golang-jwt/jwt](https://img.shields.io/badge/golang--jwt%2Fjwt--v5-v3.7.3-green)](https://github.com/golang-jwt/jwt) | JSON Web Token (JWT) implementation in Go.            |
| 3 | [![joho/godotenv](https://img.shields.io/badge/joho%2Fgodotenv-v1.3.0-green)](https://github.com/joho/godotenv) | Load environment variables from a .env file.           |
| 4 | [![golang.org/x/crypto](https://img.shields.io/badge/golang.org%2Fx%2Fcrypto-v0.0.0-9cf)](https://pkg.go.dev/golang.org/x/crypto#section-readme) | Cryptographic primitives for Go.                      |
| 5 | [![go-gorm/postgres](https://img.shields.io/badge/go--gorm%2Fpostgres-v1.11.0-green)](https://github.com/go-gorm/postgres) | PostgreSQL driver for GORM, an ORM for Go.           |
| 6 | [![gorm.io/gorm](https://img.shields.io/badge/gorm.io%2Fgorm-v1.23.8-green)](https://gorm.io/) | Feature-rich ORM for Go.                              |
| 7 | [![githubnemo/compiledaemon](https://img.shields.io/badge/githubnemo%2Fcompiledaemon-v1.2.1-blue)](https://pkg.go.dev/github.com/githubnemo/compiledaemon/v2) | Watches .go files and invokes go build if a file changed. |


## API Endpoints

### Users API Endpoints
|HTTP method|Endpoint|Request Parameters|Description|
|:----|:----|:----|:----|
|GET|/api/users|-|-|
|GET|/api/users/{id}|-|Get details of a specific user|
|POST|/api/users|JSON payload: {}|Create a new user|
|PUT|/api/users/{id}|JSON payload: {}|Update details of a specific user|
|DELETE|/api/users/{id}|-|-|



### Jobs API Endpoints
|HTTP method|Endpoint|Request Parameters|Description|
|:----|:----|:----|:----|
|GET|/api/jobs|-|Get a list of all jobs|
|GET|/api/jobs/{id}|-|Get details of a specific job|
|POST|/api/jobs|JSON payload: {}|Create a new job|
|PUT|/api/jobs/{id}|JSON payload: {}|Update details of a specific job|
|DELETE|/api/jobs/{id}|-|Delete a specific job|

### Contacts API Endpoints
|HTTP method|Endpoint|Request Parameters|Description|
|:----|:----|:----|:----|
|GET|/api/contacts|-|Get a list of all contacts|
|GET|/api/contacts/{id}|-|Get details of a specific contact|
|POST|/api/contacts|JSON payload: {}|Create a new contact|
|PUT|/api/contacts/{id}|JSON payload: {}|Update details of a specific contact|
|DELETE|/api/contacts/{id}|-|Delete a specific contact|

### Company API Endpoints
|HTTP method|Endpoint|Request Parameters|Description|
|:----|:----|:----|:----|
|GET|/api/companies|-|Get a list of all companies|
|GET|/api/companies/{id}|-|Get details of a specific company|
|POST|/api/companies|JSON payload: {}|Create a new company|
|PUT|/api/companies/{id}|JSON payload: {}|Update details of a specific company|
|DELETE|/api/companies/{id}|-|Delete a specific company|

## Reference
1. [Understanding the Five Planes of UX](https://bootcamp.uxdesign.cc/demystifying-the-ux-process-24e5b10e77b8)