package main

import (
        "fmt"
        "github.com/gin-gonic/gin"
        "net/http"
        "os"
)

func GetPORT() string {
        PORT := os.Getenv("PORT")
        if PORT == "" {
                PORT = "3001"
        }
        return ":" + PORT
}

func main() {
        PORT := GetPORT()
        router := gin.Default()
        router.StaticFS("../", http.Dir("public"))

        fmt.Printf("App running on port %s!", PORT)
        router.Run(PORT)
}
