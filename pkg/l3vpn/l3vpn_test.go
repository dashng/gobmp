package l3vpn

import (
	"reflect"
	"testing"

	"github.com/sbezverk/gobmp/pkg/base"
)

func TestUnmarshalL3VPNNLRI(t *testing.T) {
	tests := []struct {
		name   string
		input  []byte
		expect *NLRI
		fail   bool
	}{
		{
			name:  "nlri 1",
			input: []byte{120, 5, 220, 49, 0, 0, 2, 65, 0, 0, 253, 235, 3, 3, 3, 3},
			expect: &NLRI{
				Length: 15,
				Labels: []*base.Label{
					{
						Value: 24003,
						Exp:   0,
						BoS:   true,
					},
				},
				RD: &base.RD{
					Type:  0,
					Value: []byte{2, 65, 0, 0, 253, 235},
				},
				Prefix: []byte{3, 3, 3, 3},
			},
			fail: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalL3VPNNLRI(tt.input)
			if err != nil && !tt.fail {
				t.Fatalf("expected to succeed but failed with error: %+v", err)
			}
			if err == nil && tt.fail {
				t.Fatalf("expected to fail but succeeded")
			}
			if err == nil {
				if !reflect.DeepEqual(got, tt.expect) {
					t.Errorf("Expected label %+v does not match to actual label %+v", got, *tt.expect)
				}
			}
		})
	}
}