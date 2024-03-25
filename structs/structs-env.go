package structs

type SecretParams struct {
	Name     string `json:"name"`
	Region   string `json:"region"`
	Field    string `json:"field,omitempty"`
}
