
# Cheetsheet

## initiate go module
`initialize a Go module`

## Library/ Dependancy config

### Config `CompileDaemon`

Auto-restart Go program on changes by ComplieDaemon

`$ CompileDaemon -command="./<project_name>"`

### Config `Gin`

 ```go
 package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welocome to OppOrakle",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
 ```
__Note :__ better to use `port:8000` to avoid conflict with `postgres`


### Config `godotenv`

1. Paste below in `main.go`

```go
package main

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

func init() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
}

```

2. Refact for modularity
    -  `mkdir initializers`
    - `cd initializers`
    - `touch loadEnvVariables.go`
    - add below content


        ```go
        package main

        import (
            "github.com/joho/godotenv"
            "log"
            "os"
        )

        func init() {
        err := godotenv.Load()
        if err != nil {
            log.Fatal("Error loading .env file")
        }
        }

        ```

    - In `main.go`, paste below insde init function

        ```go
        func init() {
            initializers.LoadEnvVaraibles()
        }
        ```

### config `Postgres`
1. create a file `database.go` inside `initializers` 

```go
package initializers

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB () {
    var err error
    dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatalf("Fail to connect to database: %v", err)
        return
    }
    log.PrintIn("Connected to the database successfully")
}
```

- __Note :__ 
    - use dockerhub too
        - [Postgress](https://hub.docker.com/_/postgres)
        - [Pdadmin](https://hub.docker.com/r/elestio/pgadmin) - postgres GUI
    - use `.env` for env variables
        - DB_URL = "xxxx"
        - in main.go dns = OS.GetEnv("DB_URL")
        - [reference](https://pkg.go.dev/github.com/joho/godotenv@v1.4.0#section-readme)

2. Run the function in `main.go` 

    ```go
    initializers.ConnectToDb()
    ```

### Config Automigration

1. `mkdir models`
2. `cd models`
3. `touch postModel.go`
__Note :__ for table and coloumns use `SnakeCase`



    ```go
        package model

        type User struct {
            gorm.Model
            Email string
            }
    ```

#### [Migrate database](https://gorm.io/docs/migration.html)
1. `mkdir migrate`
2. `touch migrate.go`

    ```go
        package main

        func init() {
            initializers.LoadEnvVariables()
            initializers.ConnectToDB()

        }

        func main () {
            initializers.DB.AutoMigrate(&models.Contact{})
        }
    ```

3. run command `migrate/migrate.go`

4. check using __[Table Plus](https://tableplus.com/download)__ software
- copy ElephanSql url
- paste and create a new conncection

### [CRUD](https://gorm.io/docs/models.html)

#### Create Contact

1. `mkdir controllers`
2. `contactControllers.go`


    ```

3. Use postman to test
 
4. `main.go`
- delete test page details

    ```go
        r.POST("/posts", controllers.ContactCreate)
    ```

#### Get Contact List

1. update contactController.go



    ```go

        // ContactList returns a list of contacts
        func ContactList(C *gin.Context) {
            // Fetch the list of contacts from the database
            var contacts []models.Contact
            if err := initializers.DB.Find(&contacts).Error; err != nil {
                C.JSON(500, gin.H{"error": err.Error()})
                return
            }

            // Return the list of contacts
            C.JSON(200, gin.H{
                "contacts": contacts,
            })
        }
        
    ```

  

2. Use postman to test
 
3. `main.go`

    ```go
        ...
        r.GET("/contacts/". controller.ContactList)
    ```

#### Get Contact Details

1. update contactController.go

    ```go
        func ContactDetails(C *gin.Context) {
            // Get contact 
            contactID := C.Param("id")

            // Query the database to retrieve
            var contact models.Contact
            result := initializers.DB.First(&contact, contactID)

            // Check for errors
            if result.Error != nil {
                C.JSON(404, gin.H{"error": "Contact not found"})
                return
            }

            // Return the retrieved contact
            C.JSON(200, gin.H{
                "contact": contact,
            })
        }

    ```

2. Use postman to test
 
3. `main.go`

    ```go
        ...
        r.GET("/contacts/:id". controller.ContactDetails)
    ```

#### Update

1. update contactController.go

    ```go
        // ContactUpdate updates a contact
        func ContactUpdate(C *gin.Context) {
            // Get contact ID from the URL parameter
            contactID := C.Param("id")

            // Fetch the existing contact from the database
            var existingContact models.Contact
            if err := initializers.DB.First(&existingContact, contactID).Error; err != nil {
                C.JSON(404, gin.H{"error": "Contact not found"})
                return
            }

            // Get data from req body and update the existing contact
            if err := C.ShouldBindJSON(&existingContact); err != nil {
                C.JSON(400, gin.H{"error": err.Error()})
                return
            }

            // Save the updated contact to the database
            if err := initializers.DB.Save(&existingContact).Error; err != nil {
                C.JSON(500, gin.H{"error": err.Error()})
                return
            }

            // Return the updated contact
            C.JSON(200, gin.H{
                "contact": existingContact,
            })
        }

    ```

2. Use postman to test
 
3. `main.go`

    ```go
        ...
        r.PUT("/contacts/:id". controller.ContactUpdate)
    ```

#### Delete Concact

1. update contactController.go

    ```go
        // ContactDelete deletes a contact by ID
        func ContactDelete(C *gin.Context) {
            // Get contact ID from path parameters
            contactID := C.Param("id")

            // Check if the contact ID is provided
            if contactID == "" {
                C.JSON(400, gin.H{"error": "Contact ID is required"})
                return
            }

            // Attempt to delete the contact
            if err := models.DeleteContact(initializers.DB, contactID); err != nil {
                C.JSON(500, gin.H{"error": err.Error()})
                return
            }

            // Return success response
            C.JSON(200, gin.H{"message": "Contact deleted successfully"})
        }
    ```

2. Use postman to test
 
3. `main.go`

    ```go
        ...
        r.DELETE("/contacts/:id". controller.ContactDelete)
    ```