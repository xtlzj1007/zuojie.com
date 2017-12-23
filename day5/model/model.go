package model

type School struct{
	Name string 
	Addr string
}

func NewSchool(name,addr string) *School{
	return &School{
		Name:name,
		Addr:addr,
	}
}

