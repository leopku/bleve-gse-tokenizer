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
	segmenter := gse.New("zh", dictFiles)
	return &GseTokenizer{&segmenter}
}

/* func (t *GseTokenizer) Free()  {

} */

func (t *GseTokenizer) Tokenize(sentence []byte) analysis.TokenStream {
	result := make(analysis.TokenStream, 0)
	pos := 1
	segments := t.segmenter.ModeSegment(sentence, true)
	for _, seg := range segments {
		// if strings.TrimSpace(seg.Token().Text()) == "" {
		// 	continue
		// }
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
	userDicts, ok := config["user_dicts"].(string)
	if !ok {
		return NewGseTokenizer(""), nil
	}
	return NewGseTokenizer(userDicts), nil
}

func init() {
	registry.RegisterTokenizer(Name, tokenizerConstructor)
}
