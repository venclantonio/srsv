//
// Created by Antonio Vencl on 04.01.2022..
//

#ifndef SRSV_LIST_H
#define SRSV_LIST_H

#include <stdlib.h>
#include <stdio.h>

struct Node
{
    void *data;
    struct Node *next;
};

struct Node *createList() {
    struct Node *start = NULL;
    return start;
}

void append(struct Node** head_ref, void *new_data)
{
    /* 1. allocate node */
    struct Node* new_node = (struct Node*) malloc(sizeof(struct Node));

    struct Node *last = *head_ref;  /* used in step 5*/

    /* 2. put in the data  */
    new_node->data  = new_data;

    /* 3. This new node is going to be the last node, so make next of
          it as NULL*/
    new_node->next = NULL;

    /* 4. If the Linked List is empty, then make the new node as head */
    if (*head_ref == NULL)
    {
        *head_ref = new_node;
        return;
    }

    /* 5. Else traverse till the last node */
    while (last->next != NULL)
        last = last->next;

    /* 6. Change the next of last node */
    last->next = new_node;
    return;
}

void printList(struct Node* n)
{
    while (n != NULL) {
        printf(" %d ", n->data);
        n = n->next;
    }
}

void deleteList(struct Node ** head){
    struct Node *first = *head;
    *head = NULL;
}


void *deleteLast(struct Node ** head){
    void *result;
    struct Node *last = *head;
    if (last->next != NULL) {
        while (last->next->next != NULL)
            last = last->next;
    }
    if (last->next != NULL) {
        result = last->next->data;
    } else {
        result = last->data;
        *head = NULL;
        return result;
    }
    last->next = NULL;
    return result;
}

void deleteNode(struct Node** head_ref, int key)
{
    // Store head node
    struct Node *temp = *head_ref, *prev;

    // If head node itself holds the key to be deleted
    if (temp != NULL && temp->data == key) {
        *head_ref = temp->next; // Changed head
        free(temp); // free old head
        return;
    }

    // Search for the key to be deleted, keep track of the
    // previous node as we need to change 'prev->next'
    while (temp != NULL && temp->data != key) {
        prev = temp;
        temp = temp->next;
    }

    // If key was not present in linked list
    if (temp == NULL)
        return;

    // Unlink the node from linked list
    prev->next = temp->next;

    free(temp); // Free memory
}

int isEmpty(struct Node **head) {
    struct Node *temp = *head;
    if (temp == NULL) {
        return 1;
    }
    return 0;
}

int length(struct Node **head) {
    struct Node *temp = *head;
    int counter = 0;
    while (temp != NULL) {
        temp = temp->next;
        counter++;
    }
    return counter;
}



#endif //SRSV_LIST_H
