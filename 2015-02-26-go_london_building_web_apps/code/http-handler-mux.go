func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) { // HL

	// <snip>

	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}
