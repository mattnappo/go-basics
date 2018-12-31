package types

func TestWriteBlockToMemory(t *testing.T) {
	tx1, err := NewTransaction(
		"dowland",
		"matt",
		10000000,
	)

	if err != nil {
		t.Fatal(err)
	}

	transactions := []*Transaction{tx1, tx1}

	block, err := NewBlock(
		0,
		transactions,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = block.WriteBlockToMemory()
	if err != nil {
		t.Fatal(err)
	}
}

func TestReadBlockFromMemory(t *testing.T) {
	tx1, err := NewTransaction(
		"dowland",
		"matt",
		10000000,
	)

	if err != nil {
		t.Fatal(err)
	}

	transactions := []*Transaction{tx1, tx1}

	block, err := NewBlock(
		0,
		transactions,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = block.WriteBlockToMemory()
	if err != nil {
		t.Fatal(err)
	}

	// t.Log(fmt.Sprintf("here: %s\n", ))
	block, err = ReadBlockFromMemory(fmt.Sprintf("%x", block.Hash[:8]))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(block)
}
