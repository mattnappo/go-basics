package networking

import "testing"

func TestGetGenesis(t *testing.T) {
	genesis := GetGenesis()
	t.Log(genesis)
}
