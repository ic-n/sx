package testing

import (
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

	var newPoll commitrevealv1.CreatePollResponse
	{
		resp, err := c.PostForm(base+"/v1/poll", url.Values{
			"choice_1": []string{"Yes"},
			"choice_2": []string{"No"},
			"seconds":  []string{"2"},
		})
		require.NoError(t, err)

		b, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		resp.Body.Close()

		err = protojson.Unmarshal(b, &newPoll)
		require.NoError(t, err, string(b))
	}
	require.NotEmpty(t, newPoll.Address)

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
	require.Greater(t, poll.SecondsLeft, int64(0))

	{
		_, err := c.PostForm(base+"/v1/commit", url.Values{
			"address": []string{newPoll.GetAddress()},
			"secret":  []string{"1~test"},
		})
		require.NoError(t, err)
	}
	{
		_, err := c.PostForm(base+"/v1/commit", url.Values{
			"address": []string{newPoll.GetAddress()},
			"secret":  []string{"0~test"},
		})
		require.NoError(t, err)
	}

	time.Sleep(2 * time.Second)
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
}
