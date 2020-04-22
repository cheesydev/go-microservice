package product

// The func that searches and returns the rating of a product. This function
// could talk to a database of a cache, and would be called by the http or grpc
// servers providing the service in the different protocols.
func ratingByProductId(id string) (string, error) {
	return "4", nil
}
