package tzf_test

import (
	"fmt"
	"testing"

	"github.com/ringsaturn/tzf"
	tzfrel "github.com/ringsaturn/tzf-rel"
	"github.com/ringsaturn/tzf/pb"
	"google.golang.org/protobuf/proto"
)

var f *tzf.Finder

func init() {
	input := &pb.Timezones{}
	if err := proto.Unmarshal(tzfrel.LiteData, input); err != nil {
		panic(err)
	}
	finder, _ := tzf.NewFinderFromPB(input)
	f = finder
}

func BenchmarkGetTimezoneName(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		_ = f.GetTimezoneName(116.6386, 40.0786)
	}
}

func ExampleFinder_GetTimezoneName() {
	fmt.Println(f.GetTimezoneName(116.6386, 40.0786))
	// Output: Asia/Shanghai
}

func ExampleFinder_GetTimezoneLoc() {
	fmt.Println(f.GetTimezoneLoc(116.6386, 40.0786))
	// Output: Asia/Shanghai <nil>
}