#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <math.h>

#define RMAX 8
#define DEBUG 0

void get_diffs(int[], int[], int, bool);

int main() {
    char *filename = "../../input/2.ex";
    char line[32];
    char delim = ' ';
    int safe1 = 0;
    int safe2 = 0;

    FILE *input = fopen(filename, "r");
    if ( input != NULL ) {
        while(fgets(line, sizeof(line), input) != NULL) {
            // determine length of line (number of integers)
            int rlen = 1; // to account for string at end
            for (int i = 0; line[i] != '\0'; i++){
                if (line[i] == delim){
                    rlen++;
                }
            }
            if (rlen > RMAX) {
                printf("Counted more 'rlen' than allowed!\n");
                return 1;
            }
            int *report = malloc(sizeof(int[rlen]));
            int *diff = malloc(sizeof(int[rlen-1]));

            // parse line to reports
            char* token = strtok(line, &delim);
            char *end; // not used
            for(int i = 0; token != NULL; i++) {
                int level = (int)strtol(token, &end, 10);
                report[i] = level;
                if ( DEBUG ) {
                    printf("level = %d, report[%d] = %d\n", level, i, report[i]);
                }
                token = strtok(NULL, &delim);
            }
            // determine direction
            // bool descending = true;
            // if (diff[0] - diff[rlen-1] < 0) {
            //     descending = false;
            // }
            // get diffs
            for (int i = 0; i < rlen-1; i++){
                // if (descending) {
                    diff[i] = abs(report[i] - report[i+1]);
                // } else {
                //     diff[i] = report[i+1] - report[i];
                // }
            }

            // print stuff
            for (int i = 0; i < rlen; i++){
                printf("%d ", report[i]);
            }
            printf("||");
            for (int i = 0; i < rlen-1; i++){
                printf(" %d", diff[i]);
            }
            printf("\n");

            /*                 -- Safe criteria --
            * The levels are either all increasing or all decreasing.
            * Any two adjacent levels differ by at least one and at most three.
            */
            int unsafe = 0;
            for ( int i = 0; i < rlen-1; i++ ){
                if (diff[i] < 1 && diff[i] > 3 ){
                    unsafe++;
                }
            }
            if (DEBUG) {
                printf("unsafe = %d\n", unsafe);
            }

            if (unsafe == 0) {
                safe1++;
            }
        }
    } else  {
        printf("File doesn't exist.");
        return 1;
    }
    fclose(input);

    printf("Part 1: %d\n", safe1);
    printf("Part 2: %d\n", safe2);

    return 0;
}