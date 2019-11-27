package main

import (
	"crypto/tls"
	"fmt"
)

func main() {
	_, err := tls.Dial("tcp", "62.159.42.66:465", &tls.Config{ServerName: "mail.asiamed.com", MaxVersion: tls.VersionTLS11})
	fmt.Println(err)
	// verbose := flag.Bool("v", false, "verbose")
	// quiet := flag.Bool("q", false, "don't print the data")
	// insecure := flag.Bool("insecure", false, "skip certificate verification")
	// flag.Parse()
	// urls := flag.Args()

	// logger := utils.DefaultLogger

	// if *verbose {
	// 	logger.SetLogLevel(utils.LogLevelDebug)
	// } else {
	// 	logger.SetLogLevel(utils.LogLevelInfo)
	// }
	// logger.SetLogTimeFormat("")

	// pool, err := x509.SystemCertPool()
	// if err != nil {
	// 	panic(err)
	// }
	// testdata.AddRootCA(pool)
	// roundTripper := &http3.RoundTripper{
	// 	TLSClientConfig: &tls.Config{
	// 		RootCAs:            pool,
	// 		InsecureSkipVerify: *insecure,
	// 	},
	// }
	// defer roundTripper.Close()
	// hclient := &http.Client{
	// 	Transport: roundTripper,
	// }

	// var wg sync.WaitGroup
	// wg.Add(len(urls))
	// for _, addr := range urls {
	// 	logger.Infof("GET %s", addr)
	// 	go func(addr string) {
	// 		rsp, err := hclient.Get(addr)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		logger.Infof("Got response for %s: %#v", addr, rsp)

	// 		body := &bytes.Buffer{}
	// 		_, err = io.Copy(body, rsp.Body)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		if *quiet {
	// 			logger.Infof("Request Body: %d bytes", body.Len())
	// 		} else {
	// 			logger.Infof("Request Body:")
	// 			logger.Infof("%s", body.Bytes())
	// 		}
	// 		wg.Done()
	// 	}(addr)
	// }
	// wg.Wait()
}
