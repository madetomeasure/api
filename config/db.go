import (
    "encoding/json"
    "os"
    "fmt"
    "github.com/jinzhu/gorm"
    "github.com/lib/pq"
)

func connectionInfo() string {
    file, _ := os.Open("database.json")
    decoder := json.NewDecoder(file)

    var config DBConfiguration

    err := decoder.Decode(&config)

    if err != nil {
      fmt.Println("Error:", err)
    }


    fmt.Println(configuration)
    info := fmt.Sprintf(
      "user=%s password=%s host=%s encoding=%s dbname=%s",
      config.Username, config.Password, config.Host,
      config.Encoding, config.Database
    )
    return info
}

db, err := gorm.Open("postgres", connectionInfo())

// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
db.DB()

// Then you could invoke `*sql.DB`'s functions with it
db.DB().Ping()
db.DB().SetMaxIdleConns(10)
db.DB().SetMaxOpenConns(100)
