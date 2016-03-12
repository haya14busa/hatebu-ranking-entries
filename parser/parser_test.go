package parser

import "testing"

func TestNewEntries(t *testing.T) {

	tests := []struct {
		descr   string
		in      string
		wantErr bool
	}{{
		descr: "daily ranking page",
		in:    "http://b.hatena.ne.jp/ranking/daily",
	}, {
		descr: "weekly ranking page",
		in:    "http://b.hatena.ne.jp/ranking/weekly",
	}, {
		descr: "monthly ranking page",
		in:    "http://b.hatena.ne.jp/ranking/monthly",
	}, {
		descr: "daily it ranking page",
		in:    "http://b.hatena.ne.jp/ranking/daily/20160311/it",
	}, {
		descr: "daily it weekly page",
		in:    "http://b.hatena.ne.jp/ranking/weekly/20160229/it",
	}, {
		descr:   "invalid url",
		in:      "http://invalid/ranking/daily/",
		wantErr: true,
	}, {
		descr:   "not found url",
		in:      "http://b.hatena.ne.jp/ranking/daily/notfound",
		wantErr: true,
	}}

	for _, tt := range tests {
		es, err := NewEntries(tt.in)
		if err == nil && tt.wantErr {
			t.Errorf("%s: error expected but no error. in: %v", tt.descr, tt.in)
		}
		if err != nil && !tt.wantErr {
			t.Errorf("%s: Fail to get entries from %s: %v", tt.descr, tt.in, err)
		}
		if err != nil {
			continue
		}
		if len(es) < 1 {
			t.Errorf("%s: entries not found. in: %v", tt.descr, tt.in)
		}
		for _, e := range es {
			if e.Eid == "" {
				t.Errorf("%s: entry should have id: %s", tt.descr, tt.in)
			}
		}
	}

}
