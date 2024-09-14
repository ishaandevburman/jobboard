// utils/utils.go
package utils

import (
    "gorm.io/driver/sqlite" // or postgres
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    var err error
    // Use SQLite database file
    DB, err = gorm.Open(sqlite.Open("college_placements.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }
}
