package controller

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type Job struct{}

func (j *Job) UploadImagesToWebServer(lotteryType string) {
	log.Infof("Executing upload_images_to_web_server with lottery type: %s\n", lotteryType)
}

func (j *Job) AnotherTask() {
	m := map[string]any{
		"name":    "backy",
		"species": "dog",
	}
	mJson, _ := json.Marshal(m)
	log.Info(string(mJson))
	log.Infof("Executing another task without parameters.")
}
