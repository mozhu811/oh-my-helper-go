package cos

import "testing"

func TestUpload(t *testing.T) {
	url := "https://i0.hdslb.com/bfs/face/257c0607d129c45267baead593cbb16768168255.jpg"

	err := Upload(url, "test.png")
	if err != nil {
		t.Fatal(err)
	}
}
