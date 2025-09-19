package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	type args struct {
		rows int
	}
	tests := []struct {
		name string
		args args
		want Spreadsheet
	}{
		{
			name: "init spreadsheet",
			args: args{rows: 3},
			want: Spreadsheet{spreadSheet: map[string]int{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Constructor(tt.args.rows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpreadsheet_SetCell(t *testing.T) {
	type fields struct {
		spreadSheet map[string]int
	}
	type args struct {
		cell  string
		value int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "set cell A1 to 10",
			fields: fields{spreadSheet: map[string]int{}},
			args:   args{cell: "A1", value: 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Spreadsheet{
				spreadSheet: tt.fields.spreadSheet,
			}
			this.SetCell(tt.args.cell, tt.args.value)
			if got := this.spreadSheet[tt.args.cell]; got != tt.args.value {
				t.Errorf("SetCell() = %v, want %v", got, tt.args.value)
			}
		})
	}
}

func TestSpreadsheet_ResetCell(t *testing.T) {
	type fields struct {
		spreadSheet map[string]int
	}
	type args struct {
		cell string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "reset cell A1",
			fields: fields{spreadSheet: map[string]int{"A1": 10}},
			args:   args{cell: "A1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Spreadsheet{
				spreadSheet: tt.fields.spreadSheet,
			}
			this.ResetCell(tt.args.cell)
			if got := this.spreadSheet[tt.args.cell]; got != 0 {
				t.Errorf("ResetCell() = %v, want %v", got, 0)
			}
		})
	}
}

func TestSpreadsheet_GetValue(t *testing.T) {
	type fields struct {
		spreadSheet map[string]int
	}
	type args struct {
		formula string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "sum of two integers",
			fields: fields{spreadSheet: map[string]int{}},
			args:   args{formula: "=5+7"},
			want:   12,
		},
		{
			name:   "sum of cell and integer",
			fields: fields{spreadSheet: map[string]int{"A1": 10}},
			args:   args{formula: "=A1+6"},
			want:   16,
		},
		{
			name:   "sum of two cells",
			fields: fields{spreadSheet: map[string]int{"A1": 10, "B2": 15}},
			args:   args{formula: "=A1+B2"},
			want:   25,
		},
		{
			name:   "sum after reset cell",
			fields: fields{spreadSheet: map[string]int{"A1": 0, "B2": 15}},
			args:   args{formula: "=A1+B2"},
			want:   15,
		},
		{
			name:   "cell not set (should be 0)",
			fields: fields{spreadSheet: map[string]int{}},
			args:   args{formula: "=A1+B2"},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Spreadsheet{
				spreadSheet: tt.fields.spreadSheet,
			}
			if got := this.GetValue(tt.args.formula); got != tt.want {
				t.Errorf("Spreadsheet.GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
