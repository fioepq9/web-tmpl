package config

type Mode string

const (
	ModeDev Mode = "dev"
	ModeStg Mode = "stg"
	ModePrd Mode = "prd"
)

func (m Mode) Dev() bool {
	return !m.Stg() && !m.Prd()
}

func (m Mode) Stg() bool {
	return m == ModeStg
}

func (m Mode) Prd() bool {
	return m == ModePrd
}
