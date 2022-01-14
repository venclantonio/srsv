//
// Created by Antonio Vencl on 04.01.2022..
//

// C Program for Message Queue (Writer Process)
#include <stdio.h>
#include <sys/ipc.h>
#include <sys/msg.h>
#include <stdlib.h>
#include <string.h>
#include "const.h"
#include <sys/shm.h>
#include <sys/errno.h>
#include <sys/fcntl.h>
#include <sys/mman.h>
#include <unistd.h>

#define MAX 10

// structure for message queue
struct mesg_buffer {
    long mesg_type;
    //char mesg_text[100];
    Task mesg_text;
} message;

Task newTask(int id, int duration);
void printTask(Task task);
void printMessage(Task task);
void writeIdToMemory(int id, const char *);
int getLastIdFromMemory(const char *);

const int data = 8;
const char *charVal[3];
const char envName[MAX_SIZE];

void *myThreadFun(void *vargp)
{
    sleep(1);
    printf("Printing GeeksQuiz from Thread \n");
    return NULL;
}

int main(int argc, char *argv[]) {
    // pthread_t thread_id;
    // printf("Before Thread\n");
    // pthread_create(&thread_id, NULL, myThreadFun, NULL);
    //pthread_join(thread_id, NULL);
    // printf("After Thread\n");

    int J,maxDuration;
    if (argc != 3) {
        printf("Pogresan broj argumenata");
        exit(1);
    }

    J = atoi(argv[1]);
    maxDuration = atoi(argv[2]);

    stpcpy((char *) envName, getenv(ENV_NAME));//getenv(ENV_NAME));
    int lastId = getLastIdFromMemory(envName);
    // printf("Ovo je lastId iz memorije = %d", lastId);
    srand((unsigned)time(0));
    for (int i = 0; i < J; ++i) {
        key_t key;
        int msgid, openFile = 1;



        // ftok to generate unique key
        key = ftok("progfile", 65);

        // msgget creates a message queue
        // and returns identifier
        msgid = msgget(key, 0666 | IPC_CREAT);
        message.mesg_type = 1;


        Task task = newTask(i + lastId,maxDuration);
        message.mesg_text = task;
        // printTask(task);
        printMessage(task);


        // destroy the shared memory
        // shmctl(shmidread,IPC_RMID,NULL);

        // printf("%s", envName);
        writeIdToMemory(1, envName);
        // printf("\nLast id in memory is = %d\n",getLastIdFromMemory(envName));

        // ftok to generate unique key
        key_t keyShared = ftok("shmfile",65);
        if(keyShared==-1) {
            // printf("Oh dear, something went wrong with read()! %s\n", strerror(errno));
        }
        // printf("key for writing = %d\n", keyShared);

        // shmget returns an identifier in shmid
        int shmid = shmget(keyShared,1024,0666|IPC_CREAT);

        // shmat to attach to shared memory
        char *str = (char*) shmat(shmid,(void*)0,0);

        sprintf((char *) charVal, "%d", data);

        memmove(str, (const void *) charVal, 3);

        // printf("Data written in memory: %s\n",str);



        //detach from shared memory
        shmdt(str);
        msgsnd(msgid, &message, sizeof(message), 0);


    }
    // pthread_join(thread_id, NULL);
    return 0;
}

Task newTask (int id, int taskDuration) {
    Task result;
    // srand(time(0));
    result.id= id;
    result.time = rand() % taskDuration + 1;
    for (int i = 0; i < result.time; ++i) {
        int number = rand() % 1000;
        result.jobs[i] = number;
        append(&result.jobList, number);
    }
    memcpy(result.envName, getenv(ENV_NAME), strlen(getenv(ENV_NAME)) + 1);
    return result;
}

void printTask(Task task) {
    printf("id = %d  duration = %d  [", task.id, task.time);
    for (int i = 0; i < task.time; ++i) {
        printf(" %d ", task.jobs[i]);
    }
    printf("]\n");
}

void printMessage(Task task) {
    printf("G: posao %d %d /%s  [", task.id, task.time, task.envName);
    for (int i = 0; i < task.time; ++i) {
        printf(" %d ", task.jobs[i]);
    }
    printf("]\n");
}

void writeIdToMemory(int id, const char *env) {
    // printf("poslije ovog\n");
    int openFile;
    openFile = shm_open(env, O_CREAT | O_RDWR, 0777);


    int truncate = ftruncate(openFile, sizeof(Mutex_message));
    Mutex_message *memoryObj = mmap(0, sizeof(Mutex_message), PROT_WRITE | PROT_READ, MAP_SHARED,openFile,0);
    memoryObj->id++;
    //ssize_t n = write(1,memoryObj, sizeof(Mutex_message));
}

int getLastIdFromMemory(const char *env) {
    int openFile;
    openFile = shm_open(env, O_CREAT | O_RDWR, 0777);
    if(shm_open("mirko", O_CREAT | O_RDWR, 0777) == -1) {
        printf("%s", strerror(errno));
    }
    // printf("openFile = %d", openFile);
    int truncate = ftruncate(openFile, sizeof(Mutex_message));
    Mutex_message *memoryObj = mmap(0, sizeof(Mutex_message), PROT_WRITE | PROT_READ, MAP_SHARED,openFile,0);
    //printf("%d", memoryObj->id);

    if (memoryObj->id < 1) {
        return 1;
    } else {
        return memoryObj->id;
    }
}

