build:	
	xgo --deps=https://gmplib.org/download/gmp/gmp-6.0.0a.tar.bz2 \
			--targets=linux/amd64 -out bin/cache \
			./
	mv bin/cache-linux-amd64 bin/cache