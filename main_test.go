package main

import (
	"reflect"
	"testing"
	"time"
	"unsafe"
)

func TestNewHashMap(t *testing.T) {
	type testData struct {
		size     uint
		expected *HashMap
	}
	testCases := []testData{
		{size: 4, expected: &HashMap{buckets: make([][8][2]unsafe.Pointer, 4), bitMask: 3}},
		{size: 8, expected: &HashMap{buckets: make([][8][2]unsafe.Pointer, 4), bitMask: 3}},
		{size: 16, expected: &HashMap{buckets: make([][8][2]unsafe.Pointer, 8), bitMask: 7}},
	}
	for _, tc := range testCases {
		result := NewHashMap(tc.size)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("unexpected result in NewHashMap: got %+v,\n want %+v", result, tc.expected)
		}
	}
}

func TestWithHashCRC64(t *testing.T) {
	type testData struct {
		size     uint
		expected *CRC64
	}

	testCases := []testData{
		{size: 4, expected: &CRC64{bitMask: 3, buckets: make([][8][2]unsafe.Pointer, 4), firstEmptyInBucket: make([]int, 4)}},

		{size: 8, expected: &CRC64{bitMask: 3, buckets: make([][8][2]unsafe.Pointer, 4), firstEmptyInBucket: make([]int, 4)}},

		{size: 16, expected: &CRC64{bitMask: 7, buckets: make([][8][2]unsafe.Pointer, 8), firstEmptyInBucket: make([]int, 8)}},
	}

	for _, tc := range testCases {
		result := NewHashMap(tc.size, WithHashCRC64())
		if result.HashMaper.(*CRC64).hash == nil {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): hash is nil")
		}
		if result.HashMaper.(*CRC64).tab == nil {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): hash is nil")
		}
		if result.HashMaper.(*CRC64).bitMask != tc.expected.bitMask {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): got bitMask %v, want %v", result.HashMaper.(*CRC64).bitMask, tc.expected.bitMask)
		}
		if !reflect.DeepEqual(result.HashMaper.(*CRC64).buckets, tc.expected.buckets) {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): got buckets %v, want %v", result.HashMaper.(*CRC64).buckets, tc.expected.buckets)
		}
		if !reflect.DeepEqual(result.HashMaper.(*CRC64).firstEmptyInBucket, tc.expected.firstEmptyInBucket) {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): got firstEmptyInBucket %v, want %v", result.HashMaper.(*CRC64).firstEmptyInBucket, tc.expected.firstEmptyInBucket)
		}
	}
}

func TestWithHashCRC32(t *testing.T) {
	type testData struct {
		size     uint
		expected *CRC32
	}

	testCases := []testData{
		{size: 4, expected: &CRC32{bitMask: 3, buckets: make([][8][2]unsafe.Pointer, 4), firstEmptyInBucket: make([]int, 4)}},
		{size: 8, expected: &CRC32{bitMask: 3, buckets: make([][8][2]unsafe.Pointer, 4), firstEmptyInBucket: make([]int, 4)}},
		{size: 16, expected: &CRC32{bitMask: 7, buckets: make([][8][2]unsafe.Pointer, 8), firstEmptyInBucket: make([]int, 8)}},
	}

	for _, tc := range testCases {
		result := NewHashMap(tc.size, WithHashCRC32())
		if result.HashMaper.(*CRC32).hash == nil {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): hash is nil")
		}
		if result.HashMaper.(*CRC32).tab == nil {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): hash is nil")
		}
		if result.HashMaper.(*CRC32).bitMask != tc.expected.bitMask {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): got bitMask %v, want %v", result.HashMaper.(*CRC32).bitMask, tc.expected.bitMask)
		}
		if !reflect.DeepEqual(result.HashMaper.(*CRC32).buckets, tc.expected.buckets) {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): got buckets %v, want %v", result.HashMaper.(*CRC32).buckets, tc.expected.buckets)
		}
		if !reflect.DeepEqual(result.HashMaper.(*CRC32).firstEmptyInBucket, tc.expected.firstEmptyInBucket) {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): got firstEmptyInBucket %v, want %v", result.HashMaper.(*CRC32).firstEmptyInBucket, tc.expected.firstEmptyInBucket)
		}
	}
}

