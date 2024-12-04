#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void build(char *in, char *trans, FILE *out) {
    char *e;
    int list = 0;

    if ( *in == '"' )
        list = 1;

    if (list) {
        e = strstr(in, "\":");
        e[2] = '\0';
        fprintf(out,"%-120s\"%s\",\n",in,trans);
    } else {
        e = strstr(in, "\",");
        e[2] = '\0';
        fprintf(out,"%-120s\"%s\"],\n",in,trans);
    }
}
    
void main(int argc, char **argv) {
    FILE *in = NULL;
    FILE *trans = NULL;
    FILE *out = NULL;
    char line[1024];
    char transLine[1024];
    char *s, *t;
    
    if ( argc < 4 ) {
        fprintf(stderr, "Syntax: merge <template file> <translation file> <output file>\n");
        exit(1);
    }
    
    if ((in = fopen(argv[1], "r")) == NULL) {
        fprintf(stderr,"can't open %s\n", argv[1]);
        exit(1);
    }

    if ((trans = fopen(argv[2], "r")) == NULL) {
        fprintf(stderr,"can't open %s\n", argv[2]);
        exit(1);
    }

    if ((out = fopen(argv[3], "w")) == NULL) {
        fprintf(stderr,"can't open %s\n", argv[3]);
        exit(1);
    }
    
    // read line by line;
    while((s = fgets(line, sizeof(line),in))) {
        switch (*s) {
            case '"':
            case '[':
                if (!fgets(transLine, sizeof(transLine),trans)) {
                    fprintf(stderr,"Error: wrong translation text\n");
                    exit(1);
                }
                t = strchr(transLine,'\n');
                if (t)
                   *t = '\0';
                build(s, transLine, out);
            break;
            default:
               fprintf(out,  "%s",s);
        }
    }
    
    exit(0);
}
