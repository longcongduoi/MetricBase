package graphiteTcp

import (
	"testing"

	"github.com/msiebuhr/MetricBase/metrics"
)

// What about some basic lines
func TestSingleLineParsing(t *testing.T) {
	t.Parallel()
	var linetests = []struct {
		in  string
		out metrics.Metric
	}{
		{
			"foo 1 2",
			*metrics.NewMetric("foo", 1, 2),
		},
		{
			"a.b.c 4.2 42",
			*metrics.NewMetric("a.b.c", 4.2, 42),
		},
	}

	for i, tt := range linetests {
		_, outMetric := parseGraphiteLine(tt.in)
		if outMetric != tt.out {
			t.Errorf("%d. parseGraphiteLine(%s) => %v, want %v", i, tt.in, outMetric, tt.out)
		}
	}
}

func TestSingleLineParserFail(t *testing.T) {
	t.Parallel()
	failLines := []string{
		"",
		"one",
		"two elements",
		"four elements to go",
		"foo 1 should_be_int",
		"name should_be_float 1",
	}

	for i, tt := range failLines {
		err, _ := parseGraphiteLine(tt)
		if err == nil {
			t.Errorf("%d. Expected parseGraphiteLine(%s) to return error", i, tt)
		}
	}
}

/*
// Dummy ReadWriteCloser for testing
type RWC struct{ data []byte }

func (m RWC) Read(b []byte) (int, error) {

	// Copy over as many bytes as we can
	bytes := 0
	for bytes < len(m.data) && bytes < len(b) {
		b[bytes] = m.data[bytes]
		bytes += 1
	}
	m.data = m.data[bytes:]

	//fmt.Println("Writing", b)
	return bytes, io.EOF
}

func (m RWC) Write(b []byte) (int, error) { return 0, nil }
func (m RWC) Close() error                { m.data = make([]byte, 0); return nil }
func (m RWC) String() string              { return string(m.data) }
*/