func TestWithHashCRC16(t *testing.T) {
	type testData struct {
		size     uint
		expected *CRC16
	}

	testCases := []testData{
		{size: 4, expected: &CRC16{bitMask: 3, buckets: make([][8][2]unsafe.Pointer, 4), firstEmptyInBucket: make([]int, 4)}},
		{size: 8, expected: &CRC16{bitMask: 3, buckets: make([][8][2]unsafe.Pointer, 4), firstEmptyInBucket: make([]int, 4)}},
		{size: 16, expected: &CRC16{bitMask: 7, buckets: make([][8][2]unsafe.Pointer, 8), firstEmptyInBucket: make([]int, 8)}},
	}

	for _, tc := range testCases {
		result := NewHashMap(tc.size, WithHashCRC16())
		if result.HashMaper.(*CRC16).hash == nil {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): hash is nil")
		}
		if result.HashMaper.(*CRC16).tab == nil {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): hash is nil")
		}
		if result.HashMaper.(*CRC16).bitMask != tc.expected.bitMask {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): got bitMask %v, want %v", result.HashMaper.(*CRC16).bitMask, tc.expected.bitMask)
		}
		if !reflect.DeepEqual(result.HashMaper.(*CRC16).buckets, tc.expected.buckets) {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): got buckets %v, want %v", result.HashMaper.(*CRC16).buckets, tc.expected.buckets)
		}
		if !reflect.DeepEqual(result.HashMaper.(*CRC16).firstEmptyInBucket, tc.expected.firstEmptyInBucket) {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): got firstEmptyInBucket %v, want %v", result.HashMaper.(*CRC16).firstEmptyInBucket, tc.expected.firstEmptyInBucket)
		}
	}
}

func TestWithHashCRC8(t *testing.T) {
	type testData struct {
		size     uint
		expected *CRC8
	}

	testCases := []testData{
		{size: 4, expected: &CRC8{bitMask: 3, buckets: make([][8][2]unsafe.Pointer, 4), firstEmptyInBucket: make([]int, 4)}},
		{size: 8, expected: &CRC8{bitMask: 3, buckets: make([][8][2]unsafe.Pointer, 4), firstEmptyInBucket: make([]int, 4)}},
		{size: 16, expected: &CRC8{bitMask: 7, buckets: make([][8][2]unsafe.Pointer, 8), firstEmptyInBucket: make([]int, 8)}},
	}

	for _, tc := range testCases {
		result := NewHashMap(tc.size, WithHashCRC8())
		if result.HashMaper.(*CRC8).hash == nil {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): hash is nil")
		}
		if result.HashMaper.(*CRC8).tab == nil {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): hash is nil")
		}
		if result.HashMaper.(*CRC8).bitMask != tc.expected.bitMask {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): got bitMask %v, want %v", result.HashMaper.(*CRC8).bitMask, tc.expected.bitMask)
		}
		if !reflect.DeepEqual(result.HashMaper.(*CRC8).buckets, tc.expected.buckets) {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): got buckets %v, want %v", result.HashMaper.(*CRC8).buckets, tc.expected.buckets)
		}
		if !reflect.DeepEqual(result.HashMaper.(*CRC8).firstEmptyInBucket, tc.expected.firstEmptyInBucket) {
			t.Errorf("unexpected result in NewHashMap(WithHashCRC64): got firstEmptyInBucket %v, want %v", result.HashMaper.(*CRC8).firstEmptyInBucket, tc.expected.firstEmptyInBucket)
		}
	}
}

func TestCRC64_Set(t *testing.T) {
	testHashMap := NewHashMap(16, WithHashCRC64())
	type testData struct {
		key   string
		value interface{}
	}
	testCases := []testData{
		{key: "testKey", value: "testValue"},
		{key: "key", value: "value"},
		{key: "test", value: 15},
		{key: "the last one", value: true},
	}
	for _, tc := range testCases {
		testHashMap.Set(tc.key, tc.value)
		result, ok := testHashMap.Get(tc.key)
		if !ok {
			t.Errorf("unexpected result in CRC64_Set: value should exist")
		}
		if !reflect.DeepEqual(result, tc.value) {
			t.Errorf("unexpected result in CRC64_Set: got %v, want %v", result, tc.value)
		}
	}
}

