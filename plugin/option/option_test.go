package option

import (
	"testing"

	"github.com/caddyserver/caddy"
	"github.com/insomniacslk/dhcp/dhcpv4"
	"github.com/stretchr/testify/assert"
)

func TestCustomOption(t *testing.T) {
	cases := []struct {
		Name    string
		Value   []string
		Code    uint8
		Payload []byte
		Err     bool
	}{
		{
			"0xaa",
			[]string{"0xaabbccdd", "0xeeff"},
			0xaa,
			[]byte{0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
			false,
		},
		{
			"0b10001000",
			[]string{"fe"},
			0x88,
			[]byte{0xfe},
			false,
		},
		{
			"fo",
			nil,
			0,
			nil,
			true,
		},
		{
			"0xaa",
			[]string{"fae"},
			0xaa,
			nil,
			true,
		},
	}

	for _, c := range cases {
		o, v, err := parseCustomOption(c.Name, c.Value)

		if err == nil {
			assert.Equal(t, c.Code, o.Code())
			assert.Equal(t, c.Payload, v.ToBytes())
			assert.False(t, c.Err, "expected an error")
		} else {
			assert.True(t, c.Err)
		}
	}
}

func getController(input string) *caddy.Controller {
	return caddy.NewTestController("dhcpv4", input)
}

func TestSetupPlugin(t *testing.T) {
	cases := []struct {
		I string
		O map[dhcpv4.OptionCode]dhcpv4.OptionValue
		E bool
	}{}

	for _, c := range cases {
		ctrl := getController(c.I)

		err := setupOption(ctrl)
		if c.E {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}

	}
}
