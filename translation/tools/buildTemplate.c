#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void build(char *in, FILE *out) {
    char *s = in;
    char *e;
    int list = 0;

    if ( *in == '"' )
        list = 1;

    if (list) {
        e = strstr(in, "\":");
        e[2] = '\0';
        fprintf(out,"%-120s\"\",\n",in);
    } else {
        e = strstr(in, "\",");
        if ( e == NULL ) {
            fprintf(stderr, "input: <%s>", in);
            exit(0);
        }
        e[2] = '\0';
        fprintf(out,"%-120s\"\"],\n",in);
    }

}
    
void main(int argc, char **argv) {
    FILE *in = NULL;
    FILE *out = NULL;
    char line[1024];
    char *s, *t;
    
    if ( argc < 3 ) {
        fprintf(stderr, "Syntax: buildTemplate <input file> <output file>\n");
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
                build(s, out);
            break;
            default:
                fprintf(out,"%s",s);
        }
    }
    
    exit(0);
}