func TestCRC64_Get(t *testing.T) {
	testHashMap := NewHashMap(16, WithHashCRC64())
	type testData struct {
		key             string
		value           interface{}
		keyWithoutValue string
	}
	testCases := []testData{
		{key: "testKey", value: "testValue", keyWithoutValue: "testValue"},
		{key: "key", value: "value", keyWithoutValue: "value"},
		{key: "test", value: 15, keyWithoutValue: "15"},
		{key: "the last one", value: true, keyWithoutValue: "true"},
	}
	for _, tc := range testCases {
		testHashMap.Set(tc.key, tc.value)
		result, ok := testHashMap.Get(tc.key)
		if !ok {
			t.Errorf("unexpected result in CRC64_Set: value should exist")
		}
		if !reflect.DeepEqual(result, tc.value) {
			t.Errorf("unexpected result in CRC64_Set: got %v, want %v", result, tc.value)
		}
		_, ok = testHashMap.Get(tc.keyWithoutValue)
		if ok {
			t.Errorf("unexpected result in CRC64_Set: value should not exist")
		}
	}
}

func TestCRC32_Set(t *testing.T) {
	testHashMap := NewHashMap(16, WithHashCRC32())
	type testData struct {
		key   string
		value interface{}
	}
	testCases := []testData{
		{key: "testKey", value: "testValue"},
		{key: "key", value: "value"},
		{key: "test", value: 15},
		{key: "the last one", value: true},
	}
	for _, tc := range testCases {
		testHashMap.Set(tc.key, tc.value)
		result, ok := testHashMap.Get(tc.key)
		if !ok {
			t.Errorf("unexpected result in CRC32_Set: value should exist")
		}
		if !reflect.DeepEqual(result, tc.value) {
			t.Errorf("unexpected result in CRC32_Set: got %v, want %v", result, tc.value)
		}
	}
}

func TestCRC32_Get(t *testing.T) {
	testHashMap := NewHashMap(16, WithHashCRC32())
	type testData struct {
		key             string
		value           interface{}
		keyWithoutValue string
	}
	testCases := []testData{
		{key: "testKey", value: "testValue", keyWithoutValue: "testValue"},
		{key: "key", value: "value", keyWithoutValue: "value"},
		{key: "test", value: 15, keyWithoutValue: "15"},
		{key: "the last one", value: true, keyWithoutValue: "true"},
	}
	for _, tc := range testCases {
		testHashMap.Set(tc.key, tc.value)
		result, ok := testHashMap.Get(tc.key)
		if !ok {
			t.Errorf("unexpected result in CRC32_Get: value should exist")
		}
		if !reflect.DeepEqual(result, tc.value) {
			t.Errorf("unexpected result in CRC32_Get: got %v, want %v", result, tc.value)
		}
		_, ok = testHashMap.Get(tc.keyWithoutValue)
		if ok {
			t.Errorf("unexpected result in CRC32_Get: value should not exist")
		}
	}
}

func TestCRC16_Set(t *testing.T) {
	testHashMap := NewHashMap(16, WithHashCRC16())
	type testData struct {
		key   string
		value interface{}
	}
	testCases := []testData{
		{key: "testKey", value: "testValue"},
		{key: "key", value: "value"},
		{key: "test", value: 15},
		{key: "the last one", value: true},
	}
	for _, tc := range testCases {
		testHashMap.Set(tc.key, tc.value)
		result, ok := testHashMap.Get(tc.key)
		if !ok {
			t.Errorf("unexpected result in CRC16_Set: value should exist")
		}
		if !reflect.DeepEqual(result, tc.value) {
			t.Errorf("unexpected result in CRC16_Set: got %v, want %v", result, tc.value)
		}
	}
}

func TestCRC16_Get(t *testing.T) {
	testHashMap := NewHashMap(16, WithHashCRC16())
	type testData struct {
		key             string
		value           interface{}
		keyWithoutValue string
	}
	testCases := []testData{
		{key: "testKey", value: "testValue", keyWithoutValue: "testValue"},
		{key: "key", value: "value", keyWithoutValue: "value"},
		{key: "test", value: 15, keyWithoutValue: "15"},
		{key: "the last one", value: true, keyWithoutValue: "true"},
	}
	for _, tc := range testCases {
		testHashMap.Set(tc.key, tc.value)
		result, ok := testHashMap.Get(tc.key)
		if !ok {
			t.Errorf("unexpected result in CRC16_Get: value should exist")
		}
		if !reflect.DeepEqual(result, tc.value) {
			t.Errorf("unexpected result in CRC16_Get: got %v, want %v", result, tc.value)
		}
		_, ok = testHashMap.Get(tc.keyWithoutValue)
		if ok {
			t.Errorf("unexpected result in CRC16_Get: value should not exist")
		}
	}
}

