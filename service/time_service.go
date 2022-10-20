package service

import (
	"fmt"
	"time"
)

type TimeService struct {
}

type TimeServiceRequest struct {
}

type TimeServiceResponse struct {
	Status      string
	CurrentTime time.Time
}

func (t *TimeService) GetTime(request TimeServiceRequest, response *TimeServiceResponse) error {
	// Directly reply the current time
	response.CurrentTime = time.Now()
	response.Status = "success"
	fmt.Println(response.CurrentTime)
	return nil
}
