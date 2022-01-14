//
// Created by Antonio Vencl on 04.01.2022..
//

#include <pthread.h>
#include "list.h"
#ifndef SRSV_CONST_H
#define SRSV_CONST_H

static const int MAX_DURATION = 5;
static const int MAX_SIZE = 26;
static const char ENV_NAME[] = "SRSV_LAB5";
static const int MAX_NUMBER_JOBS = 20;
static const int WAIT_TIME = 10;

typedef struct {
    int id;
    pthread_mutex_t mutexObj;
}Mutex_message;

struct poruka {
    int taskID;
    struct Node *jobList;
    struct Node *taskList;
} poruka;

typedef struct {
    int id;
    int time;
    int jobs[MAX_NUMBER_JOBS];
    struct Node *jobList;
    char envName[MAX_SIZE];
}Task;

typedef struct {
    int messageType;
    char messageText[MAX_SIZE];
}Message;





#endif //SRSV_CONST_H
