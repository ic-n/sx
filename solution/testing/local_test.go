package testing

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	commitrevealv1 "github.com/ic-n/sx/pkg/api/gen/commitreveal/v1"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
)

const base = "http://127.0.0.1:80"

func TestAPI(t *testing.T) {
	c := http.Client{}

	resp, err := c.Get(base + "/v1/health")
	if err != nil {
		t.Logf("skipping: %s", err.Error())
		t.Skip()
		return
	}
	if resp.StatusCode != http.StatusOK {
		t.Logf("skipping: HTTP Status %s", resp.Status)
		t.Skip()
		return
	}

	var newPoll commitrevealv1.CreatePollResponse
	{
		resp, err := c.PostForm(base+"/v1/poll", url.Values{
			"choice_1": []string{"Yes"},
			"choice_2": []string{"No"},
			"seconds":  []string{"5"},
		})
		require.NoError(t, err)

		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		resp.Body.Close()

		err = protojson.Unmarshal(b, &newPoll)
		require.NoError(t, err, string(b))
	}
	require.NotEmpty(t, newPoll.Address)

	<-time.After(time.Second)

	var poll commitrevealv1.GetPollResponse
	{
		resp, err := c.Get(base + "/v1/poll?address=" + newPoll.GetAddress())
		require.NoError(t, err)

		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		resp.Body.Close()

		err = protojson.Unmarshal(b, &poll)
		require.NoError(t, err, string(b))
	}
	require.Equal(t, "Yes", poll.Choice_1)
	require.Equal(t, "No", poll.Choice_2)
	require.GreaterOrEqual(t, poll.SecondsLeft, int64(4))

	for i := 0; i < 9; i++ {
		resp, err := c.PostForm(base+"/v1/commit", url.Values{
			"address": []string{newPoll.GetAddress()},
			"secret":  []string{fmt.Sprintf("%d~test%d", i%2+1, i)},
		})
		if resp.StatusCode == http.StatusBadRequest {
			break
		}
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.NoError(t, err)
	}

	<-time.After(time.Second * 4)

	{
		resp, err := c.Get(base + "/v1/poll?address=" + newPoll.GetAddress())
		require.NoError(t, err)

		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		resp.Body.Close()

		err = protojson.Unmarshal(b, &poll)
		require.NoError(t, err, string(b))
	}
	require.Equal(t, "Yes", poll.Choice_1)
	require.Equal(t, "No", poll.Choice_2)
	require.Equal(t, int64(0), poll.SecondsLeft)

	for i := 0; i < 9; i++ {
		resp, err := c.PostForm(base+"/v1/reveal", url.Values{
			"address": []string{newPoll.GetAddress()},
			"secret":  []string{fmt.Sprintf("%d~test%d", i%2+1, i)},
		})
		if resp.StatusCode == http.StatusBadRequest {
			break
		}
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.NoError(t, err)
	}

	{
		resp, err := c.Get(base + "/v1/poll?address=" + newPoll.GetAddress())
		require.NoError(t, err)

		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		resp.Body.Close()

		err = protojson.Unmarshal(b, &poll)
		require.NoError(t, err, string(b))
	}

	require.Equal(t, "Yes", poll.Choice_1)
	require.Equal(t, "No", poll.Choice_2)
	require.EqualValues(t, 0, poll.SecondsLeft)
	require.Greater(t, poll.Count_1, int64(0))
	require.Greater(t, poll.Count_2, int64(0))
}
