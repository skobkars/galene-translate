#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char *build(char *in) {
    char *out;
    char *s = in;
    char *e;

    if ( *s == '[' )
        s++;
    s++;
    e = strstr(s, "\":");
    if ( !e )
       e = strstr(s, "\",");
    if ( e )
       *e = '\0';
    return s;
}
    
void main(int argc, char **argv) {
    FILE *in = NULL;
    FILE *out = NULL;
    char line[1024];
    char *s, *t;
    
    if ( argc < 3 ) {
        fprintf(stderr, "Syntax: extract <input file> <output file>\n");
        exit(1);
    }
    
    if ((in = fopen(argv[1], "r")) == NULL) {
        fprintf(stderr,"can't open %s\n", argv[1]);
        exit(1);
    }

    if ((out = fopen(argv[2], "w")) == NULL) {
        fprintf(stderr,"can't open %s\n", argv[2]);
        exit(1);
    }
    
    // read line by line;
    while((s = fgets(line, sizeof(line),in))) {
        switch (*s) {
            case '"':
            case '[':
                t = build(s);
                fprintf(out, "%s\n",t);
            break;
        }
    }
    
    exit(0);
}
