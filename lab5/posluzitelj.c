//
// Created by Antonio Vencl on 04.01.2022..
//

#include <stdio.h>
#include <sys/ipc.h>
#include <sys/msg.h>
#include <signal.h>
#include <stdlib.h>
#include "list.h"
#include "const.h"
#include <sys/shm.h>
#include <string.h>
#include <unistd.h>

// structure for message queue
struct mesg_buffer {
    long mesg_type;
    Task mesg_text;
} message;

struct threadParams {
    pthread_mutex_t *mutex;
    struct Node *jobList;
    struct Node *taskList;
    int number;
} threadParams;

void sigIntHandler(int sig) {
    int key = ftok("progfile", 65);
    int msgid = msgget(key, 0666 | IPC_CREAT);
    // to destroy the message queue
    if(msgctl(msgid, IPC_RMID, NULL)== 0) {
        printf("uspjesno ugasio queue\n");
    }
    exit(1);
}

void printMemory();
void printThreadStatus();


int getPodatak(struct Node **pNode);

int jobNumber(struct Node **pNode);

void appendJobList(struct Node **pNode, int *polje, int size);

const char envName[MAX_SIZE];
int timeoutDuration;
int emptyList = 1;
int prviProlaz = 1;

// Declaration of thread condition variable
pthread_cond_t cond1 = PTHREAD_COND_INITIALIZER;

// declaring mutex
pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;

void *myThreadFun(void *vargp) {
    pthread_mutex_lock(&mutex);
    struct threadParams *params = vargp;
    struct Node *jobList = params->jobList;
    struct Node *taskList = params->taskList;
    int taskId = 0;

    int threadNumber = params->number;
    pthread_cond_signal(&cond1);
    pthread_mutex_unlock(&mutex);

    if(isEmpty(&taskList) == 1) {
        printf("R%d: nema poslova, spavam\n",threadNumber);
    } else {
        if (taskList != NULL) {
            taskId = deleteLast(&taskList);
        }
        // printf("ja radim na task id = %d\n", taskId);
        while (1) {
            sleep(1);
            int brojZadatka = jobNumber(&jobList) +1;
            int podatak = getPodatak(&jobList);
            int duljina = length(&jobList);
            printf("R%d: id:%d obrada podatka: %d (%d/%d)\n", threadNumber, taskId, podatak, brojZadatka, duljina);
            if (brojZadatka == duljina) break;
        }
        printf("R%d: id:%d obrada gotova\n",threadNumber, taskId);
        printf("R%d: nema poslova, spavam\n",threadNumber);
    }

    return NULL;
}

int jobNumber(struct Node **pNode) {
    struct Node *temp = *pNode;
    int counter = 0;
    while (temp->next != NULL) {
        if (temp->data == -1){
            temp = temp->next;
            counter++;
        } else {
            temp = temp->next;
        }
    }
    return counter;
}

int getPodatak(struct Node **pNode) {
    struct Node *temp = *pNode;
    int result;
    if (temp->next == NULL){
        result = temp->data;
        temp->data = -1;
        return result;
    }
    while(temp->next != NULL) {

        if (temp->data != -1) {
            break;
        }
        temp = temp->next;
    }
    result = temp->data;
    temp->data = -1;
    return result;
}


int main(int argc, char *argv[]) {
    struct Node *taskList = createList();
    struct Node *jobList = createList();
    struct Node *jobList2 = createList();
    struct Node *listaPoruka = createList();

    int timer = 0;
    time_t start_t, end_t;
    time(&start_t);



    if (argc != 3) {
        printf("pogresan broj argumenata");
        exit(1);
    }

    timeoutDuration = atoi(argv[2]);

    pthread_t thread_id;

    int threadNumber = atoi(argv[1]);

    struct threadParams params[threadNumber];
    for (int i = 0; i < threadNumber; ++i) {
        params[i].jobList = NULL;
        params[i].taskList = NULL;
        params[i].number = i+1;
        params[i].mutex = &mutex;
    }


    while (1) {
        struct Node *porukaTaskLista = createList();
        // printf("antonio\n");
        // printThreadStatus();
        pthread_t threadId[threadNumber];

        key_t key;
        int msgid;

        // ftok to generate unique key
        key = ftok("progfile", 65);

        // msgget creates a message queue
        // and returns identifier
        msgid = msgget(key, 0666 | IPC_CREAT);


        signal(SIGINT, sigIntHandler);
        // msgrcv to receive message


        msgrcv(msgid, &message, sizeof(message), 1, 0);

        time(&end_t);
        Task task = message.mesg_text;
        printf("P: zaprimio %d %d / [", task.id, task.time);
        for (int i = 0; i < task.time; ++i) {
            printf(" %d ", task.jobs[i]);
        }
        printf("]\n");
        append(&porukaTaskLista, task.id);
        struct Node *jobListToAppend = createList();
        appendJobList(&jobListToAppend, (int *) &task.jobs, task.time);
        struct poruka *por = malloc(sizeof (struct poruka));
        por->taskID = task.id;
        por->taskList = porukaTaskLista;
        por->jobList = jobListToAppend;
        append(&listaPoruka, por);
        task.id++;
        // array size to but flag on last element;
        int arraySize = message.mesg_text.time;
        message.mesg_text.jobs[arraySize] = -1;
        for (int i = 0; i < arraySize; ++i) {
            append(&jobList, task.jobs[i]);
            append(&jobList2, task.jobs[i]);
        }


        if (prviProlaz != timeoutDuration) {
            prviProlaz++;
            continue;
        }
        prviProlaz = 1;
        struct poruka *p;
        int i = 0;
        while(listaPoruka != NULL) {
            for (int j = 0; j < threadNumber; ++j) {
                params[j].taskList = NULL;
                params[j].jobList = NULL;
                if (listaPoruka != NULL) {
                    p = deleteLast(&listaPoruka);
                } else {
                    break;
                }
                params[j].taskList = p->taskList;
                params[j].jobList = p->jobList;
            }
            for (int j = 0; j < threadNumber; ++j) {
                pthread_mutex_lock(&mutex);
                params[j].number = j+1;
                pthread_create(&threadId[j], NULL, myThreadFun, &params[j]);
                pthread_cond_wait(&cond1, &mutex);
                pthread_mutex_unlock(&mutex);
            }

            // p = deleteLast(&listaPoruka);
            // params[1].taskList = p->taskList;
            // params[1].jobList = p->jobList;
            //i++;
            if(i > threadNumber) break;
        }


        emptyList = isEmpty(&taskList);
        if (emptyList != 1){
            pthread_cond_signal(&cond1);
        }


        // display the message

        // printList(taskList);



        strcpy((char *) envName, task.envName);

        // printMemory();



        // to destroy the message queue
        // msgctl(msgid, IPC_RMID, NULL);
    }

    return 0;
}

void appendJobList(struct Node **pNode, int *polje, int size) {
    for (int i = 0; i < size; ++i) {
        append(pNode, polje[i]);
    }
}

void printMemory() {
    // ftok to generate unique key
    key_t keyShared = ftok("shmfile",65);

    // shmget returns an identifier in shmid
    int shmid = shmget(keyShared,1024,0666|IPC_CREAT);

    // shmat to attach to shared memory
    char *str = (char*) shmat(shmid,(void*)0,0);

    printf("Data read from memory: %s\n",str);

    //detach from shared memory
    shmdt(str);

    // destroy the shared memory
    // shmctl(shmid,IPC_RMID,NULL);
}



