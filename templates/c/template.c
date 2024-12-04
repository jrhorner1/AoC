#include <stdio.h>

int main() {
    char *filename = "../../input/0";
    char line[32];

    FILE *input = fopen(filename, "r");
    if ( input != NULL ) {
        while(fgets(line, sizeof(line), input) != NULL) {
            // parse line .. do stuff
        }
    } else  {
        printf("File doesn't exist.");
        return 1;
    }
    fclose(input);

    printf("Part 1: %d\n", 42);
    printf("Part 2: %d\n", 5/7);

    return 0;
}