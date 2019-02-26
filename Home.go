func Home(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "index/html; charset=utf-8")

	fullData := map[string]interface{}{
		render(w, r, homepage, "homepage_view", fullData)
	}
	
}