package azure

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	databricks "github.com/xinsnake/databricks-sdk-golang"
)

func Test_MultipleCallsToAPI(t *testing.T) {

	// Setup before running:
	// - run mock api
	// - run `watch -n  1 /tmp/netstat.sh` to monitor number of open connections (Will see 10 extra in ESTABLISHED state while paused which then switch to TIME_WAIT)
	// Goal: update the SDK to reuse connections so that the above conditions no longer occur :-)
	// (go clean -testcache && go test -v ./azure -run Test_MultipleCallsToAPI)
	// ( watch -n 1 /tmp/netstat.sh  )

	// /tmp/netstat.sh:
	/*
		echo "ESTABLISHED: $(netstat -t | grep localhost | grep EST | wc -l)"
		echo "TIME_WAIT: $(netstat -t | grep localhost | grep TIME_WAIT | wc -l)"
		echo "TOTAL: $(netstat -t | grep localhost | wc -l)"
		echo "--------------"
		#netstat -t | grep localhost
	*/

	//Arrange
	var o databricks.DBClientOption
	o.Host = "http://localhost:8080"
	o.Token = "dummy - not checked by mock api"

	var c DBClient
	c.Init(o)

	//Act
	for index := 0; index < 10; index++ {
		_, err := c.Jobs().List()
		// TODO - check the error ;-)
		assert.Nil(t, err)
	}
	_ = "foo1"

	time.Sleep(30 * time.Second) // sleep to keep the process alive so that we can check the number of open sockets

	// Assert
	// assert.False(false) // TODO - decide whether we want to validate the number of sockets here or manually ;-)
}
