Porter Stemmer for Go
=====================

This is a fairly straighforward port of Martin Porter's C implementation
of the Porter stemming algorithm. The C version this port is based on is
available for download here:
[http://tartarus.org/~martin/PorterStemmer/c_thread_safe.txt](http://tartarus.org/~martin/PorterStemmer/c_thread_safe.txt)

The original algorithm is described in the paper:

    M.F. Porter, 1980, An algorithm for suffix stripping, Program, 14(3) pp
130-137.


While the internal implementation and interface is nearly identical to
the original implementation, the Go interface is much simplified. The
stemmer can be called as follows:

    import "porter"
    ...
    stemmed := porter.Stem(word_to_stem)


Limitations
-----------

While the implementation is fairly robust, this is a work in progress.
In particular, a new interface will likely be provided to prevent
excessive conversions between `string`s and `[]byte`. Currently, on
calling `Stem` the string argument is converted to a byte slice which
the algorithm works on and is converted back into a string before
returning.

Also, the implementation is not particularly robust at handling Unicode
input, currently, only bytes with the high bit set are ignored. It's up
to the caller to make sure the string contains only ASCII characters.
Since the algorithm itself operates on English words only, this doens't
restrict the functionality, but it is nuisance.

TODO:
----- 
* docs, inline, howto goinstall
* attribution
