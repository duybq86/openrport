package simple_client_connects_test

import (
	"context"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/cloudradar-monitoring/rport/bdd/helpers"
)

func TestClientConnects(t *testing.T) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*5)
	defer cancel()

	rd, rc := helpers.StartClientAndServerAndWaitForConnection(ctx, t)

	defer func() {
		helpers.LogAndIgnore(rd.Process.Kill())
		helpers.LogAndIgnore(rc.Process.Kill())
	}()

	assertProcessiesAreNotDead(t, rd, rc)
}

func assertProcessiesAreNotDead(t *testing.T, rd *exec.Cmd, rc *exec.Cmd) {
	assert.Nil(t, rd.ProcessState)
	assert.Nil(t, rc.ProcessState)
}
