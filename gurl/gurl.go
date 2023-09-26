package main

import (
	"crypto/tls"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptrace"
	"os"
	"strconv"
	"strings"
)

func main() {
	argsLength := len(os.Args[1:])
	slog.Debug("Number of CLI args:", "length", argsLength)
	if argsLength < 1 {
		slog.Error("URL is required!")
		os.Exit(1)
	}

	url := os.Args[1]
	if !strings.HasPrefix(strings.ToLower(url), "https://") {
		slog.Error("Only HTTPS protocol is supported!")
		os.Exit(1)
	}

	//doGet(url)
	req, _ := http.NewRequest("GET", url, nil)
	doRequest(req)
}

func doGet(url string) {
	slog.Info("GET", "URL", url)
	resp, err := http.Get(url)
	if err != nil {
		slog.Error("HTTP GET error", "message", err.Error())
		os.Exit(1)
	}
	slog.Info("Request successful!", "phrase", resp.Status)
	defer resp.Body.Close()
	//body, err := io.ReadAll(resp.Body)
	body, err := io.Copy(os.Stdout, resp.Body)
	if err != nil {
		slog.Error("Error when reading response body", "message", err.Error())
	}
	slog.Info("Response body", "length", body)
}

func buildTrace() *httptrace.ClientTrace {
	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			slog.Info("Get connection", "target", hostPort)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			slog.Info("Got Conn: ",
				slog.String("network", connInfo.Conn.RemoteAddr().Network()),
				slog.String("IP", connInfo.Conn.RemoteAddr().String()),
			)
		},
		ConnectStart: func(network, addr string) {
			slog.Info("start connection", "network", network, "address", addr)
		},
		DNSStart: func(info httptrace.DNSStartInfo) {
			slog.Info("DNS lookup", "target", info.Host)
		},
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			slog.Info("DNS Info: ",
				slog.Any("sd", dnsInfo.Addrs),
			)
		},
		TLSHandshakeStart: func() {
			slog.Info("TLS handshake start")
		},
		TLSHandshakeDone: func(state tls.ConnectionState, err error) {
			if err != nil {
				slog.Error("TLS handshake error", "message", err.Error())
				return
			}
			slog.Info("TLS handshake done: ",
				slog.String("protocol", state.NegotiatedProtocol),
				slog.String("version", strconv.Itoa(int(state.Version))),
				slog.String("server", state.ServerName),
				slog.Uint64("cipher", uint64(state.CipherSuite)),
			)
		},
	}
	return trace
}

func doRequest(req *http.Request) {
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), buildTrace()))
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		slog.Error(err.Error())
	}
}