func TestCRC8_Set(t *testing.T) {
	testHashMap := NewHashMap(16, WithHashCRC8())
	type testData struct {
		key   string
		value interface{}
	}
	testCases := []testData{
		{key: "testKey", value: "testValue"},
		{key: "key", value: "value"},
		{key: "test", value: 15},
		{key: "the last one", value: true},
	}
	for _, tc := range testCases {
		testHashMap.Set(tc.key, tc.value)
		result, ok := testHashMap.Get(tc.key)
		if !ok {
			t.Errorf("unexpected result in CRC8_Set: value should exist")
		}
		if !reflect.DeepEqual(result, tc.value) {
			t.Errorf("unexpected result in CRC8_Set: got %v, want %v", result, tc.value)
		}
	}
}

func TestCRC8_Get(t *testing.T) {
	testHashMap := NewHashMap(16, WithHashCRC8())
	type testData struct {
		key             string
		value           interface{}
		keyWithoutValue string
	}
	testCases := []testData{
		{key: "testKey", value: "testValue", keyWithoutValue: "testValue"},
		{key: "key", value: "value", keyWithoutValue: "value"},
		{key: "test", value: 15, keyWithoutValue: "15"},
		{key: "the last one", value: true, keyWithoutValue: "true"},
	}
	for _, tc := range testCases {
		testHashMap.Set(tc.key, tc.value)
		result, ok := testHashMap.Get(tc.key)
		if !ok {
			t.Errorf("unexpected result in CRC8_Get: value should exist")
		}
		if !reflect.DeepEqual(result, tc.value) {
			t.Errorf("unexpected result in CRC8_Get: got %v, want %v", result, tc.value)
		}
		_, ok = testHashMap.Get(tc.keyWithoutValue)
		if ok {
			t.Errorf("unexpected result in CRC8_Get: value should not exist")
		}
	}
}

func TestMeasureTime(t *testing.T) {
	type testData struct {
		function    func()
		expectedMin time.Duration
		expectedMax time.Duration
	}
	testCases := []testData{
		{function: func() { time.Sleep(10 * time.Millisecond) }, expectedMin: 10 * time.Millisecond, expectedMax: 11 * time.Millisecond},
		{function: func() { time.Sleep(20 * time.Millisecond) }, expectedMin: 20 * time.Millisecond, expectedMax: 21 * time.Millisecond},
		{function: func() { time.Sleep(30 * time.Millisecond) }, expectedMin: 30 * time.Millisecond, expectedMax: 31 * time.Millisecond},
		{function: func() { time.Sleep(40 * time.Millisecond) }, expectedMin: 40 * time.Millisecond, expectedMax: 41 * time.Millisecond},
	}
	for _, tc := range testCases {
		result := MeasureTime(tc.function)
		if result < tc.expectedMin {
			t.Errorf("unexpected result in Measure time: got %v, want res < %v", result, tc.expectedMin)
		}
		if result > tc.expectedMax {
			t.Errorf("unexpected result in Measure time: got %v, want res > %v", result, tc.expectedMax)
		}
	}
}

func BenchmarkCRC64_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testHashMap := NewHashMap(16, WithHashCRC64())
		testHashMap.Set("key", "value")
	}
}

func BenchmarkCRC64_Get(b *testing.B) {
	testHashMap := NewHashMap(16, WithHashCRC64())
	testHashMap.Set("key", "value")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testHashMap.Get("key")
	}
}

func BenchmarkCRC32_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testHashMap := NewHashMap(16, WithHashCRC32())
		testHashMap.Set("key", "value")
	}
}

func BenchmarkCRC32_Get(b *testing.B) {
	testHashMap := NewHashMap(16, WithHashCRC32())
	testHashMap.Set("key", "value")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testHashMap.Get("key")
	}
}

func BenchmarkCRC16_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testHashMap := NewHashMap(16, WithHashCRC16())
		testHashMap.Set("key", "value")
	}
}

func BenchmarkCRC16_Get(b *testing.B) {
	testHashMap := NewHashMap(16, WithHashCRC16())
	testHashMap.Set("key", "value")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testHashMap.Get("key")
	}
}

func BenchmarkCRC8_Set(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testHashMap := NewHashMap(16, WithHashCRC8())
		testHashMap.Set("key", "value")
	}
}

func BenchmarkCRC8_Get(b *testing.B) {
	testHashMap := NewHashMap(16, WithHashCRC8())
	testHashMap.Set("key", "value")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testHashMap.Get("key")
	}
}
