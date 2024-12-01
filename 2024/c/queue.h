#ifndef QUEUE_H
# define QUEUE_H

#include <limits.h>
#include <stdio.h>
#include <stdlib.h>
#include "linklist.h"

// Structure to implement queue operations using a linked
// list
typedef struct Queue {

    // Pointer to the front and the rear of the linked list
    Node *front, *rear;
} Queue;

Queue* createQueue();
void sortQueue(Queue*);
int isEmpty(Queue*);
void enqueue(Queue*, int);
void dequeue(Queue*);
int getFront(Queue*);
int getRear(Queue*);

#endif