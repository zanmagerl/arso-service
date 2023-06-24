package main

import (
	"arso-service/internal/backup"
	"net/http"
)

func main() {
	backupService := backup.NewBackupService()
	http.HandleFunc("/rest/charts/updateChart", backupService.BackupARSOAnimation)
	http.ListenAndServe(":8080", nil)
}
