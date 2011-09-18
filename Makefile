# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

all: install

DIRS=\
	sdl\
	go2d\
	test/basic\
	test/gui\
	skeleton\
	
clean.dirs: $(addsuffix .clean, $(DIRS))
install.dirs: $(addsuffix .install, $(DIRS))

%.clean:
	+cd $* && gomake clean

%.install:
	+cd $* && gomake install

clean: clean.dirs
install: install.dirs

echo-dirs:
	@echo $(DIRS)
