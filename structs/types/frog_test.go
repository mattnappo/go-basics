package types

import (
	"testing"
)

func TestNewFrog(t *testing.T) {
	MyFroggy, err := NewFrog(
		"big hunk",
		700,
		true,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(MyFroggy)
}

func TestRee(t *testing.T) {
	MyFroggy, err := NewFrog(
		"big hunk",
		700,
		true,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = MyFroggy.Ree()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")
}

func TestWriteToMemory(t *testing.T) {
	MyFroggy, err := NewFrog(
		"big hunk",
		700,
		true,
	)
	if err != nil {
		t.Fatal(err)
	}

	err = MyFroggy.WriteToMemory()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success")

}

func TestReadFrogFromMemory(t *testing.T) {
	myFroggy, err := NewFrog(
		"big hunk",
		700,
		true,
	)
	if err != nil {
		t.Fatal(err)
	}

	err = myFroggy.WriteToMemory()
	if err != nil {
		t.Fatal(err)
	}
	myFroggy, err = ReadFrogFromMemory(myFroggy.Name)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(myFroggy)
}

func TestString(t *testing.T) {
	myFroggy, err := NewFrog(
		"big hunk",
		700,
		true,
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(myFroggy.String())
}

func TestBytes(t *testing.T) {
	myFroggy, err := NewFrog(
		"big hunk",
		700,
		true,
	)

	if err != nil {
		t.Fatal(err)
	}

	t.Log(myFroggy.Bytes())
}

func TestFrogFromBytes(t *testing.T) {
	myFroggy, err := NewFrog(
		"big hunk",
		700,
		true,
	)

	if err != nil {
		t.Fatal(err)
	}

	byteVal := myFroggy.Bytes()

	myFroggy, err = FrogFromBytes(byteVal) // Decode byte value

	if err != nil {
		t.Fatal(err)
	}

	t.Log(myFroggy)
}
