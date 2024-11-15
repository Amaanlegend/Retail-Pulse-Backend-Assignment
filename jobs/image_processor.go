package jobs

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

func ProcessImages(storeID string, imageURLs []string) error {
    for _, url := range imageURLs {
        img, err := downloadImage(url)
        if err != nil {
            return fmt.Errorf("failed to download image: %v", err)
        }

        // Calculate perimeter
        perimeter := 2 * (img.Width + img.Height)

        // Simulate processing delay
        time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
    }
    return nil
}

func downloadImage(url string) (*Image, error) {
    resp, err := http.Get(url)
    if err != nil || resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to fetch image: %v", err)
    }
    defer resp.Body.Close()

    img, _, err := image.Decode(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("failed to decode image: %v", err)
    }

    return &Image{Width: img.Bounds().Dx(), Height: img.Bounds().Dy()}, nil
}

type Image struct {
    Width  int
    Height int
}
