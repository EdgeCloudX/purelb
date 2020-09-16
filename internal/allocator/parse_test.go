package allocator

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"

	"purelb.io/pkg/apis/v1"
)

func TestParse(t *testing.T) {
	tests := []struct {
		desc string
		raw  []*v1.ServiceGroup
		want map[string]Pool
	}{
		{desc: "empty config",
			raw:  []*v1.ServiceGroup{},
			want: map[string]Pool{},
		},

		{desc: "config using all features",
			raw: []*v1.ServiceGroup{
				localServiceGroup("pool1", "10.20.0.0/16"),
				localServiceGroup("pool2", "30.0.0.0/8"),
				localServiceGroup("pool3", "40.0.0.0/25"),
				localServiceGroup("pool4", "2001:db8::/126"),
				egwServiceGroup("pool5", "url"),
			},
			want: map[string]Pool{
				"pool1": mustLocalPool(t, "10.20.0.0/16"),
				"pool2": mustLocalPool(t, "30.0.0.0/8"),
				"pool3": mustLocalPool(t, "40.0.0.0/25"),
				"pool4": mustLocalPool(t, "2001:db8::/126"),
				"pool5": mustEGWPool(t, "url"),
			},
		},

		{desc: "invalid CIDR",
			raw: []*v1.ServiceGroup{
				localServiceGroup("pool1", "100.200.300.400/24"),
			},
		},

		{desc: "invalid CIDR prefix length",
			raw: []*v1.ServiceGroup{
				localServiceGroup("pool1", "1.2.3.0/33"),
			},
		},

		{desc: "duplicate group name",
			raw: []*v1.ServiceGroup{
				localServiceGroup("pool1", "10.20.0.0/16"),
				localServiceGroup("pool1", "30.0.0.0/8"),
			},
		},

		{desc: "duplicate CIDRs",
			raw: []*v1.ServiceGroup{
				localServiceGroup("pool1", "10.0.0.0/8"),
				localServiceGroup("pool2", "10.0.0.0/8"),
			},
		},

		{desc: "overlapping CIDRs",
			raw: []*v1.ServiceGroup{
				localServiceGroup("pool1", "10.0.0.0/8"),
				localServiceGroup("pool2", "10.0.0.0/16"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got, err := parseConfig(test.raw)
			if err != nil && test.want != nil {
				t.Errorf("%q: parse failed: %s", test.desc, err)
				return
			}
			if test.want == nil && err == nil {
				t.Errorf("%q: parse succeeded but should have failed", test.desc)
				return
			}
			egwComparer := cmp.Comparer(func(x, y EGWPool) bool {
				return reflect.DeepEqual(x.url, y.url)
			})
			iprangeComparer := cmp.Comparer(func(x, y IPRange) bool {
				return reflect.DeepEqual(x.from, y.from) && reflect.DeepEqual(x.to, y.to)
			})
			if diff := cmp.Diff(test.want, got, iprangeComparer, egwComparer, cmp.AllowUnexported(LocalPool{})); diff != "" {
				t.Errorf("%q: parse returned wrong result (-want, +got)\n%s", test.desc, diff)
			}
		})
	}
}
