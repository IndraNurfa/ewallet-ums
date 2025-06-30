package services

import (
	"ewallet-ums/internal/interfaces"
	"fmt"
	"time"
)

type Healthcheck struct {
	HealthcheckRepository interfaces.IHealthcheckRepo
}

func (s *Healthcheck) HealthcheckServices() (string, error) {
	time := time.Now().UTC().Local().Format("2006-01-02 15:04:05")
	return fmt.Sprintf("service healthy, time: %s", time), nil
}
