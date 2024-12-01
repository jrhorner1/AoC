#include <stdio.h>
#include <stdlib.h>
#include "linklist.h"

// Function to create a new node
Node* createNode(int new_data)
{
    Node* new_node = (Node*)malloc(sizeof(Node));
    new_node->data = new_data;
    new_node->next = NULL;
    return new_node;
}

Node* sortedInsert(struct Node* createNode,  struct Node* sorted) {
    
    // Special case for the head end
    if (sorted == NULL || 
        sorted->data >= createNode->data) {
        createNode->next = sorted;
        sorted = createNode;
    }
    else {
        struct Node* curr = sorted;
        
        // Locate the node before the point of insertion
        while (curr->next != NULL && 
               curr->next->data < createNode->data) {
            curr = curr->next;
        }
        createNode->next = curr->next;
        curr->next = createNode;
    }
    
    return sorted;
}

Node* insertionSort(struct Node* head) {
    
    // Initialize sorted linked list
    struct Node* sorted = NULL;
    struct Node* curr = head;
    
    // Traverse the given linked list and insert
    // every node to sorted
    while (curr != NULL) {
        
        // Store next for next iteration
        struct Node* next = curr->next;
        
        // Insert current in sorted linked list
        sorted = sortedInsert(curr, sorted);
        
        // Update current
        curr = next;
    }
    
    return sorted;
}
