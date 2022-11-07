package culture

func ExtractLanguageSlugs(cultures []*Culture) []string {
	out := make([]string, len(cultures))
	for i := range out {
		out[i] = cultures[i].LanguageSlug
	}

	return out
}
