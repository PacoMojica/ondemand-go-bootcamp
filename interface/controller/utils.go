package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func writeJSON(data []byte, res http.ResponseWriter) {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func writeError(message string, res http.ResponseWriter) {
	log.Println(message)
	fmt.Fprint(res, message)
}

func isMethodValid(method string, res http.ResponseWriter, req *http.Request) bool {
	if req.Method != method {
		writeError(fmt.Sprintf("unsupported method [%v]", req.Method), res)
		return false
	}

	return true
}

func getPaths(res http.ResponseWriter, req *http.Request) ([]string, bool) {
	paths := strings.Split(req.URL.Path[1:], "/")

	if len(paths) != 2 {
		res.WriteHeader(http.StatusNotFound)
		writeError("404 page not found", res)
		return nil, false
	}

	return paths, true
}

func getPathID(paths []string, res http.ResponseWriter) (uint, bool) {
	ID, err := strconv.ParseUint(paths[1], 10, 64)
	if err != nil {
		writeError(fmt.Sprintf("invalid ID: %v", err), res)
		return 0, false
	}

	return uint(ID), true
}
