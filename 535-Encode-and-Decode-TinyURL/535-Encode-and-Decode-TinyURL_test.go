package leetcode

import (
	"reflect"
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name string
		want Codec
	}{
		{
			name: "Test case 1: Verify Constructor initializes empty Codec",
			want: Codec{
				lookup:        map[string]string{},
				reverseLookup: map[string]string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Constructor()

			// Verify that the maps are properly initialized and empty
			if got.lookup == nil || len(got.lookup) != 0 {
				t.Errorf("Constructor() = lookup map not properly initialized")
			}
			if got.reverseLookup == nil || len(got.reverseLookup) != 0 {
				t.Errorf("Constructor() = reverseLookup map not properly initialized")
			}

			// Ensure the structure matches the expected structure
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Constructor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_encode(t *testing.T) {
	type fields struct {
		lookup        map[string]string
		reverseLookup map[string]string
	}
	type args struct {
		longUrl string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "Test case 1: Basic encoding",
			fields: fields{lookup: map[string]string{}, reverseLookup: map[string]string{}},
			args:   args{longUrl: "https://leetcode.com/problems/design-tinyurl"},
			want:   "https://silentFellow/XXXXXXXXX", // Match based on generated random value
		},
		{
			name: "Test case 2: URL already encoded",
			fields: fields{
				lookup: map[string]string{
					"https://silentFellow/ABCDEFGH": "https://example.com",
				},
				reverseLookup: map[string]string{
					"https://example.com": "https://silentFellow/ABCDEFGH",
				},
			},
			args: args{longUrl: "https://example.com"},
			want: "https://silentFellow/ABCDEFGH",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{
				lookup:        tt.fields.lookup,
				reverseLookup: tt.fields.reverseLookup,
			}
			got := this.encode(tt.args.longUrl)
			if got == "" {
				t.Errorf("Codec.encode() = empty string, which is invalid")
			}
			if tt.args.longUrl != this.decode(got) {
				t.Errorf("Mismatch between encoded and decoded values")
			}
		})
	}
}

func TestCodec_decode(t *testing.T) {
	type fields struct {
		lookup        map[string]string
		reverseLookup map[string]string
	}
	type args struct {
		shortUrl string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Test case 1: Valid decoding",
			fields: fields{
				lookup: map[string]string{
					"https://silentFellow/XYZ123": "https://google.com",
				},
				reverseLookup: map[string]string{
					"https://google.com": "https://silentFellow/XYZ123",
				},
			},
			args: args{shortUrl: "https://silentFellow/XYZ123"},
			want: "https://google.com",
		},
		{
			name:   "Test case 2: Non-existent short URL",
			fields: fields{lookup: map[string]string{}, reverseLookup: map[string]string{}},
			args:   args{shortUrl: "https://silentFellow/INVALID"},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{
				lookup:        tt.fields.lookup,
				reverseLookup: tt.fields.reverseLookup,
			}
			if got := this.decode(tt.args.shortUrl); got != tt.want {
				t.Errorf("Codec.decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateRandom(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{name: "Test case 1: Length 5", length: 5},
		{name: "Test case 2: Length 10", length: 10},
		{name: "Test case 3: Length 15", length: 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandom(tt.length)
			if len(got) != tt.length {
				t.Errorf(
					"generateRandom() = %v, length mismatch, got length = %d, want length = %d",
					got,
					len(got),
					tt.length,
				)
			}
		})
	}
}
