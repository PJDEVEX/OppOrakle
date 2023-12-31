
# Cheetsheet



Initialize go module



## initiate go module

As Go modules enabled (GO111MODULE=on), 

cd /path/to/your/project

`go mod init github.com/yourusername/yourproject`

e.g.: `go mod init github.com/PJDEVEX/OppOrakle`


## Library/ Dependancy config

```bash
go get \
> github.com/gin-gonic/gin \
> github.com/golang-jwt/jwt/v4 \
> golang.org/x/crypto \
> gorm.io/driver/postgres \
> gorm.io/gorm
```

### Config `CompileDaemon`

`go get -u github.com/githubnemo/CompileDaemon`

install - `go install github.com/githubnemo/CompileDaemon`

Auto-restart Go program on changes by ComplieDaemon

`$ CompileDaemon -command="./<project_name>"`

if the project is not in system's `PATH`, use full path,

`/home/pjlinux/go/bin/CompileDaemon -command="./OppOrakle"`

### [Config `Gin`](https://gin-gonic.com/docs/quickstart/)



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
            package initializers

            import (
                "log"

                "github.com/joho/godotenv"
            )

            // Load env variables from .env
            func LoadEnvVarialbles() {
                err := godotenv.Load()
                if err != nil {
                    log.Fatal("Error in loading .env file")
                }
            }


        ```

    - In `main.go`, paste below insde init function

        ```go
                    func init() {
                initializers.LoadEnvVarialbles()
            }

            func main() {
                // Gin router
                r := gin.Default()
                // Route paths
                r.GET("/", func(c *gin.Context) {
                    c.JSON(200, gin.H{
                        "message": "pong",
                    })
                })
                // Set local port
                r.Run(":" + os.Getenv("PORT"))
            }
        ```

### config `Postgres`

Create a PostgreSql database in VM

```bash
    CREATE DATABASE 'databasename';
    CREATE USER 'username' WITH ENCRYPTED PASSWORD 'password';
    ALTER ROLE 'username' SET client_encoding TO 'utf8';
    ALTER ROLE 'username' SET default_transaction_isolation TO 'read committed';
    ALTER ROLE 'username' SET timezone TO 'UTC';
    GRANT ALL PRIVILEGES ON DATABASE 'databasename' TO 'username';
