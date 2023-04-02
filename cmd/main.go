package main

import (
	_ "github.com/joho/godotenv/autoload"
	"hte-danger-zone-job/internal/job"
)

func main() {
	j := job.New()
	j.Run()
}
