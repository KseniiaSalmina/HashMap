package main

import (
	"fmt"
	"github.com/sigurn/crc16"
	"github.com/sigurn/crc8"
	"hash/crc32"
	"hash/crc64"
	"math/bits"
	"time"
	"unsafe"
)

type HashMaper interface {
	Set(key string, value interface{})
	Get(key string) (interface{}, bool)
}

type CRC64 struct {
	hash               func(data []byte, tab *crc64.Table) uint64
	tab                *crc64.Table
	bitMask            uint64
	buckets            [][8][2]unsafe.Pointer
	firstEmptyInBucket []int
}

func (crc64 *CRC64) Set(key string, value interface{}) {
	hashKey := crc64.hash([]byte(key), crc64.tab)
	numBucket := hashKey & crc64.bitMask
	position := crc64.firstEmptyInBucket[numBucket]
	if position >= 8 {
		fmt.Printf("Bucket %v is full, cannot add value!\n", numBucket)
		return
	}
	crc64.buckets[numBucket][position][0] = unsafe.Pointer(&key)
	crc64.buckets[numBucket][position][1] = unsafe.Pointer(&value)
	crc64.firstEmptyInBucket[numBucket]++
}

func (crc64 *CRC64) Get(key string) (interface{}, bool) {
	hashKey := crc64.hash([]byte(key), crc64.tab)
	numBucket := hashKey & crc64.bitMask
	for i := 0; i < crc64.firstEmptyInBucket[numBucket]; i++ {
		if key == *(*string)(crc64.buckets[numBucket][i][0]) {
			return *(*interface{})(crc64.buckets[numBucket][i][1]), true
		}
	}
	return "", false
}

type CRC32 struct {
	hash               func(data []byte, tab *crc32.Table) uint32
	tab                *crc32.Table
	bitMask            uint32
	buckets            [][8][2]unsafe.Pointer
	firstEmptyInBucket []int
}

func (crc32 *CRC32) Set(key string, value interface{}) {
	hashKey := crc32.hash([]byte(key), crc32.tab)
	numBucket := hashKey & crc32.bitMask
	position := crc32.firstEmptyInBucket[numBucket]
	if position >= 8 {
		fmt.Printf("Bucket %v is full, cannot add value!\n", numBucket)
		return
	}
	crc32.buckets[numBucket][position][0] = unsafe.Pointer(&key)
	crc32.buckets[numBucket][position][1] = unsafe.Pointer(&value)
	crc32.firstEmptyInBucket[numBucket]++
}

func (crc32 *CRC32) Get(key string) (interface{}, bool) {
	hashKey := crc32.hash([]byte(key), crc32.tab)
	numBucket := hashKey & crc32.bitMask
	for i := 0; i < crc32.firstEmptyInBucket[numBucket]; i++ {
		if key == *(*string)(crc32.buckets[numBucket][i][0]) {
			return *(*interface{})(crc32.buckets[numBucket][i][1]), true
		}
	}
	return "", false
}

type CRC16 struct {
	hash               func(data []byte, tab *crc16.Table) uint16
	tab                *crc16.Table
	bitMask            uint16
	buckets            [][8][2]unsafe.Pointer
	firstEmptyInBucket []int
}

func (crc16 *CRC16) Set(key string, value interface{}) {
	hashKey := crc16.hash([]byte(key), crc16.tab)
	numBucket := hashKey & crc16.bitMask
	position := crc16.firstEmptyInBucket[numBucket]
	if position >= 8 {
		fmt.Printf("Bucket %v is full, cannot add value!\n", numBucket)
		return
	}
	crc16.buckets[numBucket][position][0] = unsafe.Pointer(&key)
	crc16.buckets[numBucket][position][1] = unsafe.Pointer(&value)
	crc16.firstEmptyInBucket[numBucket]++
}

