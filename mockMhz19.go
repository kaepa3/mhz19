package mhz19

type MockMHZ19 struct {
	Value int
}

func (m *MockMHZ19) Connect() error {
	return nil
}

func (m *MockMHZ19) ReadCO2() (int, error) {
	return m.Value, nil
}
