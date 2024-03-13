package db

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
    dsn := "host=localhost user=postgres password=230798 dbname=orders_service port=5432 sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}
