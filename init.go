var homepage *template.Template
func init() {
homepageHTML := assets.MustAssetString("templates/index.html")
	homepage = template.Must(template.New("homepage_view").Parse(homepageHTML))
}