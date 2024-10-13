package teapods

import (
	"testing"

	"github.com/enuesaa/teatime/pkg/router/apptest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	app := apptest.New(t)
	res, err := app.Post(Create,
		apptest.UseBody(`{"name": "teapod-links"}`),
	)
	require.NoError(t, err)

	assert.Equal(t, 200, res.Code)
	assert.Equal(t, true, res.GetB("data.ok"))
}
