package DataStructure

import (
	"reflect"
	"sync"
	"testing"
)

func TestSynchronizedMap_New(t *testing.T) {
	type fields struct {
		rw   *sync.RWMutex
		data map[interface{}]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   *SynchronizedMap
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SynchronizedMap{
				rw:   tt.fields.rw,
				data: tt.fields.data,
			}
			if got := sm.New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SynchronizedMap.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSynchronizedMap_Put(t *testing.T) {
	type fields struct {
		rw   *sync.RWMutex
		data map[interface{}]interface{}
	}
	type args struct {
		in1 interface{}
		in2 interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SynchronizedMap{
				rw:   tt.fields.rw,
				data: tt.fields.data,
			}
			sm.Put(tt.args.in1, tt.args.in2)
		})
	}
}

func TestSynchronizedMap_Get(t *testing.T) {
	type fields struct {
		rw   *sync.RWMutex
		data map[interface{}]interface{}
	}
	type args struct {
		k interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SynchronizedMap{
				rw:   tt.fields.rw,
				data: tt.fields.data,
			}
			if got := sm.Get(tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SynchronizedMap.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSynchronizedMap_Delete(t *testing.T) {
	type fields struct {
		rw   *sync.RWMutex
		data map[interface{}]interface{}
	}
	type args struct {
		k interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SynchronizedMap{
				rw:   tt.fields.rw,
				data: tt.fields.data,
			}
			sm.Delete(tt.args.k)
		})
	}
}

func TestSynchronizedMap_Each(t *testing.T) {
	type fields struct {
		rw   *sync.RWMutex
		data map[interface{}]interface{}
	}
	type args struct {
		cb func(interface{}, interface{})
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SynchronizedMap{
				rw:   tt.fields.rw,
				data: tt.fields.data,
			}
			sm.Each(tt.args.cb)
		})
	}
}

func TestSynchronizedMap_Map(t *testing.T) {
	type fields struct {
		rw   *sync.RWMutex
		data map[interface{}]interface{}
	}
	type args struct {
		handle func(interface{}, interface{}) interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SynchronizedMap{
				rw:   tt.fields.rw,
				data: tt.fields.data,
			}
			sm.Map(tt.args.handle)
		})
	}
}

func TestSynchronizedMap_Filter(t *testing.T) {
	type fields struct {
		rw   *sync.RWMutex
		data map[interface{}]interface{}
	}
	type args struct {
		filter func(interface{}, interface{}) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SynchronizedMap{
				rw:   tt.fields.rw,
				data: tt.fields.data,
			}
			sm.Filter(tt.args.filter)
		})
	}
}

func TestSynchronizedMap_Size(t *testing.T) {
	type fields struct {
		rw   *sync.RWMutex
		data map[interface{}]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"test1", fields{rw: &sync.RWMutex{}, data: map[interface{}]interface{}{},}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SynchronizedMap{
				rw:   tt.fields.rw,
				data: tt.fields.data,
			}
			if got := sm.Size(); got != tt.want {
				t.Errorf("SynchronizedMap.Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSynchronizedMap_Keys(t *testing.T) {
	type fields struct {
		rw   *sync.RWMutex
		data map[interface{}]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SynchronizedMap{
				rw:   tt.fields.rw,
				data: tt.fields.data,
			}
			if got := sm.Keys(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SynchronizedMap.Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSynchronizedMap_Values(t *testing.T) {
	type fields struct {
		rw   *sync.RWMutex
		data map[interface{}]interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sm := &SynchronizedMap{
				rw:   tt.fields.rw,
				data: tt.fields.data,
			}
			if got := sm.Values(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SynchronizedMap.Values() = %v, want %v", got, tt.want)
			}
		})
	}
}
