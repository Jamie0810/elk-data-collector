package handler

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLambdaHandler(t *testing.T) {
	ctx := context.Background()

	awsLogs := AwsLogs{Data: "H4sIAAAAAAAAAO1S22rbQBD9FbMU8mLFe7/ozVAnFNIGbD81DmKlHScCXdyV5MQN+fesV016DYS+Ffq0MHPmnJmz5wHV0HX2BtaHHaAUvZ+v59nHxWo1P1+gKWrvGvChbAiVTHFGFTWhXLU3574ddqEzs3fdrLJ17uwMCucTB/ukaOu6be6gqpKd7Uto+qQD64vbZESOFKveg60DB8WUzLCcUTW7encxXy9W62tOMCswl0aA5kay3PB8mwNRlG+dtSpQdEPeFb7c9WXbnJVVD75D6RVytrdhg6qCom89uo5ai31Y4th+QKULkkxSxrGgRkpBCaZaKU0Il0JLSqXSRnIiNNVEUGyMElRrI7g+Ht+XwbLe1uF6IilXhhCqDebTZysD/cMGVbCHaoPSDfrw6exyg6ah1Bb2uGys3trGVeBTSmPv22xsXUSTJnAPxXDET4Kc78FF4It8hB6tS7BMqFoTnAqaYn0advkcoR34fVmMnH/8h4gKVrksKgRg7wcIte3QFEflrLH1OP/mv42cL/M11K0/ZF35daQJVv0MsH60I7xpiFI6kqRDl4Dt+oSkP0YvfR5L/3IdD1+G4F5WuigqqXYsdyZRuVAJ56RIdM5dghmXIDAmW8CRIB+6sgk/lLm2tuW48XflCLn39pD13hbwzE4SiZ0WOlArpayWOJygHNgQG0UKVjC5QY+bBj1Of02lUpxrZpihghkuJTOSYhbyKQkhWhlKNGGCC0ONYQy/ksow/3oqF8vl5fL1WDL8WyyXwbzSg5ucbG1dVocYjpNJ2U2atp9s26Fxp2+KqDkNi/2P6D8V0evHJ9SBw5ErBgAA"}
	event := Event{
		AwsLogs: awsLogs,
	}

	resp, err := LambdaHandler(ctx, event)
	if err != nil {
		t.Fatalf("Job failed: %v", err)
	}
	assert.Equal(t, "Job completed", resp.Message)
}
