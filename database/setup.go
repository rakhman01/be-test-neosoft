package database

import (
    "fmt"
    "log"
    "net/url"
    "os"
    "strings"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func Connect() *gorm.DB {
    var dsn string

    // cek ENV DATABASE_URL
    rawURL := os.Getenv("DATABASE_URL")
    if rawURL != "" {
        // pakai Railway / live database
        u, err := url.Parse(rawURL)
        if err != nil {
            log.Fatal("Invalid DATABASE_URL: ", err)
        }
        user := u.User.Username()
        pass, _ := u.User.Password()
        host := u.Hostname()
        port := u.Port()
        dbName := strings.TrimPrefix(u.Path, "/")

        dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
            user, pass, host, port, dbName)
        fmt.Println("Menggunakan DATABASE_URL dari ENV:", rawURL)
    } else {
        // fallback ke database lokal
        dsn = "root:root@tcp(127.0.0.1:8889)/golang-invoice?charset=utf8mb4&parseTime=True&loc=Local"
        fmt.Println("ENV DATABASE_URL tidak ditemukan, menggunakan database lokal")
    }

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Gagal konek database: ", err)
    }

    fmt.Println("Berhasil konek ke database!")
    return db
}
