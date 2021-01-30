package gsebleve

import (
	"os"

	"github.com/blevesearch/bleve/v2/analysis"
	"github.com/blevesearch/bleve/v2/registry"
	"github.com/go-ego/gse"
)

const (
	Name = "gse"
)

type GseTokenizer struct {
	segmenter *gse.Segmenter
}

func NewGseTokenizer(dictFiles string) *GseTokenizer {
	// segmenter := gse.New("./data/dict/zh/dict.txt", dictFiles)
	// segmenter.MoreLog = false
	// segmenter.SkipLog = true
	dictFile := "./data/dict/zh/dict.txt"
	var segmenter gse.Segmenter
	segmenter.SkipLog = true
	if IsExist(dictFile) {
		segmenter.LoadDict(dictFile, dictFiles)
	}
	return &GseTokenizer{&segmenter}
}

/* func (t *GseTokenizer) Free()  {

} */

func (t *GseTokenizer) Tokenize(sentence []byte) analysis.TokenStream {
	result := make(analysis.TokenStream, 0)
	pos := 1
	//segments := t.segmenter.ModeSegment(sentence, true)
	segments := t.segmenter.Segment(sentence)
	for _, seg := range segments {
		token := analysis.Token{
			Term:     []byte(seg.Token().Text()),
			Start:    seg.Start(),
			End:      seg.End(),
			Position: pos,
			Type:     analysis.Ideographic,
		}
		result = append(result, &token)
		pos++
	}
	return result
}

func tokenizerConstructor(config map[string]interface{}, cache *registry.Cache) (analysis.Tokenizer, error) {
	dicts := ""
	userDicts, ok := config["user_dicts"].(string)
	if !ok {
		dicts = userDicts
	}
	return NewGseTokenizer(dicts), nil
}

func init() {
	registry.RegisterTokenizer(Name, tokenizerConstructor)
}

// IsExist checks if the given file exist
func IsExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}
