
#include "queue.h"

// Function to create a queue
Queue* createQueue()
{
    Queue* q = (Queue*)malloc(sizeof(Queue));
    q->front = q->rear = NULL;
    return q;
}

void sortQueue(Queue* q)
{
    q->front = insertionSort(q->front);
}

// Function to check if the queue is empty
int isEmpty(Queue* q)
{

    // If the front and rear are null, then the queue is
    // empty, otherwise it's not
    if (q->front == NULL && q->rear == NULL) {
        return 1;
    }
    return 0;
}

// Function to add an element to the queue
void enqueue(Queue* q, int new_data)
{

    // Create a new linked list node
    Node* new_node = createNode(new_data);

    // If queue is empty, the new node is both the front
    // and rear
    if (q->rear == NULL) {
        q->front = q->rear = new_node;
        return;
    }

    // Add the new node at the end of the queue and
    // change rear
    q->rear->next = new_node;
    q->rear = new_node;
}

// Function to remove an element from the queue
void dequeue(Queue* q)
{

    // If queue is empty, return
    if (isEmpty(q)) {
        printf("Queue Underflow\n");
        return;
    }

    // Store previous front and move front one node
    // ahead
    Node* temp = q->front;
    q->front = q->front->next;

    // If front becomes null, then change rear also
    // to null
    if (q->front == NULL)
        q->rear = NULL;

    // Deallocate memory of the old front node
    free(temp);
}

// Function to get the front element of the queue
int getFront(Queue* q)
{

    // Checking if the queue is empty
    if (isEmpty(q)) {
        printf("Queue is empty\n");
        return INT_MIN;
    }
    return q->front->data;
}

// Function to get the rear element of the queue
int getRear(Queue* q)
{

    // Checking if the queue is empty
    if (isEmpty(q)) {
        printf("Queue is empty\n");
        return INT_MIN;
    }
    return q->rear->data;
}