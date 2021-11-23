package mhz19

type MocMHZ19 struct {
	Value int
}

func (m *MocMHZ19) Connect() error {
	return nil
}

func (m *MocMHZ19) ReadCO2() (int, error) {
	return m.Value, nil
}
