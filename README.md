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
| 3 | [![joho/godotenv](https://img.shields.io/badge/joho%2Fgodotenv-v1.5.1-green)](https://github.com/joho/godotenv) | Load environment variables from a .env file.           |
| 4 | [![golang.org/x/crypto](https://img.shields.io/badge/golang.org%2Fx%2Fcrypto-v0.0.0-9cf)](https://pkg.go.dev/golang.org/x/crypto#section-readme) | Cryptographic primitives for Go.                      |
| 5 | [![go-gorm/postgres](https://img.shields.io/badge/go--gorm%2Fpostgres-v1.11.0-green)](https://github.com/go-gorm/postgres) | PostgreSQL driver for GORM, an ORM for Go.           |
| 6 | [![gorm.io/gorm](https://img.shields.io/badge/gorm.io%2Fgorm-v1.23.8-green)](https://gorm.io/) | Feature-rich ORM for Go.                              |
| 7 | [![githubnemo/compiledaemon](https://img.shields.io/badge/githubnemo%2Fcompiledaemon-v1.2.1-blue)](https://pkg.go.dev/github.com/githubnemo/compiledaemon/v2) | Watches .go files and invokes go build if a file changed. |
| 8 | [![aws-sdk-go](https://img.shields.io/badge/aws--sdk--go-v1.49.13-brightorange?color=%23FF9900)](https://github.com/aws/aws-sdk-go/releases/tag/v1.49.13) | Official AWS SDK for the Go programming language.  |

**Change Log:**
- Include your log entries here.
 | | |




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

## Testing

### Manual Testing
#### Backend

#### Initial Configarations working fine

<details>
  <summary>Initial setup works</summary>
  <img src="https://res.cloudinary.com/pjdevex/image/upload/v1703972542/oppOrakle/devServerUp_brdjme.png" alt="Initial setup works">
</details>

#### Automigration of database works fine

<details>
  <summary>Automigration of database</summary>
  <img src="https://res.cloudinary.com/pjdevex/image/upload/v1703973068/oppOrakle/automigrationsWorks_b0krl4.png" alt="Automigration of database works fine">
</details>

#### CRUD operation

##### Application - AddApplication - static data
<details>
  <summary>Application - postApplicationPostman</summary>
  <img src="https://res.cloudinary.com/pjdevex/image/upload/v1703985108/oppOrakle/postApplicationPostman_j46bv1.jpg" alt="Application - postApplicationPostman">
</details>

<details>
  <summary>Application - postApplicationView-localhost</summary>
  <img src="https://res.cloudinary.com/pjdevex/image/upload/v1703985108/oppOrakle/postApplicationView_jjaidd.jpg" alt="Application - postApplicationView-localhost">
</details>

<details>
  <summary>Application - postApplicationDataPlus</summary>
  <img src="https://res.cloudinary.com/pjdevex/image/upload/v1703985661/oppOrakle/postApplicationDataPlusView_hkuwsy.jpg" alt="Application - postApplicationDataPlusView">
</details>

##### Application - AddApplication - dynamic data
<details>
  <summary>AddAppDynamicData-postman</summary>
  <img src="https://res.cloudinary.com/pjdevex/image/upload/v1703989729/oppOrakle/AddAppDynamicData-postman_zjh1o6.jpg" alt="AddAppDynamicData-postman">
</details>

<details>
  <summary>AddAppDynamicData-dataPlus</summary>
  <img src="https://res.cloudinary.com/pjdevex/image/upload/v1703989729/oppOrakle/AddAppDynamicData-dataPlus_phaza7.jpg" alt="AddAppDynamicData-dataPlus">
</details>

<details>
  <summary>AddAppDynamicData-localhost</summary>
  <img src="https://res.cloudinary.com/pjdevex/image/upload/v1703989729/oppOrakle/AddAppDynamicData-localhost_m9cs8v.jpg" alt="AddAppDynamicData-localhost">
</details>

#### Refactorting - Separation of concerns
<details>
  <summary>Refactorting-SeparationOfConcerns-postman</summary>
  <img src="https://res.cloudinary.com/pjdevex/image/upload/v1704001568/oppOrakle/Refactorting-SeparationOfConcerns-postman_kcodpm.jpg" alt="Refactoring - Separation of Concerns in Postman">
</details>

<details>
  <summary>Refactorting-SeparationOfConcerns-TablePlus</summary>
  <img src="https://res.cloudinary.com/pjdevex/image/upload/v1704001317/oppOrakle/Refactorting-SeparationOfConcerns-TablePlus_bix2bw.jpg" alt="Refactoring - Separation of Concerns in TablePlus">
</details>

<details>
  <summary>Refactorting-SeparationOfConcerns-localhost</summary>
  <img src="https://res.cloudinary.com/pjdevex/image/upload/v1704000889/oppOrakle/Refactorting-SeparationOfConcerns-localhost_onzew4.png" alt="Refactoring - Separation of Concerns on localhost">
</details>


##### Application - ListApplication

<details>
  <summary>Application - ApplicationList-localhost</summary>
  <img src="https://res.cloudinary.com/pjdevex/image/upload/v1703980325/oppOrakle/getApplicationList_eiq9cy.jpg" alt="Application - ApplicationList-localhost">
</details>
<details>
  <summary>Application - ApplicationList-postman</summary>
  <img src="https://res.cloudinary.com/pjdevex/image/upload/v1703980325/oppOrakle/getApplicationListPostman_sihc5b.jpg" alt="Application - ApplicationList-postman">
</details>

## Reference
1. [Understanding the Five Planes of UX](https://bootcamp.uxdesign.cc/demystifying-the-ux-process-24e5b10e77b8)
2. [aws s3](https://github.com/brightfame/ssh-iam)
3. [File Uploads in Go/Gin (AWS S3 uploads and save to disk)](https://www.youtube.com/watch?v=ssLcGwHv7Hc)

