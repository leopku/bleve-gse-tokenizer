package gsebleve

import (
	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/registry"
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
	var segmenter gse.Segmenter
	segmenter.SkipLog = true
	segmenter.LoadDict(dictFiles)
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
	dicts, ok := config["user_dicts"].(string)
	if !ok || dicts == "" {
		panic("Initialize Gse tokenizer error: user dicts should NOT be empty")
	}
	return NewGseTokenizer(dicts), nil
}

func init() {
	registry.RegisterTokenizer(Name, tokenizerConstructor)
}
