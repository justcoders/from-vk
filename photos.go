package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"strconv"
)

type Photo struct {
	Photo75   string `json:"photo_75"`
	Photo130  string `json:"photo_130"`
	Photo604  string `json:"photo_604"`
	Photo807  string `json:"photo_807"`
	Photo1280 string `json:"photo_1280"`
	Photo2560 string `json:"photo_2560"`
}

type PhotosResponse struct {
	Response struct {
		Count int
		Items []Photo
	}
}

func importPhotos() {
	photosList, err := getAllPhotos()
	if err != nil {
		log.Fatal(err)
	}

	wg := sync.WaitGroup{}

	i := 0

	for _, photo := range photosList {
		size, maxPhotoURL := getMaxPhotoSize(photo)

		wg.Add(1)
		go downloadFile(maxPhotoURL, func(){
			i++
			wg.Done()
			fmt.Printf("%d) %s %s\n", i, size, maxPhotoURL)
		})
	}

	wg.Wait()
}

func getAllPhotos() (photosList []Photo, err error){
	count, err := getPhotosCount()

	for i :=0 ; i <= count; i += 200  {
		nextPhotoSet, e := getPhotos(strconv.Itoa(i))
		if e != nil {
			fmt.Errorf("offset %d %s", i, e)
			continue
		}
		photosList = append(photosList, nextPhotoSet...)
	}

	return
}

func getPhotos(offset string) (photosList []Photo, err error) {

	res, err := http.Get(prepareRequest(
		"photos.getAll",
		"count=200&extended=0&photo_sizes=0&skip_hidden=0&no_service_albums=0&need_hidden=0&offset=" + offset))

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var photosResponse PhotosResponse
	err = json.Unmarshal(body, &photosResponse)

	photosList = photosResponse.Response.Items

	return
}

func getPhotosCount() (count int, err error) {
	res, e := http.Get(prepareRequest(
		"photos.getAll",
		"count=200&extended=0&photo_sizes=0&skip_hidden=0&no_service_albums=0&need_hidden=0"))

	if e != nil {
		err = e
	}

	body, e := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if e != nil {
		err = e
	}

	var photosResponse PhotosResponse
	json.Unmarshal(body, &photosResponse)

	count = photosResponse.Response.Count

	return
}
