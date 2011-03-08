package go2d

import "container/list"

var g_resources *list.List = list.New()

type IResource interface {
	release()
}

func addResource(_resource IResource) {
	g_resources.PushBack(_resource)
}

func freeResources() {
	for i := g_resources.Front(); i != nil; i = i.Next() {
		res, valid := i.Value.(IResource)
		if valid {
			res.release()
		}
	}
}
