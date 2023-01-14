package downloader

// DownloadFromUrl downloads files from urls using downloader.
func DownloadFromUrl(downloadDir string, urls map[string]string) error {
	var u []string
	for _, v := range urls {
		u = append(u, v)
	}
	r, err := NewDownloader(downloadDir, u)
	if err != nil {
		return err
	}
	r.Start()
	return nil
}
