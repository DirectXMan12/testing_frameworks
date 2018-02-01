package internal

import (
	"fmt"
	"net/url"
)

func MakeEtcdArgs(input DefaultedProcessInput) []string {
	args := []string{
		"--debug",
		"--listen-peer-urls=http://localhost:0",
		fmt.Sprintf("--advertise-client-urls=%s", input.URL.String()),
		fmt.Sprintf("--listen-client-urls=%s", input.URL.String()),
		fmt.Sprintf("--data-dir=%s", input.Dir),
	}
	return args
}

func isSecureScheme(scheme string) bool {
	// https://github.com/coreos/etcd/blob/d9deeff49a080a88c982d328ad9d33f26d1ad7b6/pkg/transport/listener.go#L53
	if scheme == "https" || scheme == "unixs" {
		return true
	}
	return false
}

func GetEtcdStartMessage(listenUrl url.URL) string {
	if isSecureScheme(listenUrl.Scheme) {
		// https://github.com/coreos/etcd/blob/a7f1fbe00ec216fcb3a1919397a103b41dca8413/embed/serve.go#L167
		return "serving client requests on " + listenUrl.Hostname()
	}

	// https://github.com/coreos/etcd/blob/a7f1fbe00ec216fcb3a1919397a103b41dca8413/embed/serve.go#L124
	return "serving insecure client requests on " + listenUrl.Hostname()
}
