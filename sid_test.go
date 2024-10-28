package sid

import "testing"

func TestNew(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "new 一个id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId := New()
			t.Logf("New() = %v", gotId)
			t.Logf("New() = %d", gotId)
		})
	}
}

func TestFromNum(t *testing.T) {
	type args struct {
		num uint64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "从数字创建",
			args: args{
				num: 529437569033699329,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId := FromNum(tt.args.num)
			t.Logf("FromNum() = %v", gotId)
			t.Logf("FromNum() = %d", gotId)
		})
	}
}

func TestFromStr(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name   string
		args   args
		wantId Id
	}{
		{
			name: "从字符串创建",
			args: args{
				str: "AQAA-anwWAc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId := FromStr(tt.args.str)
			t.Logf("FromStr() = %v", gotId)
			t.Logf("FromStr() = %d", gotId)
		})
	}
}
