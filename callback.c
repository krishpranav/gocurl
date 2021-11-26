#include <stdio.h>
#include <string.h>
#include "callback.h"
#include "_cgo_export.h"

size_t header_function( char *ptr, size_t size, size_t nmemb, void *ctx) {
	return goCallHeaderFunction(ptr, size*nmemb, ctx);
}

void *return_header_function() {
    return (void *)&header_function;
}


size_t write_function( char *ptr, size_t size, size_t nmemb, void *ctx) {
	return goCallWriteFunction(ptr, size*nmemb, ctx);
}

void *return_write_function() {
    return (void *)&write_function;
}

size_t read_function( char *ptr, size_t size, size_t nmemb, void *ctx) {
	return goCallReadFunction(ptr, size*nmemb, ctx);
}

void *return_read_function() {
    return (void *)&read_function;
}

int progress_function(void *ctx, double dltotal, double dlnow, double ultotal, double ulnow) {
	return goCallProgressFunction(dltotal, dlnow, ultotal, ulnow, ctx);
}

void *return_progress_function() {
    return (void *)progress_function;
}