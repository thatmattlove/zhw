package zhw_test

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
	"go.mdl.wtf/zhw"
)

func Test_Writer(t *testing.T) {
	u, _ := url.Parse("http://localhost")
	w, err := zhw.NewWriter(
		zhw.WithURL(u),
		zhw.WithMethod(http.MethodPost),
	)
	require.NoError(t, err)
	defer w.Close()
	zerolog.TimestampFieldName = "dt"
	log.Logger = log.Output(w)
	_log := log.With().Str("str", "string").Int("int", 15).Uint("uint", 15).Strs("slice", []string{"one", "two"}).Logger()
	_log.Info().Msg("message 1")
	_log.Info().Msg("message 2")
	_log.Err(fmt.Errorf("test error")).Msg("an error")
}
