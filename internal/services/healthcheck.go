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
	time := time.Now()
	return fmt.Sprintf("service healthy, time: %v", time), nil
}
