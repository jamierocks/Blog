package main

import (
    "github.com/jamierocks/blog/models"
    "github.com/jamierocks/blog/modules"
)

func main() {
    // Initialise config
    modules.InitConfig()

    // Initialise database
    modules.InitDatabase()
    models.AutoMigrate()
}