```


1. create a file `database.go` inside `initializers` 

```go
package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connecting to the DB
func ConnectToDB() {
	var err error
	// Retrieve the DB from env veriables.
	dsn := os.Getenv("VMDATABASE_URL")
 	// Trying to open a conncection to the DB.
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}) 
    // it is "=", NOT ":="
    // Use package-level DB variable

	if err != nil {
		// log error
		log.Printf("Fail to connect to database: %v", err)
		return
	}
	// log success
	log.Println("Connected to the database successfully")

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
    <p style="color: red; font-weight: bold;">Remember: run ☝in the root directory</p>
    
    [SuccessfulDatabaseMigration](https://drive.google.com/file/d/1Vy1Vfgt1up0c3XEYSbpoI-7FMGGrVPdb/view?usp=drive_link)


4. check using __[Table Plus](https://tableplus.com/download)__ software
- copy ElephanSql url
- paste and create a new conncection

### [CRUD](https://gorm.io/docs/models.html)

#### Create Application

1. `mkdir controllers`
2. `applicationController.go`

    ```go
        package controllers

        import (
            "time"

            "github.com/PJDEVEX/OppOrakle/initializers"
            "github.com/PJDEVEX/OppOrakle/models"
            "github.com/gin-gonic/gin"
        )

        // POST application
        func AddApplication(c *gin.Context) {
            // Sample for testing
            application := models.ApplicationModel{
                Title:           "Software Engineer",
                Company:         "Example Corp",
                Source:          "LinkedIn",
                Location:        "New York",
                WorkArrangement: "Remote",
                AdvertPdf:       "https://example.com/job_ad.pdf",
                TechStack:       "Go, React, PostgreSQL",
                Softskills:      "Communication, Teamwork",
                Recruiter:       "John Doe",
                HiringManager:   "Jane Smith",
                DatePosted:      time.Now(),
                Deadline:        time.Now().AddDate(0, 0, 14), // Two weeks from now
                AppliedDate:     time.Now().AddDate(0, 0, -7), // One week ago
                ResponseType:    "Pending",
                ResponseDate:    "",
                TimeElapsed:     7 * 24 * time.Hour, // One week in duration
                Note:            "Great potential candidate",
            }

            // Create data in DB
            result := initializers.DB.Create(&application)

            // Error handling
            if result.Error != nil {
                c.Status(400)
                return
            }

            // Confirm the creation
            c.JSON(201, gin.H{
                "application": application,
            })

        }
    ```



3. Use postman to test
 
4. `main.go`
- delete test page details

    ```go
        r.POST("/posts", controllers.AddApplication)
    ```

#### Get Application List

1. update applicationController.go



    ```go

        package main

        import (
            "os"

            "github.com/PJDEVEX/OppOrakle/controllers"
            "github.com/PJDEVEX/OppOrakle/initializers"
            "github.com/gin-gonic/gin"
        )

        // Initialize key comp before the main
        func init() {
            initializers.LoadEnvVarialbles()
            initializers.ConnectToDB()
        }

        func main() {
            // Gin router
            r := gin.Default()
            // Route paths
            r.GET("/applications/", controllers.ApplicationLIst)

            // Set local port
            r.Run(":" + os.Getenv("PORT"))
        }

        
    ```
 

2. Use postman to test
 
3. `main.go`

    ```go
        ...
        r.GET("/contacts/". controller.ContactList)
    ```

#### Refactorting
##### Refactorting - use DYNAMIC data

1. Update applicationController.go

```go
    package controllers

    import (
        "time"

        "github.com/PJDEVEX/OppOrakle/initializers"
        "github.com/PJDEVEX/OppOrakle/models"
        "github.com/gin-gonic/gin"
    )

    // POST application
    func AddApplication(c *gin.Context) {
        // Sample for testing
        var input struct {
            Title           string        `gorm:"column:title;not null" json:"title"`
            Company         string        `gorm:"column:company;not null" json:"company"`
            Source          string        `gorm:"column:source" json:"source"`
            Location        string        `gorm:"column:location" json:"location"`
            WorkArrangement string        `gorm:"column:work_arrangement" json:"work_arrangement"`
            AdvertPdf       string        `gorm:"column:advert_pdf" json:"advert_pdf"`
            TechStack       string        `gorm:"column:tech_stack" json:"tech_stack"`
            Softskills      string        `gorm:"column:softskills" json:"softskills"`
            Recruiter       string        `gorm:"column:recruiter" json:"recruiter"`
            HiringManager   string        `gorm:"column:hiring_manager" json:"hiring_manager"`
            DatePosted      time.Time     `gorm:"column:date_posted;not null" json:"date_posted"`
            Deadline        time.Time     `gorm:"column:deadline" json:"deadline"`
            AppliedDate     time.Time     `gorm:"column:applied_date" json:"applied_date"`
            ResponseType    string        `gorm:"column:response_type" json:"response_type"`
            ResponseDate    string        `gorm:"column:response_date" json:"response_date"`
            TimeElapsed     time.Duration `gorm:"column:time_elapsed" json:"time_elapsed"`
            Note            string        `gorm:"column:note" json:"note"`
        }

        c.Bind(&input)

        application := models.ApplicationModel{
            Title:           input.Title,
            Company:         input.Company,
            Source:          input.Source,
            Location:        input.Location,
            WorkArrangement: input.WorkArrangement,
            AdvertPdf:       input.AdvertPdf,
            TechStack:       input.TechStack,
            Softskills:      input.Softskills,
            Recruiter:       input.Recruiter,
            HiringManager:   input.HiringManager,
            DatePosted:      input.DatePosted,
            Deadline:        input.Deadline,
            AppliedDate:     input.AppliedDate,
            ResponseType:    input.ResponseType,
            ResponseDate:    input.ResponseDate,
            TimeElapsed:     input.TimeElapsed,
            Note:            input.Note,
        }

        // Create data in DB
        result := initializers.DB.Create(&application)

        // Error handling
        if result.Error != nil {
            c.Status(400)
            return
        }

        // Confirm the creation
        c.JSON(201, gin.H{
            "application": application,
        })

    }
```

⛔⛔⛔⛔⛔⛔⛔⛔⛔
<p style="color: red; font-weight: bold ">Refactoring for modularity, scalability, maintenance and separation of concerns</p>

#### Refactorting - Separation of concerns


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