// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !appengine

// This file implements a stand-alone blog server.

package main

import (
	"flag"
	"log"
	"net/http"

	"code.google.com/p/go.blog/pkg/blog"
)

var (
	httpAddr     = flag.String("http", "localhost:8080", "HTTP listen address")
	contentPath  = flag.String("content", "content/", "path to content files")
	templatePath = flag.String("template", "template/", "path to template files")
	staticPath   = flag.String("static", "static/", "path to static files")
)

func main() {
	flag.Parse()
	config.ContentPath = *contentPath
	config.TemplatePath = *templatePath
	s, err := blog.NewServer(config)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", s)
	fs := http.FileServer(http.Dir(*staticPath))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}
