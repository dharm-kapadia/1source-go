package api

type Client struct {
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services
	PartiesApi *PartiesApiService
}

type service struct {
	client *Client
}
