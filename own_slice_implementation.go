package main

type mySlice struct {
	data []int
}

func (m *mySlice) cap(data []int) int {
	return 
}

func (m *mySlice) append(data ...int) {
	m.data = append(m.data, data...)
}

func (m *mySlice) remove(index int) bool {
	if index >= 0 && index < len(m.data) {
		m.data = append(m.data[:index], m.data[index+1:]...)

		return true
	}

	return false
}

func (m *mySlice) get(index int) int {
	if index >= 0 && index < len(m.data) {
		return m.data[index]
	}

	return -1
}