# go-qrcode
> CLI to generate QR code PNG files of URLs

## Usage

```
go-qrcode -url https://google.com
# output: 2023/12/22 21:24:23 INFO written file '<path>/qr.png' size=1024
```

## Help Output

```bash
  -filePath string
    	file path of the .png file (default "./qr.png")
  -size int
    	width/height length of the .png file in px (default 1024)
  -url string
    	URL to encode in the QR code
```