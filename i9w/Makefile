.PHONY: all
all: beattai

beattai: i9w.c vendor/libtai/libtai.a
	cc -o $@ i9w.c vendor/libtai/libtai.a -lm

.PHONY: clean
clean:
	-rm beattai
