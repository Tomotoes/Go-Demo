package SafeMap

import "testing"

func TestSafeMap_Delete(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		sm   *SafeMap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sm.Delete(tt.args.key)
		})
	}
}

func TestSafeMap_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		sm    *SafeMap
		args  args
		want  string
		want1 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.sm.Get(tt.args.key)
			if got != tt.want {
				t.Errorf("SafeMap.Get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SafeMap.Get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSafeMap_Put(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name string
		sm   *SafeMap
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sm.Put(tt.args.key, tt.args.value)
		})
	}
}
