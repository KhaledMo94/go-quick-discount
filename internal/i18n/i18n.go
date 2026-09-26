package i18n

import (
	"encoding/json"
	"iter"
	"os"
	"path/filepath"
	"strings"
	//"sync"
)

type Translator struct{
	fallback 	string
	dicts 		map[string]map[string]string
	//mu 			sync.RWMutex
}

func New(langDir, fallback string) (*Translator, error) {
	t := &Translator{
		fallback: fallback,
		dicts:    map[string]map[string]string{},
	}
	entries, err := os.ReadDir(langDir)
	if err != nil {
		return nil, err
	}
	for _, e := range entries {
		if e.IsDir() || filepath.Ext(e.Name()) != ".json" {
			continue
		}
		locale := strings.TrimSuffix(e.Name(), ".json")
		b, err := os.ReadFile(filepath.Join(langDir, e.Name()))
		if err != nil {
			return nil, err
		}
		var raw map[string]any
		if err := json.Unmarshal(b, &raw); err != nil {
			return nil, err
		}
		m := make(map[string]string)
		for key, value := range flatten("", raw) {
			m[key] = value
		}
		t.dicts[locale] = m
	}
	return t, nil
}

func (t *Translator) T(key , locale string) string {
	//t.mu.RLock()
	//defer t.mu.RUnlock()

	if d,ok := t.dicts[locale];ok{
		if v,ok := d[key]; ok {
			return v
		}
	}

	if d,ok := t.dicts[t.fallback]; ok{
		if v,ok := d[key]; ok{
			return v
		}
	}

	return key
}

func flatten(prefix string , v any) iter.Seq2[string,string]{
	return func(yield func(string, string) bool) {
		switch t:=v.(type){
		case string:
			if prefix != "" {
				yield(prefix, t)
			}
		case map[string]any:
			for k,child := range t{
				key := k
				if prefix != ""{
					key = prefix + "." + k
				}

				for fk , fv := range flatten(key,child){
					if !yield(fk,fv){
						return 
					}
				}
			}
		}

		
	}
}