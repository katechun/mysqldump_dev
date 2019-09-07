package main

func main() {
	type Export interface {
		SelectRows()
		WriteFile()
	}

	type Import interface {
		ReadFile()
		InsertRows()
	}

}
