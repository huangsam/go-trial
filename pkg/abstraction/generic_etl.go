package abstraction

// RunETL is a generic function that performs Extract, Transform, Load (ETL) operations.
//
// It takes three functions as parameters:
//
//   - A function that returns a channel of type A, from which data is extracted.
//   - A function that transforms data of type A into type B.
//   - A function that loads data of type B.
func RunETL[A any, B any](extractor func() <-chan A, transformer func(A) B, loader func(B)) {
	// Extract
	dataChannel := extractor()

	// Transform and Load
	for data := range dataChannel {
		transformedData := transformer(data)
		loader(transformedData)
	}
}
