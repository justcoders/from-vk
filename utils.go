package main

import (
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"fmt"
)

func prepareRequest(method, params string) string {
	url := []string{API_URL, method, "?v=5.52&access_token=" + token}
	if params != "" {
		url = append(url, "&"+params)
	}
	return strings.Join(url, "")
}

func createDir(dir string) (err error) {
	return os.Mkdir("."+string(filepath.Separator)+dir, 0777)
}

func getMaxPhotoSize(photo Photo) (size, maxPhoto string) {
	switch {
	case photo.Photo2560 != "":
		size = "2560"
		maxPhoto = photo.Photo2560
	case photo.Photo1280 != "":
		size = "1280"
		maxPhoto = photo.Photo1280
	case photo.Photo807 != "":
		size = "807"
		maxPhoto = photo.Photo807
	case photo.Photo604 != "":
		size = "604"
		maxPhoto = photo.Photo604
	case photo.Photo130 != "":
		size = "130"
		maxPhoto = photo.Photo130
	case photo.Photo75 != "":
		size = "75"
		maxPhoto = photo.Photo75
	}
	return
}

func downloadFile(url string, cb func()) (err error) {
	fileName := path.Base(url)
	out, err := os.Create(outputDir + string(filepath.Separator) + fileName)
	if err != nil {
		fmt.Errorf("%s", err)
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Errorf("%s", err)
		return err
	}
	defer resp.Body.Close()
	defer cb()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Errorf("%s", err)
		return err
	}

	return nil
}
