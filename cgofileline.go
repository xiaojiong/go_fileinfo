package go_fileinfo

/*
#include <stdio.h>

int scanlineaa(char *filename)
{
    FILE * f = 0;
    char line[256]="";
    int lines = 0;
    f = fopen(filename, "r");
    if(!f) return 0;
    while(!feof(f)) { fgets(line, 256, f); lines++;}
    fclose(f);
    return lines;
}

*/
import "C"

func FileLine(str *string) int {
	CStr := C.CString(*str)

	line := int(C.scanlineaa(CStr))
	return line
}
