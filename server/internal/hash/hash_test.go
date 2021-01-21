package hash
import (
	"testing"
	_ "fmt"
	"github.com/stretchr/testify/require"
)

func TestGetHash(t *testing.T) {
	req := require.New(t)
	hash := GetHash("hash")
	req.Equal("30163935c002fc4e1200906c3d30a9c4956b4af9f6dcaef1eb4b1fcb8fba69e7a7acdc491ea5b1f2864ea8c01b01580ef09defc3b11b3f183cb21d236f7f1a6b", hash)
}

