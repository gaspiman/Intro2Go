package main

func main() {
	//	simple()
	//	routines()
	//	syncFunc()
	//	worker()
	//	race_main()
	//	race_fixed()
	server()
}

// Cross compiling:
// GOOS=windows GOARCH=amd64 go build ./...
var webList = []string{
	"http://cnn.com",
	"http://reddit.com",
	"http://google.com",
	"http://news.ycombinator.com",
	"http://theverge.com",
	"http://techcrunch.com",
	"http://golang.org",
	"http://mashable.com",
	"http://facebook.com",
	"http://dir.bg",
	"http://news.bg",
	"http://readwriteweb.com",
	"http://wimp.com",
	"http://twitter.com",
	"http://snapchat.com",
	"http://bestofyoutube.com",
	"http://torrentz.eu",
	"http://github.com",
	"http://uber.com",
	"http://eleven.bg",
	"http://digg.com",
	"http://mediapool.com",
	"http://alchemyapi.com",
	"http://globul.bg",
	"http://t-mobile.com",
	"http://att.com",
	"http://yahoo.com",
	"http://wikipedia.org",
	"http://youtube.com",
	"http://linkedin.com",
	"http://imgur.com",
	"http://bing.com",
	"http://espn.com",
	"http://blogspot.com",
	"http://cnn.com",
	"http://reddit.com",
	"http://google.com",
	"http://news.ycombinator.com",
	"http://theverge.com",
	"http://techcrunch.com",
	"http://golang.org",
	"http://mashable.com",
	"http://facebook.com",
	"http://dir.bg",
	"http://news.bg",
	"http://readwriteweb.com",
	"http://wimp.com",
	"http://twitter.com",
	"http://snapchat.com",
	"http://bestofyoutube.com",
	"http://torrentz.eu",
	"http://github.com",
	"http://uber.com",
	"http://eleven.bg",
	"http://digg.com",
	"http://mediapool.com",
	"http://alchemyapi.com",
	"http://globul.bg",
	"http://t-mobile.com",
	"http://att.com",
	"http://yahoo.com",
	"http://wikipedia.org",
	"http://youtube.com",
	"http://linkedin.com",
	"http://imgur.com",
	"http://bing.com",
	"http://espn.com",
	"http://blogspot.com",
}
