package app

const HUB = "https://hub.iot-noob.com"


func Download(id string) error {
	url := HUB + "/app/"+id+"/download"
	

	return nil
}
