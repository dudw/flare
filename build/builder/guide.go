package builder

func TaskForGuideAssets(src string, dest string) {
	_PrepareDirectory(dest)
	_CopyDirectory(src, dest)
}