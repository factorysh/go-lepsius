package filter

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"hash"
	"io"

	"github.com/factorysh/go-lepsius/tick/model"
)

var hashes map[string]func() hash.Hash

func init() {
	hashes = map[string]func() hash.Hash{
		"sha1": sha1.New,
	}
}

func DoFingerprintFilter(method string, format string, sourceList []string, target string, in *model.Line) error {
	h, ok := hashes[method]
	if !ok {
		return fmt.Errorf("Hash method not found : %s", method)
	}
	hh := h()
	for _, s := range sourceList {
		io.WriteString(hh, fmt.Sprintf("%v", in.Data[s]))
	}
	var v interface{}
	switch format {
	case "base64":
		v = base64.StdEncoding.EncodeToString(hh.Sum(nil))
	default:
		v = hh.Sum(nil)
	}
	in.Data[target] = v
	return nil
}
