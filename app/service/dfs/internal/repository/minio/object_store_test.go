package minio

import "testing"

func TestAccessHashRoundTripStorageFileType(t *testing.T) {
	storageType := int32(0x007efe0e)
	accessHash := MakeAccessHash(storageType, 0x01020304)
	if got := StorageTypeFromAccessHash(accessHash); got != storageType {
		t.Fatalf("storage type = %#x, want %#x", got, storageType)
	}
}

func TestObjectPathUsesDatSuffix(t *testing.T) {
	if got := ObjectPath(12345); got != "12345.dat" {
		t.Fatalf("ObjectPath() = %q, want %q", got, "12345.dat")
	}
}