func (crc16 *CRC16) Get(key string) (interface{}, bool) {
	hashKey := crc16.hash([]byte(key), crc16.tab)
	numBucket := hashKey & crc16.bitMask
	for i := 0; i < crc16.firstEmptyInBucket[numBucket]; i++ {
		if key == *(*string)(crc16.buckets[numBucket][i][0]) {
			return *(*interface{})(crc16.buckets[numBucket][i][1]), true
		}
	}
	return "", false
}

type CRC8 struct {
	hash               func(data []byte, tab *crc8.Table) uint8
	tab                *crc8.Table
	bitMask            uint8
	buckets            [][8][2]unsafe.Pointer
	firstEmptyInBucket []int
}

func (crc8 *CRC8) Set(key string, value interface{}) {
	hashKey := crc8.hash([]byte(key), crc8.tab)
	numBucket := hashKey & crc8.bitMask
	position := crc8.firstEmptyInBucket[numBucket]
	if position >= 8 {
		fmt.Printf("Bucket %v is full, cannot add value!\n", numBucket)
		return
	}
	crc8.buckets[numBucket][position][0] = unsafe.Pointer(&key)
	crc8.buckets[numBucket][position][1] = unsafe.Pointer(&value)
	crc8.firstEmptyInBucket[numBucket]++
}

func (crc8 *CRC8) Get(key string) (interface{}, bool) {
	hashKey := crc8.hash([]byte(key), crc8.tab)
	numBucket := hashKey & crc8.bitMask
	for i := 0; i < crc8.firstEmptyInBucket[numBucket]; i++ {
		if key == *(*string)(crc8.buckets[numBucket][i][0]) {
			return *(*interface{})(crc8.buckets[numBucket][i][1]), true
		}
	}
	return "", false
}

type HashMap struct {
	HashMaper
	bitMask int
	buckets [][8][2]unsafe.Pointer
}

type HashMapOption func(hm *HashMap)

func WithHashCRC64() HashMapOption {
	return func(hm *HashMap) {
		HashCRC64 := &CRC64{hash: crc64.Checksum, tab: crc64.MakeTable(crc64.ISO), bitMask: uint64(hm.bitMask),
			buckets: hm.buckets, firstEmptyInBucket: make([]int, len(hm.buckets))}
		hm.HashMaper = HashCRC64
	}
}

func WithHashCRC32() HashMapOption {
	return func(hm *HashMap) {
		HashCRC32 := &CRC32{hash: crc32.Checksum, tab: crc32.MakeTable(crc32.IEEE), bitMask: uint32(hm.bitMask),
			buckets: hm.buckets, firstEmptyInBucket: make([]int, len(hm.buckets))}
		hm.HashMaper = HashCRC32
	}
}

func WithHashCRC16() HashMapOption {
	return func(hm *HashMap) {
		HashCRC16 := &CRC16{hash: crc16.Checksum, tab: crc16.MakeTable(crc16.CRC16_ARC), bitMask: uint16(hm.bitMask),
			buckets: hm.buckets, firstEmptyInBucket: make([]int, len(hm.buckets))}
		hm.HashMaper = HashCRC16
	}
}

func WithHashCRC8() HashMapOption {
	return func(hm *HashMap) {
		HashCRC8 := &CRC8{hash: crc8.Checksum, tab: crc8.MakeTable(crc8.CRC8), bitMask: uint8(hm.bitMask),
			buckets: hm.buckets, firstEmptyInBucket: make([]int, len(hm.buckets))}
		hm.HashMaper = HashCRC8
	}
}

func NewHashMap(size uint, options ...HashMapOption) *HashMap {
	minimumAmountOfBuckets := (size >> 2) + 1
	bitMask := (1 << bits.Len(minimumAmountOfBuckets)) - 1
	actualAmountOfBuckets := bitMask + 1
	hm := &HashMap{buckets: make([][8][2]unsafe.Pointer, actualAmountOfBuckets), bitMask: bitMask}
	for _, option := range options {
		option(hm)
	}
	return hm
}

func MeasureTime(f func()) time.Duration {
	start := time.Now()
	f()
	return time.Since(start)
}
