#ifndef LINKLIST_H
# define LINKLIST_H

// Node structure representing a single node in the linked
// list
typedef struct Node {
    int data;
    struct Node* next;
} Node;

Node* createNode(int);
Node* sortedInsert(struct Node*,  struct Node*);
Node* insertionSort(struct Node*);

#endif