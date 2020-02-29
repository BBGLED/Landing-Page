package main

import (
        "fmt"
        "github.com/joho/godotenv"
        "log"
        "net/http"
        "os"
)

func LoadEnv() {
        err := godotenv.Load()
        if err != nil {
                log.Fatal("Error loading .env file")
        }
}

func main() {
        fmt.Println("RUNNING MAIN, WOO")
        LoadEnv()
        PORT := os.Getenv("PORT")
        fmt.Printf("TESTING PORT ON HEROKU: %s!", PORT)
        http.Handle("/", http.FileServer(http.Dir("./public")))

        fmt.Printf("App running on port %s!", PORT)
        http.ListenAndServe(PORT, nil)
}
