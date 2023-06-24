package backup

import (
	"arso-service/internal/gcs"
	"fmt"
	"io"
	"net/http"
)

type BackupService struct {
	gcsService gcs.GCSService
}

func NewBackupService() BackupService {
	return BackupService{gcsService: gcs.NewGCSService("arso-weather-chart")}
}

const arsoWeatherChartUrl = "https://meteo.arso.gov.si/uploads/probase/www/observ/radar/si0-rm-anim.gif"

func (service *BackupService) BackupARSOAnimation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Updating chart")
	response, _ := http.Get(arsoWeatherChartUrl)
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Printf("We have received non-200 status code: %d\n", response.StatusCode)
		fmt.Printf("Headers: %s", response.Header)
		body, _ := io.ReadAll(response.Body)
		fmt.Printf("Body: %s", string(body))
		panic(fmt.Errorf("non-200 status code from ARSO server"))
	}
	file, _ := io.ReadAll(response.Body)
	service.gcsService.WriteFileToBucket(file, "chart.gif")
	fmt.Println("Chart was successfully updated")
}
