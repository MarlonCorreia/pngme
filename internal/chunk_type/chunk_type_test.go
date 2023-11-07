package chunktype_test

import (
	"bytes"
	chunktype "pngme/internal/chunk_type"
	"testing"
)

func TestTypeFromBytes(t *testing.T) {
	expected := []byte{82, 117, 83, 116}
	c := chunktype.ChunkTypeFromBytes(expected)

	if !bytes.Equal(expected, c.GetBytes()) {
		t.Errorf("Expected %d got %d", expected, c.GetBytes())
	}
}

func TestTypeFromStr(t *testing.T) {
	str := "RuSt"
	c := chunktype.ChunkTypeFromStr(str)
	if !bytes.Equal(c.GetBytes(), []byte(str)) {
		t.Errorf("Expected %d got %d", []byte(str), c.GetBytes())
	}
}

func TestTypeIsCritical(t *testing.T) {
	c := chunktype.ChunkTypeFromStr("RuSt")

	if !c.IsCritical() {
		t.Error("Expected true got false")
	}
}

func TestTypeIsNotCritical(t *testing.T) {
	c := chunktype.ChunkTypeFromStr("ruSt")

	if c.IsCritical() {
		t.Error("Expected false got true")
	}
}

func TestTypeIsPrivate(t *testing.T) {
	c := chunktype.ChunkTypeFromStr("RuSt")
	if !c.IsPrivate() {
		t.Error("Expected true, got false")
	}
}

func TestTypeIsNotPrivate(t *testing.T) {
	c := chunktype.ChunkTypeFromStr("RUSt")
	if c.IsPrivate() {
		t.Error("Expected false, got true")
	}
}

func TestTypeIsReservedValid(t *testing.T) {
	c := chunktype.ChunkTypeFromStr("RuSt")
	if !c.IsReservedValid() {
		t.Error("Expected true got false")
	}
}

func TestTypeIsReservedNotValid(t *testing.T) {
	c := chunktype.ChunkTypeFromStr("Rust")
	if c.IsReservedValid() {
		t.Error("Expected false got true")
	}
}

func TestTypeIsSafeToCopy(t *testing.T) {
	c := chunktype.ChunkTypeFromStr("RuSt")
	if !c.IsSafeToCopy() {
		t.Error("Expected true got false")
	}
}

func TestTypeIsNotSafeToCopy(t *testing.T) {
	c := chunktype.ChunkTypeFromStr("RuST")
	if c.IsSafeToCopy() {
		t.Error("Expected false got true")
	}
}

func TestTypeIsValid(t *testing.T) {
	c := chunktype.ChunkTypeFromStr("RuSt")
	if !c.IsValid() {
		t.Error("Expected true got false")
	}
}

func TestTypeIsNotValid(t *testing.T) {
	c := chunktype.ChunkTypeFromStr("Ru1t")
	if c.IsValid() {
		t.Error("Expected false got true")
	}
}

func TestTypeStr(t *testing.T) {
	typeStr := "RuSt"
	c := chunktype.ChunkTypeFromStr(typeStr)

	if c.GetTypeStr() != typeStr {
		t.Errorf("expected %s got %s", typeStr, c.GetTypeStr())
	}
}
