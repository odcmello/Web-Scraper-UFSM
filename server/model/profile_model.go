package model

type Profile struct {
	ID        int64  `json:"id"`
	Name      string `json:"nome"`
	JobTitle  string `json:"job_title"`
	Company   string `json:"company"`
	Location  string `json:"location"`
	URL       string `json:"url"`
	Curso     string `json:"curso"`
	AnoEvasao int    `json:"ano_evasao"`
}

var FiltersMapper = map[string]func(filters map[string]interface{}, value interface{}){
	"nome":        func(filters map[string]interface{}, value interface{}) { filters["nome"] = value },
	"empresa":     func(filters map[string]interface{}, value interface{}) { filters["empresa"] = value },
	"localizacao": func(filters map[string]interface{}, value interface{}) { filters["localizacao"] = value },
	"curso":       func(filters map[string]interface{}, value interface{}) { filters["curso"] = value },
	"ano_evasao":  func(filters map[string]interface{}, value interface{}) { filters["ano_evasao"] = value },
}
