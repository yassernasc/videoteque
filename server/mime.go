package server

import "videoteque/fs"

func getMime(filename string) string {
	// https://developer.mozilla.org/en-US/docs/Web/Media/Formats/Containers#browser_compatibility
	switch fs.Ext(filename) {
	case ".3gp":
		return "video/3gpp"
	case ".m4p", ".m4v", ".mp4":
		return "video/mp4"
	case ".mpeg", ".mpg":
		return "video/mpeg"
	case ".ogg", ".ogv":
		return "video/ogg"
	case ".webm":
		return "video/webm"
	default:
		return "video/mp4" // use mp4 mime as fallback
	}
}
