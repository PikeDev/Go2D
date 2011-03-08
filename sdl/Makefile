# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=sdl
CGOFILES=\
	sdl.go\
	ttf.go\
	image.go\
	renderer.go\
	event.go\
	
GOFILES=\
	const.go\

CGO_CFLAGS=`sdl-config --cflags`
CGO_LDFLAGS=`sdl-config --libs` -lSDL_image -lSDL_ttf

CLEANFILES+=sdl

include $(GOROOT)/src/Make.pkg

%: install %.go
	$(GC) $*.go
	$(LD) -o $@ $*.$O
