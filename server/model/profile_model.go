package model

type Profile struct {
	ID        int64
	Name      string
	JobTitle  string
	Company   string
	Location  string
	URL       string
	Curso     string
	AnoEvasao int
}

var FiltersMapper = map[string]func(filters map[string]interface{}, value interface{}){
	"nome":        func(filters map[string]interface{}, value interface{}) { filters["nome"] = value },
	"empresa":     func(filters map[string]interface{}, value interface{}) { filters["empresa"] = value },
	"localizacao": func(filters map[string]interface{}, value interface{}) { filters["localizacao"] = value },
	"curso":       func(filters map[string]interface{}, value interface{}) { filters["curso"] = value },
	"ano_evasao":  func(filters map[string]interface{}, value interface{}) { filters["ano_evasao"] = value },
}
