package analytics

import (
	"fmt"
	"time"
	log "github.com/sirupsen/logrus"
)

// dev_type:
// 	0: UE
// 	1: GNB
func LogUeTime(dev_type uint8, id string, task string) {

	now := time.Now()

	go ShowUeLog(dev_type, id, task, now)
}

func ShowUeLog(dev_type uint8, id string, task string, now time.Time) {
	if id > 0 {
		nsec_now := now.UnixNano()
		log.Info(fmt.Sprintf("[Analytics] %d, %s, %s, %d", dev_type, id, task, nsec_now))
	}
}