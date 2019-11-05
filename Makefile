#!/bin/bash

fmt:
	gofmt -l -w ./

push:
    git add . && git commit -m "New Algo" && git push