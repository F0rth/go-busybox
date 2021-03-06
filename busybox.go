package busybox

/*
#cgo CFLAGS: -Iinclude
#cgo LDFLAGS: -L0_lib -lbusybox

#include <unistd.h>
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

extern int lbb_main(char **argv);

char* bb_applet(char* applet){
	int fd;
	int fd2;
	fpos_t pos;
	long lSize;
	char *buffer;
	char *argument;
	char **strarray;

	strarray = malloc(strlen(applet));
	argument = strtok(applet, " ");
	int i=0;
	while(argument != NULL){
		strarray[i] = malloc(sizeof(argument) * sizeof(char));
		strcpy(strarray[i], argument);
		argument = strtok(NULL," ");
		i++;
	}

	strarray[i+1]=0;
	//char *strarray[] = {applet, 0};
	// capture stdout
	fflush(stdout);
	fgetpos(stdout, &pos);
	fd = dup(fileno(stdout));
	// call the busybox applet
	lbb_main(strarray);
	//
    fflush(stdout);
    fseek(stdout, 0L, SEEK_END);
	lSize = ftell(stdout);
	rewind(stdout);
	// allocate memory for entire content
	buffer = calloc( 1, lSize+1 );
	// copy the file into the buffer
	fread( buffer , lSize, 1 , stdout);
	// release stdout
	dup2(fd, fileno(stdout));
	close(fd);
	clearerr(stdout);
	fsetpos(stdout, &pos);
	//
	return buffer;
}
*/
import "C"

func Applet(command string) (string) {
	return C.GoString(C.bb_applet(C.CString(command)))
}
