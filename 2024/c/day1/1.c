#include <stdio.h>
#include "../queue.h"

int part1(Queue*, Queue*);
int part2(Queue*, Queue*);

int main() {
    char *filename = "../../input/1";   
    // char *filename = "example";

    FILE *input = fopen(filename, "r");
    if ( input == NULL ) {
        printf("File doesn't exist.");
        return 1;
    }

    Queue* q1 = createQueue();
    Queue* q2 = createQueue();
    int a, b;
    while ( fscanf(input, "%d   %d", &a, &b) == 2 ) {
        enqueue(q1, a);
        enqueue(q2, b);
    }

    fclose(input);

    int p2 = part2(q1, q2);
    int p1 = part1(q1, q2);

    printf("Part 1: %d\n", p1);
    printf("Part 2: %d\n", p2);

    return 0;
}

int part1(Queue* q1, Queue* q2) {
    sortQueue(q1);
    sortQueue(q2);

    Node *c1 = q1->front;
    Node *c2 = q2->front;
    int total_dist = 0;
    while ( c1 != NULL ) {
        int dist;
        if (c2->data > c1->data) {
            dist = c2->data - c1->data;
        } else {
            dist = c1->data - c2->data;
        }
        total_dist = total_dist + dist;
        c1 = c1->next;
        c2 = c2->next;
    }

    return total_dist;
}

int part2(Queue* q1, Queue* q2) {
    Node *c1 = q1->front;
    int similarity_score = 0;
    while ( c1 != NULL ) {
        int count = 0;
        Node *c2 = q2->front;
        while ( c2 != NULL ) {
            if ( c2->data == c1->data ) {
                count++;
            }
            c2 = c2->next;
        }
        int similarity = c1->data * count;
        similarity_score += similarity;
        c1 = c1->next;
    }

    return similarity_score;
}