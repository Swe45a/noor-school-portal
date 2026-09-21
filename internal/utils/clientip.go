package utils

import (
	"net"
	"net/http"
	"strings"
)

// ClientIP extracts the caller's IP for rate-limiting purposes. It trusts X-Forwarded-For only
// as a fallback since this service may sit behind a reverse proxy in front of the school network.
func ClientIP(r *http.Request) string {
	if fwd := r.Header.Get("X-Forwarded-For"); fwd != "" {
		parts := strings.Split(fwd, ",")
		ip := strings.TrimSpace(parts[0])
		if ip != "" {
			return ip
		}
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
