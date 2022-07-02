package common

import (
	"fmt"
	"testing"

	"github.com/speedwheel/bitget-go/internal"
)

func TestSigner_Sign(t *testing.T) {
	signer := new(Signer)
	result := signer.Sign("GET", "www.bitget.com", "aaaa", internal.TimesStamp())
	fmt.Print(result)
}
