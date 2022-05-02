package main

import "strings"

func getResourceNameFromSelfLink(link string) string {
	parts := strings.Split(link, "/")
	return parts[len(parts)-1]
}

func getProjectNameFromSlefLink(link string) string {
	parts := strings.SplitAfter(link, "projects/")
	project := strings.Split(parts[1], "/")
	return project[0]
}
