package ga

import (
	"os"

	"github.com/denisbrodbeck/machineid"
	"github.com/google/uuid"
)

var gaClient *GAClient

func init() {
	trackingID := os.Getenv("GA_TRACKING_ID") // Tracking ID for Google Analytics
	clientID, err := machineid.ProtectedID("hrp")
	if err != nil {
		nodeUUID, _ := uuid.NewUUID()
		clientID = nodeUUID.String()
	}
	gaClient = NewGAClient(trackingID, clientID)
}

func SendEvent(e IEvent) error {
	return gaClient.SendEvent(e)
}